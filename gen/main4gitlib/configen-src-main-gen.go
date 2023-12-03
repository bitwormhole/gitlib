package main4gitlib
import (
    pa06da1b0f "github.com/bitwormhole/gitlib/git/repositories"
    peee821702 "github.com/bitwormhole/gitlib/src/main/golang/code/implements"
    p115564418 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components"
    p4bc7581ec "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4repository"
    p8303bffc6 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4session"
    p66e7b268c "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4submodule"
    pa6edb103e "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system"
    p64aca1026 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4user"
    p2b096b882 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4worktree"
    pf2f602615 "github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders"
    p0d2a11d16 "github.com/starter-go/afs"
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



// type p4bc7581ec.RepoConfigLoaderRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4repository
//
// id:com-4bc7581ec55229b0-com4repository-RepoConfigLoaderRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p4bc7581ec5_com4repository_RepoConfigLoaderRegistry struct {
}

func (inst* p4bc7581ec5_com4repository_RepoConfigLoaderRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4bc7581ec55229b0-com4repository-RepoConfigLoaderRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4bc7581ec5_com4repository_RepoConfigLoaderRegistry) new() any {
    return &p4bc7581ec.RepoConfigLoaderRegistry{}
}

func (inst* p4bc7581ec5_com4repository_RepoConfigLoaderRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4bc7581ec.RepoConfigLoaderRegistry)
	nop(ie, com)

	


    return nil
}



// type p4bc7581ec.RepoFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4repository
//
// id:com-4bc7581ec55229b0-com4repository-RepoFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p4bc7581ec5_com4repository_RepoFacadeRegistry struct {
}

func (inst* p4bc7581ec5_com4repository_RepoFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4bc7581ec55229b0-com4repository-RepoFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4bc7581ec5_com4repository_RepoFacadeRegistry) new() any {
    return &p4bc7581ec.RepoFacadeRegistry{}
}

func (inst* p4bc7581ec5_com4repository_RepoFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4bc7581ec.RepoFacadeRegistry)
	nop(ie, com)

	


    return nil
}



// type p8303bffc6.SessionFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4session
//
// id:com-8303bffc6ea5c092-com4session-SessionFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p8303bffc6e_com4session_SessionFacadeRegistry struct {
}

func (inst* p8303bffc6e_com4session_SessionFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8303bffc6ea5c092-com4session-SessionFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8303bffc6e_com4session_SessionFacadeRegistry) new() any {
    return &p8303bffc6.SessionFacadeRegistry{}
}

func (inst* p8303bffc6e_com4session_SessionFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8303bffc6.SessionFacadeRegistry)
	nop(ie, com)

	


    return nil
}



// type p66e7b268c.SubmoduleFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4submodule
//
// id:com-66e7b268c7ef713a-com4submodule-SubmoduleFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p66e7b268c7_com4submodule_SubmoduleFacadeRegistry struct {
}

func (inst* p66e7b268c7_com4submodule_SubmoduleFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-66e7b268c7ef713a-com4submodule-SubmoduleFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p66e7b268c7_com4submodule_SubmoduleFacadeRegistry) new() any {
    return &p66e7b268c.SubmoduleFacadeRegistry{}
}

func (inst* p66e7b268c7_com4submodule_SubmoduleFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p66e7b268c.SubmoduleFacadeRegistry)
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



// type pa6edb103e.RepositoryFinderRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system
//
// id:com-a6edb103ebb3bc50-com4system-RepositoryFinderRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type pa6edb103eb_com4system_RepositoryFinderRegistry struct {
}

func (inst* pa6edb103eb_com4system_RepositoryFinderRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a6edb103ebb3bc50-com4system-RepositoryFinderRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa6edb103eb_com4system_RepositoryFinderRegistry) new() any {
    return &pa6edb103e.RepositoryFinderRegistry{}
}

func (inst* pa6edb103eb_com4system_RepositoryFinderRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa6edb103e.RepositoryFinderRegistry)
	nop(ie, com)

	


    return nil
}



