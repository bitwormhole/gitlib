// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	repository0x5aaf5a "github.com/bitwormhole/gitlib/git/repository"
	unit0xa250fd "github.com/bitwormhole/gitlib/src/test/golang/unit"
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

	// component: com0-unit0xa250fd.UnitTest1
	cominfobuilder.Next()
	cominfobuilder.ID("com0-unit0xa250fd.UnitTest1").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUnitTest1{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUnitTest1 : the factory of component: com0-unit0xa250fd.UnitTest1
type comFactory4pComUnitTest1 struct {

    mPrototype * unit0xa250fd.UnitTest1

	
	mRMSelector config.InjectionSelector
	mRepoPathSelector config.InjectionSelector

}

func (inst * comFactory4pComUnitTest1) init() application.ComponentFactory {

	
	inst.mRMSelector = config.NewInjectionSelector("#git-repository-manager",nil)
	inst.mRepoPathSelector = config.NewInjectionSelector("${test.repo.path}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUnitTest1) newObject() * unit0xa250fd.UnitTest1 {
	return & unit0xa250fd.UnitTest1 {}
}

func (inst * comFactory4pComUnitTest1) castObject(instance application.ComponentInstance) * unit0xa250fd.UnitTest1 {
	return instance.Get().(*unit0xa250fd.UnitTest1)
}

func (inst * comFactory4pComUnitTest1) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUnitTest1) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUnitTest1) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUnitTest1) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComUnitTest1) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUnitTest1) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RM = inst.getterForFieldRMSelector(context)
	obj.RepoPath = inst.getterForFieldRepoPathSelector(context)
	return context.LastError()
}

//getterForFieldRMSelector
func (inst * comFactory4pComUnitTest1) getterForFieldRMSelector (context application.InstanceContext) repository0x5aaf5a.Manager {

	o1 := inst.mRMSelector.GetOne(context)
	o2, ok := o1.(repository0x5aaf5a.Manager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-unit0xa250fd.UnitTest1")
		eb.Set("field", "RM")
		eb.Set("type1", "?")
		eb.Set("type2", "repository0x5aaf5a.Manager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoPathSelector
func (inst * comFactory4pComUnitTest1) getterForFieldRepoPathSelector (context application.InstanceContext) string {
    return inst.mRepoPathSelector.GetString(context)
}




