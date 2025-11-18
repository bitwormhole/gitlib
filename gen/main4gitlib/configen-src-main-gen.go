package main4gitlib
import (
    pa849a78ce "github.com/bitwormhole/gitlib/src/main/golang/u4go"
     "github.com/starter-go/application"
)

// type pa849a78ce.Demo in package:github.com/bitwormhole/gitlib/src/main/golang/u4go
//
// id:com-a849a78ceff349fe-u4go-Demo
// class:
// alias:
// scope:singleton
//
type pa849a78cef_u4go_Demo struct {
}

func (inst* pa849a78cef_u4go_Demo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a849a78ceff349fe-u4go-Demo"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa849a78cef_u4go_Demo) new() any {
    return &pa849a78ce.Demo{}
}

func (inst* pa849a78cef_u4go_Demo) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa849a78ce.Demo)
	nop(ie, com)

	


    return nil
}


