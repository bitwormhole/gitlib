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
    inst.register(&p2b096b8823_com4worktree_WorktreeFacadeRegistry{})
    inst.register(&p4bc7581ec5_com4repository_RepoConfigLoaderRegistry{})
    inst.register(&p4bc7581ec5_com4repository_RepoFacadeRegistry{})
    inst.register(&p64aca10263_com4user_UserConfigLoaderRegistry{})
    inst.register(&p64aca10263_com4user_UserFacadeRegistry{})
    inst.register(&p66e7b268c7_com4submodule_SubmoduleFacadeRegistry{})
    inst.register(&p8303bffc6e_com4session_SessionFacadeRegistry{})
    inst.register(&pa6edb103eb_com4system_LibComReg{})
    inst.register(&pa6edb103eb_com4system_RepositoryFinderRegistry{})
    inst.register(&pa6edb103eb_com4system_RepositoryLocatorRegistry{})
    inst.register(&pa6edb103eb_com4system_SystemConfigLoaderRegistry{})
    inst.register(&pa6edb103eb_com4system_SystemFacadeRegistry{})
    inst.register(&peee8217021_implements_GitlibAgentImpl{})
    inst.register(&peee8217021_implements_LibAgentImpl{})
    inst.register(&pf2f602615f_loaders_RepositoryContextLoaderImpl{})
    inst.register(&pf2f602615f_loaders_SessionContextLoaderImpl{})
    inst.register(&pf2f602615f_loaders_SubmoduleContextLoaderImpl{})
    inst.register(&pf2f602615f_loaders_SystemContextLoaderImpl{})
    inst.register(&pf2f602615f_loaders_UserContextLoaderImpl{})
    inst.register(&pf2f602615f_loaders_WorktreeContextLoaderImpl{})


    return nil
}
