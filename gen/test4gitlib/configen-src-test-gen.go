package test4gitlib
import (
    pa06da1b0f "github.com/bitwormhole/gitlib/git/repositories"
    p4fac42bf2 "github.com/bitwormhole/gitlib/src/test/golang/code/cases"
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
     "github.com/starter-go/application"
)

// type p4fac42bf2.RepositoryPathProviderImpl in package:github.com/bitwormhole/gitlib/src/test/golang/code/cases
//
// id:com-4fac42bf2f307abb-cases-RepositoryPathProviderImpl
// class:
// alias:alias-4fac42bf2f307abbe185e1941901afc0-RepositoryPathProvider
// scope:singleton
//
type p4fac42bf2f_cases_RepositoryPathProviderImpl struct {
}

func (inst* p4fac42bf2f_cases_RepositoryPathProviderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4fac42bf2f307abb-cases-RepositoryPathProviderImpl"
	r.Classes = ""
	r.Aliases = "alias-4fac42bf2f307abbe185e1941901afc0-RepositoryPathProvider"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4fac42bf2f_cases_RepositoryPathProviderImpl) new() any {
    return &p4fac42bf2.RepositoryPathProviderImpl{}
}

func (inst* p4fac42bf2f_cases_RepositoryPathProviderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4fac42bf2.RepositoryPathProviderImpl)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p4fac42bf2f_cases_RepositoryPathProviderImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p4fac42bf2f_cases_RepositoryPathProviderImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p4fac42bf2.ContextLoadersTest in package:github.com/bitwormhole/gitlib/src/test/golang/code/cases
//
// id:com-4fac42bf2f307abb-cases-ContextLoadersTest
// class:
// alias:
// scope:singleton
//
type p4fac42bf2f_cases_ContextLoadersTest struct {
}

func (inst* p4fac42bf2f_cases_ContextLoadersTest) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4fac42bf2f307abb-cases-ContextLoadersTest"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4fac42bf2f_cases_ContextLoadersTest) new() any {
    return &p4fac42bf2.ContextLoadersTest{}
}

func (inst* p4fac42bf2f_cases_ContextLoadersTest) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4fac42bf2.ContextLoadersTest)
	nop(ie, com)

	
    com.LibAgent = inst.getLibAgent(ie)
    com.Paths = inst.getPaths(ie)


    return nil
}


func (inst*p4fac42bf2f_cases_ContextLoadersTest) getLibAgent(ie application.InjectionExt)pa06da1b0f.LibAgent{
    return ie.GetComponent("#alias-a06da1b0f12870119f837ccacb2eabeb-LibAgent").(pa06da1b0f.LibAgent)
}


func (inst*p4fac42bf2f_cases_ContextLoadersTest) getPaths(ie application.InjectionExt)p4fac42bf2.RepositoryPathProvider{
    return ie.GetComponent("#alias-4fac42bf2f307abbe185e1941901afc0-RepositoryPathProvider").(p4fac42bf2.RepositoryPathProvider)
}



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


