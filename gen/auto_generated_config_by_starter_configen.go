// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	files0x00707a "github.com/bitwormhole/gitlib/git/files"
	repository0x5aaf5a "github.com/bitwormhole/gitlib/git/repository"
	support0x074feb "github.com/bitwormhole/gitlib/git/support"
	localfilesys0x6be3ff "github.com/bitwormhole/gitlib/git/support/localfilesys"
	config0x71b4a2 "github.com/bitwormhole/gitlib/git/support/localfilesys/config"
	head0xb6393b "github.com/bitwormhole/gitlib/git/support/localfilesys/head"
	index0x3eb559 "github.com/bitwormhole/gitlib/git/support/localfilesys/index"
	modules0xe8c0dc "github.com/bitwormhole/gitlib/git/support/localfilesys/modules"
	objects0x11508a "github.com/bitwormhole/gitlib/git/support/localfilesys/objects"
	refs0x4e5472 "github.com/bitwormhole/gitlib/git/support/localfilesys/refs"
	views0x7d154b "github.com/bitwormhole/gitlib/git/support/localfilesys/views"
	worktrees0xa9c0a9 "github.com/bitwormhole/gitlib/git/support/localfilesys/worktrees"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-config0x71b4a2.LocalGitConfigFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com0-config0x71b4a2.LocalGitConfigFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalGitConfigFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-head0xb6393b.GitHeadFileFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com1-head0xb6393b.GitHeadFileFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitHeadFileFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-index0x3eb559.GitIndexFileFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com2-index0x3eb559.GitIndexFileFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitIndexFileFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-modules0xe8c0dc.GitModulesDirFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com3-modules0xe8c0dc.GitModulesDirFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitModulesDirFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-objects0x11508a.LocalGitObjectsFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com4-objects0x11508a.LocalGitObjectsFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalGitObjectsFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-refs0x4e5472.LocalGitRefsFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com5-refs0x4e5472.LocalGitRefsFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalGitRefsFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-localfilesys0x6be3ff.LocalRepoDriver
	cominfobuilder.Next()
	cominfobuilder.ID("com6-localfilesys0x6be3ff.LocalRepoDriver").Class("git-repository-driver").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepoDriver{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-local-repository-factory
	cominfobuilder.Next()
	cominfobuilder.ID("git-local-repository-factory").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepoFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-local-repository-layout
	cominfobuilder.Next()
	cominfobuilder.ID("git-local-repository-layout").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepoLayout{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-local-repository-locator
	cominfobuilder.Next()
	cominfobuilder.ID("git-local-repository-locator").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepoLocator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com10-views0x7d154b.CoreDirFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com10-views0x7d154b.CoreDirFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCoreDirFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com11-views0x7d154b.ShellDirFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com11-views0x7d154b.ShellDirFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComShellDirFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com12-views0x7d154b.RepoViewFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com12-views0x7d154b.RepoViewFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepoViewFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com13-views0x7d154b.WorkingDirFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com13-views0x7d154b.WorkingDirFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWorkingDirFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com14-worktrees0xa9c0a9.GitWorktreesDirFactory
	cominfobuilder.Next()
	cominfobuilder.ID("com14-worktrees0xa9c0a9.GitWorktreesDirFactory").Class("git-local-element-factory").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitWorktreesDirFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-repository-manager
	cominfobuilder.Next()
	cominfobuilder.ID("git-repository-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitRepoManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalGitConfigFactory : the factory of component: com0-config0x71b4a2.LocalGitConfigFactory
type comFactory4pComLocalGitConfigFactory struct {

    mPrototype * config0x71b4a2.LocalGitConfigFactory

	

}

func (inst * comFactory4pComLocalGitConfigFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalGitConfigFactory) newObject() * config0x71b4a2.LocalGitConfigFactory {
	return & config0x71b4a2.LocalGitConfigFactory {}
}

func (inst * comFactory4pComLocalGitConfigFactory) castObject(instance application.ComponentInstance) * config0x71b4a2.LocalGitConfigFactory {
	return instance.Get().(*config0x71b4a2.LocalGitConfigFactory)
}

func (inst * comFactory4pComLocalGitConfigFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalGitConfigFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalGitConfigFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalGitConfigFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitConfigFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitConfigFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitHeadFileFactory : the factory of component: com1-head0xb6393b.GitHeadFileFactory
type comFactory4pComGitHeadFileFactory struct {

    mPrototype * head0xb6393b.GitHeadFileFactory

	

}

func (inst * comFactory4pComGitHeadFileFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitHeadFileFactory) newObject() * head0xb6393b.GitHeadFileFactory {
	return & head0xb6393b.GitHeadFileFactory {}
}

func (inst * comFactory4pComGitHeadFileFactory) castObject(instance application.ComponentInstance) * head0xb6393b.GitHeadFileFactory {
	return instance.Get().(*head0xb6393b.GitHeadFileFactory)
}

func (inst * comFactory4pComGitHeadFileFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitHeadFileFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitHeadFileFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitHeadFileFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitHeadFileFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitHeadFileFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitIndexFileFactory : the factory of component: com2-index0x3eb559.GitIndexFileFactory
type comFactory4pComGitIndexFileFactory struct {

    mPrototype * index0x3eb559.GitIndexFileFactory

	

}

func (inst * comFactory4pComGitIndexFileFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitIndexFileFactory) newObject() * index0x3eb559.GitIndexFileFactory {
	return & index0x3eb559.GitIndexFileFactory {}
}

func (inst * comFactory4pComGitIndexFileFactory) castObject(instance application.ComponentInstance) * index0x3eb559.GitIndexFileFactory {
	return instance.Get().(*index0x3eb559.GitIndexFileFactory)
}

func (inst * comFactory4pComGitIndexFileFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitIndexFileFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitIndexFileFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitIndexFileFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitIndexFileFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitIndexFileFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitModulesDirFactory : the factory of component: com3-modules0xe8c0dc.GitModulesDirFactory
type comFactory4pComGitModulesDirFactory struct {

    mPrototype * modules0xe8c0dc.GitModulesDirFactory

	

}

func (inst * comFactory4pComGitModulesDirFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitModulesDirFactory) newObject() * modules0xe8c0dc.GitModulesDirFactory {
	return & modules0xe8c0dc.GitModulesDirFactory {}
}

func (inst * comFactory4pComGitModulesDirFactory) castObject(instance application.ComponentInstance) * modules0xe8c0dc.GitModulesDirFactory {
	return instance.Get().(*modules0xe8c0dc.GitModulesDirFactory)
}

func (inst * comFactory4pComGitModulesDirFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitModulesDirFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitModulesDirFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitModulesDirFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitModulesDirFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitModulesDirFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalGitObjectsFactory : the factory of component: com4-objects0x11508a.LocalGitObjectsFactory
type comFactory4pComLocalGitObjectsFactory struct {

    mPrototype * objects0x11508a.LocalGitObjectsFactory

	

}

func (inst * comFactory4pComLocalGitObjectsFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalGitObjectsFactory) newObject() * objects0x11508a.LocalGitObjectsFactory {
	return & objects0x11508a.LocalGitObjectsFactory {}
}

func (inst * comFactory4pComLocalGitObjectsFactory) castObject(instance application.ComponentInstance) * objects0x11508a.LocalGitObjectsFactory {
	return instance.Get().(*objects0x11508a.LocalGitObjectsFactory)
}

func (inst * comFactory4pComLocalGitObjectsFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalGitObjectsFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalGitObjectsFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalGitObjectsFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitObjectsFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitObjectsFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalGitRefsFactory : the factory of component: com5-refs0x4e5472.LocalGitRefsFactory
type comFactory4pComLocalGitRefsFactory struct {

    mPrototype * refs0x4e5472.LocalGitRefsFactory

	

}

func (inst * comFactory4pComLocalGitRefsFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalGitRefsFactory) newObject() * refs0x4e5472.LocalGitRefsFactory {
	return & refs0x4e5472.LocalGitRefsFactory {}
}

func (inst * comFactory4pComLocalGitRefsFactory) castObject(instance application.ComponentInstance) * refs0x4e5472.LocalGitRefsFactory {
	return instance.Get().(*refs0x4e5472.LocalGitRefsFactory)
}

func (inst * comFactory4pComLocalGitRefsFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalGitRefsFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalGitRefsFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalGitRefsFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitRefsFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalGitRefsFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepoDriver : the factory of component: com6-localfilesys0x6be3ff.LocalRepoDriver
type comFactory4pComLocalRepoDriver struct {

    mPrototype * localfilesys0x6be3ff.LocalRepoDriver

	
	mLayoutSelector config.InjectionSelector
	mLocatorSelector config.InjectionSelector
	mFactorySelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepoDriver) init() application.ComponentFactory {

	
	inst.mLayoutSelector = config.NewInjectionSelector("#git-local-repository-layout",nil)
	inst.mLocatorSelector = config.NewInjectionSelector("#git-local-repository-locator",nil)
	inst.mFactorySelector = config.NewInjectionSelector("#git-local-repository-factory",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepoDriver) newObject() * localfilesys0x6be3ff.LocalRepoDriver {
	return & localfilesys0x6be3ff.LocalRepoDriver {}
}

func (inst * comFactory4pComLocalRepoDriver) castObject(instance application.ComponentInstance) * localfilesys0x6be3ff.LocalRepoDriver {
	return instance.Get().(*localfilesys0x6be3ff.LocalRepoDriver)
}

func (inst * comFactory4pComLocalRepoDriver) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepoDriver) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepoDriver) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepoDriver) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoDriver) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoDriver) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Layout = inst.getterForFieldLayoutSelector(context)
	obj.Locator = inst.getterForFieldLocatorSelector(context)
	obj.Factory = inst.getterForFieldFactorySelector(context)
	return context.LastError()
}

//getterForFieldLayoutSelector
func (inst * comFactory4pComLocalRepoDriver) getterForFieldLayoutSelector (context application.InstanceContext) files0x00707a.Layout {

	o1 := inst.mLayoutSelector.GetOne(context)
	o2, ok := o1.(files0x00707a.Layout)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-localfilesys0x6be3ff.LocalRepoDriver")
		eb.Set("field", "Layout")
		eb.Set("type1", "?")
		eb.Set("type2", "files0x00707a.Layout")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocatorSelector
func (inst * comFactory4pComLocalRepoDriver) getterForFieldLocatorSelector (context application.InstanceContext) files0x00707a.RepositoryLocator {

	o1 := inst.mLocatorSelector.GetOne(context)
	o2, ok := o1.(files0x00707a.RepositoryLocator)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-localfilesys0x6be3ff.LocalRepoDriver")
		eb.Set("field", "Locator")
		eb.Set("type1", "?")
		eb.Set("type2", "files0x00707a.RepositoryLocator")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFactorySelector
func (inst * comFactory4pComLocalRepoDriver) getterForFieldFactorySelector (context application.InstanceContext) localfilesys0x6be3ff.RepoFactory {

	o1 := inst.mFactorySelector.GetOne(context)
	o2, ok := o1.(localfilesys0x6be3ff.RepoFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-localfilesys0x6be3ff.LocalRepoDriver")
		eb.Set("field", "Factory")
		eb.Set("type1", "?")
		eb.Set("type2", "localfilesys0x6be3ff.RepoFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepoFactory : the factory of component: git-local-repository-factory
type comFactory4pComLocalRepoFactory struct {

    mPrototype * localfilesys0x6be3ff.LocalRepoFactory

	
	mElementsSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepoFactory) init() application.ComponentFactory {

	
	inst.mElementsSelector = config.NewInjectionSelector(".git-local-element-factory",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepoFactory) newObject() * localfilesys0x6be3ff.LocalRepoFactory {
	return & localfilesys0x6be3ff.LocalRepoFactory {}
}

func (inst * comFactory4pComLocalRepoFactory) castObject(instance application.ComponentInstance) * localfilesys0x6be3ff.LocalRepoFactory {
	return instance.Get().(*localfilesys0x6be3ff.LocalRepoFactory)
}

func (inst * comFactory4pComLocalRepoFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepoFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepoFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepoFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Elements = inst.getterForFieldElementsSelector(context)
	return context.LastError()
}

//getterForFieldElementsSelector
func (inst * comFactory4pComLocalRepoFactory) getterForFieldElementsSelector (context application.InstanceContext) []localfilesys0x6be3ff.ElementFactory {
	list1 := inst.mElementsSelector.GetList(context)
	list2 := make([]localfilesys0x6be3ff.ElementFactory, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(localfilesys0x6be3ff.ElementFactory)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepoLayout : the factory of component: git-local-repository-layout
type comFactory4pComLocalRepoLayout struct {

    mPrototype * localfilesys0x6be3ff.LocalRepoLayout

	

}

func (inst * comFactory4pComLocalRepoLayout) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepoLayout) newObject() * localfilesys0x6be3ff.LocalRepoLayout {
	return & localfilesys0x6be3ff.LocalRepoLayout {}
}

func (inst * comFactory4pComLocalRepoLayout) castObject(instance application.ComponentInstance) * localfilesys0x6be3ff.LocalRepoLayout {
	return instance.Get().(*localfilesys0x6be3ff.LocalRepoLayout)
}

func (inst * comFactory4pComLocalRepoLayout) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepoLayout) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepoLayout) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepoLayout) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoLayout) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoLayout) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepoLocator : the factory of component: git-local-repository-locator
type comFactory4pComLocalRepoLocator struct {

    mPrototype * localfilesys0x6be3ff.LocalRepoLocator

	

}

func (inst * comFactory4pComLocalRepoLocator) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepoLocator) newObject() * localfilesys0x6be3ff.LocalRepoLocator {
	return & localfilesys0x6be3ff.LocalRepoLocator {}
}

func (inst * comFactory4pComLocalRepoLocator) castObject(instance application.ComponentInstance) * localfilesys0x6be3ff.LocalRepoLocator {
	return instance.Get().(*localfilesys0x6be3ff.LocalRepoLocator)
}

func (inst * comFactory4pComLocalRepoLocator) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepoLocator) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepoLocator) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepoLocator) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoLocator) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepoLocator) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCoreDirFactory : the factory of component: com10-views0x7d154b.CoreDirFactory
type comFactory4pComCoreDirFactory struct {

    mPrototype * views0x7d154b.CoreDirFactory

	

}

