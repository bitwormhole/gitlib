package components

import (
	"github.com/bitwormhole/gitlib/git/repositories"
)

// ExampleComponentRegistry ...
type ExampleComponentRegistry struct {

	//starter:component
	_as func(repositories.ComponentRegistry) //starter:as(".")

}

func (inst *ExampleComponentRegistry) _impl() repositories.ComponentRegistry { return inst }

// ListRegistrations ...
func (inst *ExampleComponentRegistry) ListRegistrations() []*repositories.ComponentRegistration {
	r1 := &repositories.ComponentRegistration{
		Enabled:         true,
		OnInitForSystem: inst.create,
	}
	return []*repositories.ComponentRegistration{r1}
}

func (inst *ExampleComponentRegistry) create(ctx *repositories.SystemContext) (any, error) {
	com := &exampleComponent{context: ctx}
	// ctx.Facade = com
	return com, nil
}

////////////////////////////////////////////////////////////////////////////////

type exampleComponent struct {
	context *repositories.SystemContext
}

func (inst *exampleComponent) _impl() any {
	return inst
}
