package test4gitlib
import (
    p4268ba53a "github.com/bitwormhole/gitlib/src/test/golang/u4go"
     "github.com/starter-go/application"
)

// type p4268ba53a.GitRepoLocatorTest in package:github.com/bitwormhole/gitlib/src/test/golang/u4go
//
// id:com-4268ba53a5adeb35-u4go-GitRepoLocatorTest
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p4268ba53a5_u4go_GitRepoLocatorTest struct {
}

func (inst* p4268ba53a5_u4go_GitRepoLocatorTest) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4268ba53a5adeb35-u4go-GitRepoLocatorTest"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4268ba53a5_u4go_GitRepoLocatorTest) new() any {
    return &p4268ba53a.GitRepoLocatorTest{}
}

func (inst* p4268ba53a5_u4go_GitRepoLocatorTest) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4268ba53a.GitRepoLocatorTest)
	nop(ie, com)

	


    return nil
}



// type p4268ba53a.Demo in package:github.com/bitwormhole/gitlib/src/test/golang/u4go
//
// id:com-4268ba53a5adeb35-u4go-Demo
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p4268ba53a5_u4go_Demo struct {
}

func (inst* p4268ba53a5_u4go_Demo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4268ba53a5adeb35-u4go-Demo"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
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