func (inst * comFactory4pComCoreDirFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCoreDirFactory) newObject() * views0x7d154b.CoreDirFactory {
	return & views0x7d154b.CoreDirFactory {}
}

func (inst * comFactory4pComCoreDirFactory) castObject(instance application.ComponentInstance) * views0x7d154b.CoreDirFactory {
	return instance.Get().(*views0x7d154b.CoreDirFactory)
}

func (inst * comFactory4pComCoreDirFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCoreDirFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCoreDirFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCoreDirFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCoreDirFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCoreDirFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComShellDirFactory : the factory of component: com11-views0x7d154b.ShellDirFactory
type comFactory4pComShellDirFactory struct {

    mPrototype * views0x7d154b.ShellDirFactory

	

}

func (inst * comFactory4pComShellDirFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComShellDirFactory) newObject() * views0x7d154b.ShellDirFactory {
	return & views0x7d154b.ShellDirFactory {}
}

func (inst * comFactory4pComShellDirFactory) castObject(instance application.ComponentInstance) * views0x7d154b.ShellDirFactory {
	return instance.Get().(*views0x7d154b.ShellDirFactory)
}

func (inst * comFactory4pComShellDirFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComShellDirFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComShellDirFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComShellDirFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComShellDirFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComShellDirFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepoViewFactory : the factory of component: com12-views0x7d154b.RepoViewFactory
type comFactory4pComRepoViewFactory struct {

    mPrototype * views0x7d154b.RepoViewFactory

	

}

func (inst * comFactory4pComRepoViewFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepoViewFactory) newObject() * views0x7d154b.RepoViewFactory {
	return & views0x7d154b.RepoViewFactory {}
}

func (inst * comFactory4pComRepoViewFactory) castObject(instance application.ComponentInstance) * views0x7d154b.RepoViewFactory {
	return instance.Get().(*views0x7d154b.RepoViewFactory)
}

func (inst * comFactory4pComRepoViewFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepoViewFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepoViewFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepoViewFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepoViewFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepoViewFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWorkingDirFactory : the factory of component: com13-views0x7d154b.WorkingDirFactory
type comFactory4pComWorkingDirFactory struct {

    mPrototype * views0x7d154b.WorkingDirFactory

	

}

func (inst * comFactory4pComWorkingDirFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWorkingDirFactory) newObject() * views0x7d154b.WorkingDirFactory {
	return & views0x7d154b.WorkingDirFactory {}
}

func (inst * comFactory4pComWorkingDirFactory) castObject(instance application.ComponentInstance) * views0x7d154b.WorkingDirFactory {
	return instance.Get().(*views0x7d154b.WorkingDirFactory)
}

func (inst * comFactory4pComWorkingDirFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWorkingDirFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWorkingDirFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWorkingDirFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWorkingDirFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWorkingDirFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitWorktreesDirFactory : the factory of component: com14-worktrees0xa9c0a9.GitWorktreesDirFactory
type comFactory4pComGitWorktreesDirFactory struct {

    mPrototype * worktrees0xa9c0a9.GitWorktreesDirFactory

	

}

func (inst * comFactory4pComGitWorktreesDirFactory) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitWorktreesDirFactory) newObject() * worktrees0xa9c0a9.GitWorktreesDirFactory {
	return & worktrees0xa9c0a9.GitWorktreesDirFactory {}
}

func (inst * comFactory4pComGitWorktreesDirFactory) castObject(instance application.ComponentInstance) * worktrees0xa9c0a9.GitWorktreesDirFactory {
	return instance.Get().(*worktrees0xa9c0a9.GitWorktreesDirFactory)
}

func (inst * comFactory4pComGitWorktreesDirFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitWorktreesDirFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitWorktreesDirFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitWorktreesDirFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitWorktreesDirFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitWorktreesDirFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitRepoManager : the factory of component: git-repository-manager
type comFactory4pComGitRepoManager struct {

    mPrototype * support0x074feb.GitRepoManager

	
	mDriversSelector config.InjectionSelector

}

func (inst * comFactory4pComGitRepoManager) init() application.ComponentFactory {

	
	inst.mDriversSelector = config.NewInjectionSelector(".git-repository-driver",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitRepoManager) newObject() * support0x074feb.GitRepoManager {
	return & support0x074feb.GitRepoManager {}
}

func (inst * comFactory4pComGitRepoManager) castObject(instance application.ComponentInstance) * support0x074feb.GitRepoManager {
	return instance.Get().(*support0x074feb.GitRepoManager)
}

func (inst * comFactory4pComGitRepoManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitRepoManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitRepoManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitRepoManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitRepoManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitRepoManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Drivers = inst.getterForFieldDriversSelector(context)
	return context.LastError()
}

//getterForFieldDriversSelector
func (inst * comFactory4pComGitRepoManager) getterForFieldDriversSelector (context application.InstanceContext) []repository0x5aaf5a.Driver {
	list1 := inst.mDriversSelector.GetList(context)
	list2 := make([]repository0x5aaf5a.Driver, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(repository0x5aaf5a.Driver)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}




