package test4gitlib
import (
    p4268ba53a "github.com/bitwormhole/gitlib/src/test/golang/u4go"
     "github.com/starter-go/application"
)

// type p4268ba53a.Demo in package:github.com/bitwormhole/gitlib/src/test/golang/u4go
//
// id:com-4268ba53a5adeb35-u4go-Demo
// class:
// alias:
// scope:singleton
//
type p4268ba53a5_u4go_Demo struct {
}

func (inst* p4268ba53a5_u4go_Demo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4268ba53a5adeb35-u4go-Demo"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4268ba53a5_u4go_Demo) new() any {
    return &p4268ba53a.Demo{}
}

func (inst* p4268ba53a5_u4go_Demo) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4268ba53a.Demo)
	nop(ie, com)

	


    return nil
}


