package repo

import (
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/gitlib/util"
	"github.com/bitwormhole/starter/lang"
)

// FileRepositoryFactory 默认的仓库工厂
type FileRepositoryFactory struct {
	Pipeline []ElementFactory
}

func (inst *FileRepositoryFactory) _Impl() repository.Factory {
	return inst
}

func (inst *FileRepositoryFactory) isComplex(location *repository.Location) bool {
	submod := location.SubmoduleDirectory
	wktree := location.WorktreeDirectory
	if submod == nil && wktree == nil {
		return false
	}
	return true
}

func (inst *FileRepositoryFactory) initViewport(vpt *ViewportContext) error {

	// make
	err := inst.make(vpt)
	if err != nil {
		return err
	}

	// link
	err = inst.link(vpt)
	if err != nil {
		return err
	}

	// activate
	err = inst.activate(vpt)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开仓库视图
func (inst *FileRepositoryFactory) Open(location *repository.Location) (repository.Viewport, error) {

	// 准备 pool
	pool := lang.CreateReleasePool()
	poolHolder := &util.ReleasePoolHolder{}
	poolHolder.Init(pool)
	defer poolHolder.Release()

	// 准备 center
	center := &CenterContext{}
	layoutCore := &Layout{}
	core := &ViewportContext{}

	center.core = core
	center.pool = pool
	core.center = center
	core.elements = make([]Element, 0, 20)

	// layout the core
	err := layoutCore.initCore(location)
	if err != nil {
		return nil, err
	}

	if !inst.isComplex(location) {
		// as a simple viewport
		err := inst.initViewport(core)
		if err != nil {
			return nil, err
		}
		poolHolder.Disconnect()
		return core.ToViewport()
	}

	// as a complex viewport
	layoutView := &Layout{}
	err = layoutView.initView(location)
	if err != nil {
		return nil, err
	}

	view := &ViewportContext{}
	view.center = center
	view.core = core.core
	view.layout = layoutView
	view.elements = make([]Element, 0, 20)

	err = inst.initViewport(view)
	if err != nil {
		return nil, err
	}

	poolHolder.Disconnect()
	return view.ToViewport()
}

func (inst *FileRepositoryFactory) make(vc *ViewportContext) error {
	efs := inst.Pipeline
	for _, ef := range efs {
		err := ef.Make(vc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *FileRepositoryFactory) link(vc *ViewportContext) error {
	list := vc.elements
	for _, item := range list {
		err := item.Link()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *FileRepositoryFactory) activate(vc *ViewportContext) error {
	pool := vc.center.pool
	list := vc.elements
	for _, item := range list {
		err := item.Init()
		if err != nil {
			return err
		}
		pool.Push(item)
	}
	return nil
}
