package repositories

import "github.com/starter-go/base/safe"

// ComponentLifecycleManager ...
type ComponentLifecycleManager struct {
	compManager ComponentManager
	lifeManager LifecycleManager
	loaded      bool
}

// Init ...
func (inst *ComponentLifecycleManager) Init(mode safe.Mode) {
	inst.compManager = NewComponentManager(mode)
	inst.lifeManager = NewLifecycleManager(mode)
}

// GetComponents ...
func (inst *ComponentLifecycleManager) GetComponents() ComponentManager {
	return inst.compManager
}

// GetLifecycles ...
func (inst *ComponentLifecycleManager) GetLifecycles() LifecycleManager {
	return inst.lifeManager
}

// Load 把组件的生命周期加载到 LifecycleManager
func (inst *ComponentLifecycleManager) Load() {
	if inst.loaded {
		return // invoke only once
	}
	inst.loaded = true
	src := inst.compManager.ListAll()
	for _, h := range src {
		com := h.Component
		inst.loadCom(com)
	}
}

func (inst *ComponentLifecycleManager) loadCom(com any) {
	inst.lifeManager.AddComponent(com)
}
