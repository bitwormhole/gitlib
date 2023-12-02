package repositories

import (
	"fmt"
	"io"

	"github.com/bitwormhole/gitlib/utils"
	"github.com/starter-go/application"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/vlog"
)

// Lifecycle 组件生命周期的注册信息
type Lifecycle struct {
	OnCreate    application.OnCreateFunc
	OnStartPre  application.OnStartPreFunc
	OnStart     application.OnStartFunc
	OnStartPost application.OnStartPostFunc

	// OnLoop      application.OnLoopFunc

	OnStopPre  application.OnStopPreFunc
	OnStop     application.OnStopFunc
	OnStopPost application.OnStopPostFunc
	OnDestroy  application.OnDestroyFunc
}

// LifecycleRegistry ...
type LifecycleRegistry interface {
	Lifecycle() *Lifecycle
}

////////////////////////////////////////////////////////////////////////////////

// LifecycleManager ...
type LifecycleManager interface {
	AddComponent(com any) (*Lifecycle, error)
	AddLifecycle(lc *Lifecycle) (*Lifecycle, error)
	AddRegistry(reg LifecycleRegistry) (*Lifecycle, error)

	// Open a new life
	Open() (io.Closer, error)
}

// NewLifecycleManager 新建一个生命周期管理器
func NewLifecycleManager(mode safe.Mode) LifecycleManager {
	man := &innerLifecycleManager{}
	man.init(mode)
	return man
}

////////////////////////////////////////////////////////////////////////////////

type innerLifecycleManager struct {
	mode  safe.Lock
	items []*Lifecycle
}

func (inst *innerLifecycleManager) init(mode safe.Mode) {
	if mode == nil {
		mode = safe.Safe()
	}
	inst.mode = mode.NewLock()
	inst.items = make([]*Lifecycle, 0)
}

func (inst *innerLifecycleManager) lock() {
	inst.mode.Lock()
}

func (inst *innerLifecycleManager) unlock() {
	inst.mode.Unlock()
}

func (inst *innerLifecycleManager) add(lc *Lifecycle) (*Lifecycle, error) {
	if lc == nil {
		return nil, fmt.Errorf("param:Lifecycle is nil")
	}
	inst.items = append(inst.items, lc)
	return lc, nil
}

func (inst *innerLifecycleManager) AddComponent(com any) (*Lifecycle, error) {
	reg, ok := com.(LifecycleRegistry)
	if !ok {
		return nil, fmt.Errorf("Lifecycle is unsupported")
	}
	inst.lock()
	defer inst.unlock()
	return inst.add(reg.Lifecycle())
}

func (inst *innerLifecycleManager) AddLifecycle(lc *Lifecycle) (*Lifecycle, error) {
	if lc == nil {
		return nil, fmt.Errorf("param:Lifecycle is nil")
	}
	inst.lock()
	defer inst.unlock()
	return inst.add(lc)
}

func (inst *innerLifecycleManager) AddRegistry(reg LifecycleRegistry) (*Lifecycle, error) {
	if reg == nil {
		return nil, fmt.Errorf("param:LifecycleRegistry is nil")
	}
	inst.lock()
	defer inst.unlock()
	return inst.add(reg.Lifecycle())
}

// Open a new life
func (inst *innerLifecycleManager) Open() (io.Closer, error) {
	life := &innerLifecycleRuntime{parent: inst}
	err1 := life.open()
	if err1 == nil {
		return life, nil
	}
	err2 := life.Close()
	if err2 != nil {
		vlog.Error(err2.Error())
	}
	return nil, err1
}

////////////////////////////////////////////////////////////////////////////////

type innerLifecycleAdapterCloser struct {
	fn func() error
}

func (inst *innerLifecycleAdapterCloser) Close() error {
	f := inst.fn
	if f == nil {
		return nil
	}
	defer func() {
		x := recover()
		inst.handlePanic(x)
	}()
	return f()
}

