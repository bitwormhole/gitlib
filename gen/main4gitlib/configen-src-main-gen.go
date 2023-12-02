package main4gitlib
import (
    pa06da1b0f "github.com/bitwormhole/gitlib/git/repositories"
    peee821702 "github.com/bitwormhole/gitlib/src/main/golang/code/implements"
    p115564418 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components"
    pa6edb103e "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system"
     "github.com/starter-go/application"
)

// type peee821702.LibAgentImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements
//
// id:com-eee8217021100052-implements-LibAgentImpl
// class:
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-LibAgent
// scope:singleton
//
type peee8217021_implements_LibAgentImpl struct {
}

func (inst* peee8217021_implements_LibAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eee8217021100052-implements-LibAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-LibAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peee8217021_implements_LibAgentImpl) new() any {
    return &peee821702.LibAgentImpl{}
}

func (inst* peee8217021_implements_LibAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peee821702.LibAgentImpl)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*peee8217021_implements_LibAgentImpl) getLoader(ie application.InjectionExt)pa06da1b0f.SystemContextLoader{
    return ie.GetComponent("#alias-a06da1b0f12870119f837ccacb2eabeb-SystemContextLoader").(pa06da1b0f.SystemContextLoader)
}



// type peee821702.SystemContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements
//
// id:com-eee8217021100052-implements-SystemContextLoaderImpl
// class:
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-SystemContextLoader
// scope:singleton
//
type peee8217021_implements_SystemContextLoaderImpl struct {
}

func (inst* peee8217021_implements_SystemContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eee8217021100052-implements-SystemContextLoaderImpl"
	r.Classes = ""
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-SystemContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peee8217021_implements_SystemContextLoaderImpl) new() any {
    return &peee821702.SystemContextLoaderImpl{}
}

func (inst* peee8217021_implements_SystemContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peee821702.SystemContextLoaderImpl)
	nop(ie, com)

	
    com.AllComponents = inst.getAllComponents(ie)
    com.UseSafeMode = inst.getUseSafeMode(ie)


    return nil
}


func (inst*peee8217021_implements_SystemContextLoaderImpl) getAllComponents(ie application.InjectionExt)[]pa06da1b0f.ComponentRegistry{
    dst := make([]pa06da1b0f.ComponentRegistry, 0)
    src := ie.ListComponents(".class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry")
    for _, item1 := range src {
        item2 := item1.(pa06da1b0f.ComponentRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*peee8217021_implements_SystemContextLoaderImpl) getUseSafeMode(ie application.InjectionExt)bool{
    return ie.GetBool("${git.threads.use-safe-mode}")
}



// type p115564418.ExampleComponentRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components
//
// id:com-115564418d6dce6c-components-ExampleComponentRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p115564418d_components_ExampleComponentRegistry struct {
}

func (inst* p115564418d_components_ExampleComponentRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-115564418d6dce6c-components-ExampleComponentRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p115564418d_components_ExampleComponentRegistry) new() any {
    return &p115564418.ExampleComponentRegistry{}
}

func (inst* p115564418d_components_ExampleComponentRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p115564418.ExampleComponentRegistry)
	nop(ie, com)

	


    return nil
}



// type p115564418.LifecyclelifecycleExampleComRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components
//
// id:com-115564418d6dce6c-components-LifecyclelifecycleExampleComRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p115564418d_components_LifecyclelifecycleExampleComRegistry struct {
}

func (inst* p115564418d_components_LifecyclelifecycleExampleComRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-115564418d6dce6c-components-LifecyclelifecycleExampleComRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p115564418d_components_LifecyclelifecycleExampleComRegistry) new() any {
    return &p115564418.LifecyclelifecycleExampleComRegistry{}
}

func (inst* p115564418d_components_LifecyclelifecycleExampleComRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p115564418.LifecyclelifecycleExampleComRegistry)
	nop(ie, com)

	


    return nil
}



// type pa6edb103e.LibComReg in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system
//
// id:com-a6edb103ebb3bc50-com4system-LibComReg
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type pa6edb103eb_com4system_LibComReg struct {
}

func (inst* pa6edb103eb_com4system_LibComReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a6edb103ebb3bc50-com4system-LibComReg"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa6edb103eb_com4system_LibComReg) new() any {
    return &pa6edb103e.LibComReg{}
}

func (inst* pa6edb103eb_com4system_LibComReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa6edb103e.LibComReg)
	nop(ie, com)

	


    return nil
}


