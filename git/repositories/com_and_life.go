package repositories

import (
	"github.com/starter-go/base/safe"
	"github.com/starter-go/vlog"
)

// ComponentLifecycleManager ...
type ComponentLifecycleManager struct {
	compManager ComponentManager
	lifeManager LifecycleManager
}

// GetComponents ...
func (inst *ComponentLifecycleManager) GetComponents() ComponentManager {
	return inst.compManager
}

// GetLifecycles ...
func (inst *ComponentLifecycleManager) GetLifecycles() LifecycleManager {
	return inst.lifeManager
}

// Init ...
func (inst *ComponentLifecycleManager) Init(items []*ComponentRegistration, mode safe.Mode) *ComponentLifecycleManager {
	if inst.compManager != nil || inst.lifeManager != nil {
		return inst // 这个函数只能执行一次
	}
	inst.compManager = NewComponentManager(mode)
	inst.lifeManager = NewLifecycleManager(mode)
	for _, item := range items {
		inst.compManager.Register(item)
	}
	return inst
}

// CreateItems  创建各个组件的实例，并把组件的生命周期添加到 LifecycleManager
func (inst *ComponentLifecycleManager) CreateItems(callback func(h *ComponentHolder) error) {
	hlist := inst.GetComponents().ListAll()
	for _, holder := range hlist {
		if holder.Component == nil {
			err := callback(holder)
			if err != nil {
				vlog.Warn(err.Error())
			}
			if holder.Component == nil {
				continue
			}
		}
		inst.addToLifeManager(holder.Component)
	}
}

func (inst *ComponentLifecycleManager) addToLifeManager(com any) {
	inst.lifeManager.AddComponent(com)
}