// type pa6edb103e.RepositoryLocatorRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system
//
// id:com-a6edb103ebb3bc50-com4system-RepositoryLocatorRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type pa6edb103eb_com4system_RepositoryLocatorRegistry struct {
}

func (inst* pa6edb103eb_com4system_RepositoryLocatorRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a6edb103ebb3bc50-com4system-RepositoryLocatorRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa6edb103eb_com4system_RepositoryLocatorRegistry) new() any {
    return &pa6edb103e.RepositoryLocatorRegistry{}
}

func (inst* pa6edb103eb_com4system_RepositoryLocatorRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa6edb103e.RepositoryLocatorRegistry)
	nop(ie, com)

	


    return nil
}



// type pa6edb103e.SystemConfigLoaderRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system
//
// id:com-a6edb103ebb3bc50-com4system-SystemConfigLoaderRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type pa6edb103eb_com4system_SystemConfigLoaderRegistry struct {
}

func (inst* pa6edb103eb_com4system_SystemConfigLoaderRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a6edb103ebb3bc50-com4system-SystemConfigLoaderRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa6edb103eb_com4system_SystemConfigLoaderRegistry) new() any {
    return &pa6edb103e.SystemConfigLoaderRegistry{}
}

func (inst* pa6edb103eb_com4system_SystemConfigLoaderRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa6edb103e.SystemConfigLoaderRegistry)
	nop(ie, com)

	


    return nil
}



// type pa6edb103e.SystemFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4system
//
// id:com-a6edb103ebb3bc50-com4system-SystemFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type pa6edb103eb_com4system_SystemFacadeRegistry struct {
}

func (inst* pa6edb103eb_com4system_SystemFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a6edb103ebb3bc50-com4system-SystemFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa6edb103eb_com4system_SystemFacadeRegistry) new() any {
    return &pa6edb103e.SystemFacadeRegistry{}
}

func (inst* pa6edb103eb_com4system_SystemFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa6edb103e.SystemFacadeRegistry)
	nop(ie, com)

	


    return nil
}



// type p64aca1026.UserConfigLoaderRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4user
//
// id:com-64aca10263664a04-com4user-UserConfigLoaderRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p64aca10263_com4user_UserConfigLoaderRegistry struct {
}

func (inst* p64aca10263_com4user_UserConfigLoaderRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-64aca10263664a04-com4user-UserConfigLoaderRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p64aca10263_com4user_UserConfigLoaderRegistry) new() any {
    return &p64aca1026.UserConfigLoaderRegistry{}
}

func (inst* p64aca10263_com4user_UserConfigLoaderRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p64aca1026.UserConfigLoaderRegistry)
	nop(ie, com)

	


    return nil
}



// type p64aca1026.UserFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4user
//
// id:com-64aca10263664a04-com4user-UserFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p64aca10263_com4user_UserFacadeRegistry struct {
}

func (inst* p64aca10263_com4user_UserFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-64aca10263664a04-com4user-UserFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p64aca10263_com4user_UserFacadeRegistry) new() any {
    return &p64aca1026.UserFacadeRegistry{}
}

func (inst* p64aca10263_com4user_UserFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p64aca1026.UserFacadeRegistry)
	nop(ie, com)

	


    return nil
}



// type p2b096b882.WorktreeFacadeRegistry in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/components/com4worktree
//
// id:com-2b096b882319c012-com4worktree-WorktreeFacadeRegistry
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:
// scope:singleton
//
type p2b096b8823_com4worktree_WorktreeFacadeRegistry struct {
}

func (inst* p2b096b8823_com4worktree_WorktreeFacadeRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2b096b882319c012-com4worktree-WorktreeFacadeRegistry"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2b096b8823_com4worktree_WorktreeFacadeRegistry) new() any {
    return &p2b096b882.WorktreeFacadeRegistry{}
}

func (inst* p2b096b8823_com4worktree_WorktreeFacadeRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2b096b882.WorktreeFacadeRegistry)
	nop(ie, com)

	


    return nil
}



