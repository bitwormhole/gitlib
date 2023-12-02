package test4gitlib
import (
    pa06da1b0f "github.com/bitwormhole/gitlib/git/repositories"
    p4fac42bf2 "github.com/bitwormhole/gitlib/src/test/golang/code/cases"
     "github.com/starter-go/application"
)

// type p4fac42bf2.SystemContextTest in package:github.com/bitwormhole/gitlib/src/test/golang/code/cases
//
// id:com-4fac42bf2f307abb-cases-SystemContextTest
// class:
// alias:
// scope:singleton
//
type p4fac42bf2f_cases_SystemContextTest struct {
}

func (inst* p4fac42bf2f_cases_SystemContextTest) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4fac42bf2f307abb-cases-SystemContextTest"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4fac42bf2f_cases_SystemContextTest) new() any {
    return &p4fac42bf2.SystemContextTest{}
}

func (inst* p4fac42bf2f_cases_SystemContextTest) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4fac42bf2.SystemContextTest)
	nop(ie, com)

	
    com.LibAgent = inst.getLibAgent(ie)


    return nil
}


func (inst*p4fac42bf2f_cases_SystemContextTest) getLibAgent(ie application.InjectionExt)pa06da1b0f.LibAgent{
    return ie.GetComponent("#alias-a06da1b0f12870119f837ccacb2eabeb-LibAgent").(pa06da1b0f.LibAgent)
}