func (inst *innerLifecycleAdapterCloser) handlePanic(x any) {
	err := utils.RecoverError(x)
	if err != nil {
		vlog.Warn(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////

type innerLifecycleRuntime struct {
	parent      *innerLifecycleManager
	caughtPanic error
	all         []*Lifecycle
	startedList []*Lifecycle
	closerList  []io.Closer
}

func (inst *innerLifecycleRuntime) handlePanic(x any) {
	inst.caughtPanic = utils.RecoverError(x)
}

func (inst *innerLifecycleRuntime) invoke(lc *Lifecycle, fn func(lc *Lifecycle) error) error {
	if fn == nil || lc == nil {
		return nil
	}
	defer func() {
		x := recover()
		inst.handlePanic(x)
	}()
	return fn(lc)
}

func (inst *innerLifecycleRuntime) forEach(list []*Lifecycle, fn func(lc *Lifecycle) error) error {
	inst.caughtPanic = nil
	for _, item := range list {
		err := inst.invoke(item, fn)
		if err != nil {
			return err
		}
		err = inst.caughtPanic
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *innerLifecycleRuntime) openDoCreate() error {
	items := inst.all
	err := inst.forEach(items, func(lc *Lifecycle) error {
		fn1 := lc.OnCreate
		fn2 := lc.OnDestroy
		if fn1 != nil {
			err := fn1()
			if err != nil {
				return err
			}
		}
		inst.addToCloserList(fn2)
		return nil
	})
	return err
}

func (inst *innerLifecycleRuntime) openDoStart1() error {
	items := inst.all
	err := inst.forEach(items, func(lc *Lifecycle) error {
		fn := lc.OnStartPre
		if fn == nil {
			return nil
		}
		return fn()
	})
	return err
}

func (inst *innerLifecycleRuntime) openDoStart2() error {
	items := inst.all
	err := inst.forEach(items, func(lc *Lifecycle) error {
		fn := lc.OnStart
		if fn != nil {
			e2 := fn()
			if e2 != nil {
				return e2
			}
		}
		// add to started list
		inst.startedList = append(inst.startedList, lc)
		return nil
	})
	return err
}

// 为已经启动的组件添加 closer 钩子
func (inst *innerLifecycleRuntime) addCloserForStartedItems() error {
	items := inst.startedList
	inst.forEach(items, func(lc *Lifecycle) error {
		inst.addToCloserList(lc.OnStopPost)
		return nil
	})
	inst.forEach(items, func(lc *Lifecycle) error {
		inst.addToCloserList(lc.OnStop)
		return nil
	})
	inst.forEach(items, func(lc *Lifecycle) error {
		inst.addToCloserList(lc.OnStopPre)
		return nil
	})
	return nil
}

func (inst *innerLifecycleRuntime) openDoStart3() error {
	items := inst.all
	err := inst.forEach(items, func(lc *Lifecycle) error {
		fn := lc.OnStartPost
		if fn == nil {
			return nil
		}
		return fn()
	})
	return err
}

func (inst *innerLifecycleRuntime) open() error {

	inst.parent.lock()
	defer inst.parent.unlock()
	defer inst.addCloserForStartedItems()

	items := inst.parent.items
	inst.all = nil
	inst.all = append(inst.all, items...)

	steps := make([]func() error, 0)
	steps = append(steps, inst.openDoCreate)
	steps = append(steps, inst.openDoStart1)
	steps = append(steps, inst.openDoStart2)
	steps = append(steps, inst.openDoStart3)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerLifecycleRuntime) addToCloserList(fn func() error) {
	if fn == nil {
		return
	}
	adapter := &innerLifecycleAdapterCloser{fn: fn}
	inst.closerList = append(inst.closerList, adapter)
}

func (inst *innerLifecycleRuntime) Close() error {
	inst.parent.lock()
	defer inst.parent.unlock()
	list := inst.closerList
	count := len(list)
	for i := count - 1; i >= 0; i-- {
		item := list[i]
		err := item.Close()
		if err != nil {
			vlog.Warn(err.Error())
		}
	}
	return nil
}
