package components

import (
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/vlog"
)

// LifecyclelifecycleExampleComRegistry ...
type LifecyclelifecycleExampleComRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *LifecyclelifecycleExampleComRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *LifecyclelifecycleExampleComRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *LifecyclelifecycleExampleComRegistry) create(ctx *repositories.SystemContext) (any, error) {
	com := &lifecycleExampleCom{context: ctx}
	// ctx.Facade = com
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type lifecycleExampleCom struct {
	context *repositories.SystemContext
}

func (inst *lifecycleExampleCom) _impl() repositories.LifecycleRegistry {
	return inst
}

func (inst *lifecycleExampleCom) Lifecycle() *repositories.Lifecycle {
	return &repositories.Lifecycle{
		OnCreate: inst.onCreate,

		OnStartPre:  inst.onStart1,
		OnStart:     inst.onStart2,
		OnStartPost: inst.onStart3,

		OnStopPre:  inst.onStop1,
		OnStop:     inst.onStop2,
		OnStopPost: inst.onStop3,

		OnDestroy: inst.onDestroy,
	}
}

func (inst *lifecycleExampleCom) onCreate() error {
	inst.log("onCreate")
	return nil
	// return inst.e()
}

func (inst *lifecycleExampleCom) onStart1() error {
	inst.log("on_start_pre")
	// return nil
	return inst.e()
}
func (inst *lifecycleExampleCom) onStart2() error {
	inst.log("on_start")
	// return nil
	return inst.e()
}
func (inst *lifecycleExampleCom) onStart3() error {
	inst.log("on_start_post")
	// return nil
	return inst.e()
}

func (inst *lifecycleExampleCom) onStop1() error {
	inst.log("on_stop_pre")
	// return nil
	return inst.e()
}

func (inst *lifecycleExampleCom) onStop2() error {
	inst.log("on_stop")
	// return nil
	return inst.e()
}

func (inst *lifecycleExampleCom) onStop3() error {
	inst.log("on_stop_post")
	// return nil
	return inst.e()
}

func (inst *lifecycleExampleCom) onDestroy() error {
	inst.log("onDestroy")
	// return nil
	return inst.e()
}

func (inst *lifecycleExampleCom) e() error {
	// return fmt.Errorf("error for test")
	return nil
}

func (inst *lifecycleExampleCom) log(msg string) {
	vlog.Warn("lifecycleExampleCom.%s()", msg)
}
