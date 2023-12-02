package main4gitlib

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p115564418d_components_ExampleComponentRegistry{})
    inst.register(&p115564418d_components_LifecyclelifecycleExampleComRegistry{})
    inst.register(&pa6edb103eb_com4system_LibComReg{})
    inst.register(&peee8217021_implements_LibAgentImpl{})
    inst.register(&peee8217021_implements_SystemContextLoaderImpl{})


    return nil
}