// type pf2f602615.RepositoryContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-RepositoryContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-RepositoryContextLoader
// scope:singleton
//
type pf2f602615f_loaders_RepositoryContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_RepositoryContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-RepositoryContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-RepositoryContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_RepositoryContextLoaderImpl) new() any {
    return &pf2f602615.RepositoryContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_RepositoryContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.RepositoryContextLoaderImpl)
	nop(ie, com)

	


    return nil
}



// type pf2f602615.SessionContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-SessionContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-SessionContextLoader
// scope:singleton
//
type pf2f602615f_loaders_SessionContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_SessionContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-SessionContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-SessionContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_SessionContextLoaderImpl) new() any {
    return &pf2f602615.SessionContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_SessionContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.SessionContextLoaderImpl)
	nop(ie, com)

	


    return nil
}



// type pf2f602615.SubmoduleContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-SubmoduleContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-SubmoduleContextLoader
// scope:singleton
//
type pf2f602615f_loaders_SubmoduleContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_SubmoduleContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-SubmoduleContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-SubmoduleContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_SubmoduleContextLoaderImpl) new() any {
    return &pf2f602615.SubmoduleContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_SubmoduleContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.SubmoduleContextLoaderImpl)
	nop(ie, com)

	


    return nil
}



// type pf2f602615.SystemContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-SystemContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-SystemContextLoader
// scope:singleton
//
type pf2f602615f_loaders_SystemContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_SystemContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-SystemContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-SystemContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_SystemContextLoaderImpl) new() any {
    return &pf2f602615.SystemContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_SystemContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.SystemContextLoaderImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.AllComponents = inst.getAllComponents(ie)
    com.UseSafeMode = inst.getUseSafeMode(ie)


    return nil
}


func (inst*pf2f602615f_loaders_SystemContextLoaderImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*pf2f602615f_loaders_SystemContextLoaderImpl) getAllComponents(ie application.InjectionExt)[]pa06da1b0f.ComponentRegistry{
    dst := make([]pa06da1b0f.ComponentRegistry, 0)
    src := ie.ListComponents(".class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry")
    for _, item1 := range src {
        item2 := item1.(pa06da1b0f.ComponentRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pf2f602615f_loaders_SystemContextLoaderImpl) getUseSafeMode(ie application.InjectionExt)bool{
    return ie.GetBool("${git.threads.use-safe-mode}")
}



// type pf2f602615.UserContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-UserContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-UserContextLoader
// scope:singleton
//
type pf2f602615f_loaders_UserContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_UserContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-UserContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-UserContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_UserContextLoaderImpl) new() any {
    return &pf2f602615.UserContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_UserContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.UserContextLoaderImpl)
	nop(ie, com)

	


    return nil
}



// type pf2f602615.WorktreeContextLoaderImpl in package:github.com/bitwormhole/gitlib/src/main/golang/code/implements/loaders
//
// id:com-f2f602615f596b42-loaders-WorktreeContextLoaderImpl
// class:class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry
// alias:alias-a06da1b0f12870119f837ccacb2eabeb-WorktreeContextLoader
// scope:singleton
//
type pf2f602615f_loaders_WorktreeContextLoaderImpl struct {
}

func (inst* pf2f602615f_loaders_WorktreeContextLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f2f602615f596b42-loaders-WorktreeContextLoaderImpl"
	r.Classes = "class-a06da1b0f12870119f837ccacb2eabeb-ComponentRegistry"
	r.Aliases = "alias-a06da1b0f12870119f837ccacb2eabeb-WorktreeContextLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf2f602615f_loaders_WorktreeContextLoaderImpl) new() any {
    return &pf2f602615.WorktreeContextLoaderImpl{}
}

func (inst* pf2f602615f_loaders_WorktreeContextLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf2f602615.WorktreeContextLoaderImpl)
	nop(ie, com)

	


    return nil
}


