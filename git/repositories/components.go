package repositories

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// ComponentRegistration 表示组件的注册信息
type ComponentRegistration struct {
	Name     string
	Class    string // 格式参考： <html class="" />
	Alias    string // 格式参考： <html class="" />
	Enabled  bool
	Priority int

	// 以下是针对各种上下文的初始化处理函数
	// 它们的返回值为（组件对象,错误信息）
	OnInitForSystem     func(ctx *SystemContext) (any, error)
	OnInitForUser       func(ctx *UserContext) (any, error)
	OnInitForRepository func(ctx *RepositoryContext) (any, error)
	OnInitForSession    func(ctx *SessionContext) (any, error)
	OnInitForWorktree   func(ctx *WorktreeContext) (any, error)
	OnInitForSubmodule  func(ctx *SubmoduleContext) (any, error)
}

// ComponentRegistry ...
type ComponentRegistry interface {
	ListRegistrations() []*ComponentRegistration
}

// ComponentHolder ...
type ComponentHolder struct {
	Component any // 缓存的组件实例对象
	// Lifecycle    Lifecycle             // 用于组件的生命周期控制
	Registration ComponentRegistration // 注册信息
	Classes      []string              // Registration.Class 解析后的内容
	Aliases      []string              // Registration.Alias 解析后的内容
}

////////////////////////////////////////////////////////////////////////////////

// ComponentManager 代表组件管理器
type ComponentManager interface {
	Add(h *ComponentHolder) error
	Register(reg *ComponentRegistration) (*ComponentHolder, error)
	GetByName(name string) (*ComponentHolder, error)
	ListByClass(className string) ([]*ComponentHolder, error)
	ListAll() []*ComponentHolder
}

// NewComponentManager 新建 ComponentManager 对象
func NewComponentManager(mode safe.Mode) ComponentManager {
	cm := &componentManagerImpl{}
	cm.init(mode)
	return cm
}

////////////////////////////////////////////////////////////////////////////////

type componentManagerImpl struct {
	mode  safe.Lock
	list  []*ComponentHolder
	table map[string]*ComponentHolder
}

func (inst *componentManagerImpl) init(mode safe.Mode) {
	if mode == nil {
		mode = safe.Safe()
	}
	inst.mode = mode.NewLock()
	inst.list = make([]*ComponentHolder, 0)
	inst.table = make(map[string]*ComponentHolder)
}

func (inst *componentManagerImpl) addHolder(h *ComponentHolder) error {
	if h == nil {
		return fmt.Errorf("param:ComponentHolder is nil")
	}
	name := h.Registration.Name
	h.Aliases = inst.parseNameList(h.Registration.Alias + " " + name)
	h.Classes = inst.parseNameList(h.Registration.Class)
	inst.list = append(inst.list, h)
	ids := h.Aliases
	// ids = append(ids, )
	for _, id := range ids {
		old := inst.table[id]
		if old != nil {
			return fmt.Errorf("the id (name|alias) of components are duplicate: %s", id)
		}
		inst.table[id] = h
	}
	return nil
}

func (inst *componentManagerImpl) parseNameList(str string) []string {
	const (
		sep  = "\n"
		sep1 = "\r"
		sep2 = "\t"
		sep3 = string(' ')
	)
	str = strings.ReplaceAll(str, sep1, sep)
	str = strings.ReplaceAll(str, sep2, sep)
	str = strings.ReplaceAll(str, sep3, sep)
	src := strings.Split(str, sep)
	dst := make([]string, 0)
	for _, item := range src {
		// item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *componentManagerImpl) Add(h *ComponentHolder) error {
	inst.lock()
	defer inst.unlock()
	return inst.addHolder(h)
}

func (inst *componentManagerImpl) Register(reg *ComponentRegistration) (*ComponentHolder, error) {

	if reg == nil {
		return nil, fmt.Errorf("component registration is nil")
	}

	inst.lock()
	defer inst.unlock()

	h := &ComponentHolder{}
	h.Registration = *reg
	err := inst.addHolder(h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (inst *componentManagerImpl) GetByName(name string) (*ComponentHolder, error) {

	inst.lock()
	defer inst.unlock()

	h := inst.table["#"+name]
	if h == nil {
		return nil, fmt.Errorf("cannot find component with name (or alias): %s", name)
	}
	return h, nil
}

func (inst *componentManagerImpl) ListByClass(className string) ([]*ComponentHolder, error) {
	inst.lock()
	defer inst.unlock()
	dst := make([]*ComponentHolder, 0)
	src := inst.list
	for _, item := range src {
		if inst.isClassOf(className, item) {
			dst = append(dst, item)
		}
	}
	return dst, nil
}

func (inst *componentManagerImpl) isClassOf(className string, h *ComponentHolder) bool {
	if className == "" {
		return false
	}
	list := h.Classes
	for _, haveClass := range list {
		if haveClass == className {
			return true
		}
	}
	return false
}

func (inst *componentManagerImpl) ListAll() []*ComponentHolder {

	inst.lock()
	defer inst.unlock()

	src := inst.list
	dst := make([]*ComponentHolder, 0)
	dst = append(dst, src...)
	return dst
}

func (inst *componentManagerImpl) lock() {
	inst.mode.Lock()
}

func (inst *componentManagerImpl) unlock() {
	inst.mode.Unlock()
}

////////////////////////////////////////////////////////////////////////////////
