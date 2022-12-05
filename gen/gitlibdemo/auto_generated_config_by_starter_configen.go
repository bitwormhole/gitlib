// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibdemo

import (
	demo0x52dcb1 "github.com/bitwormhole/gitlib/etc/demo"
	testcmds0x82bca1 "github.com/bitwormhole/gitlib/etc/demo/testcmds"
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
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

	// component: com0-testcmds0x82bca1.TestListObjectsInPack
	cominfobuilder.Next()
	cominfobuilder.ID("com0-testcmds0x82bca1.TestListObjectsInPack").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestListObjectsInPack{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-testcmds0x82bca1.TestReadObjects
	cominfobuilder.Next()
	cominfobuilder.ID("com1-testcmds0x82bca1.TestReadObjects").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestReadObjects{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-testcmds0x82bca1.TestReadPackIdx
	cominfobuilder.Next()
	cominfobuilder.ID("com2-testcmds0x82bca1.TestReadPackIdx").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestReadPackIdx{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-demo0x52dcb1.TestPoint
	cominfobuilder.Next()
	cominfobuilder.ID("com3-demo0x52dcb1.TestPoint").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPoint{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestListObjectsInPack : the factory of component: com0-testcmds0x82bca1.TestListObjectsInPack
type comFactory4pComTestListObjectsInPack struct {

    mPrototype * testcmds0x82bca1.TestListObjectsInPack

	
	mWDSelector config.InjectionSelector
	mLASelector config.InjectionSelector

}

func (inst * comFactory4pComTestListObjectsInPack) init() application.ComponentFactory {

	
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)
	inst.mLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestListObjectsInPack) newObject() * testcmds0x82bca1.TestListObjectsInPack {
	return & testcmds0x82bca1.TestListObjectsInPack {}
}

func (inst * comFactory4pComTestListObjectsInPack) castObject(instance application.ComponentInstance) * testcmds0x82bca1.TestListObjectsInPack {
	return instance.Get().(*testcmds0x82bca1.TestListObjectsInPack)
}

func (inst * comFactory4pComTestListObjectsInPack) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestListObjectsInPack) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestListObjectsInPack) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestListObjectsInPack) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestListObjectsInPack) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestListObjectsInPack) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.WD = inst.getterForFieldWDSelector(context)
	obj.LA = inst.getterForFieldLASelector(context)
	return context.LastError()
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestListObjectsInPack) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}

//getterForFieldLASelector
func (inst * comFactory4pComTestListObjectsInPack) getterForFieldLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-testcmds0x82bca1.TestListObjectsInPack")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestReadObjects : the factory of component: com1-testcmds0x82bca1.TestReadObjects
type comFactory4pComTestReadObjects struct {

    mPrototype * testcmds0x82bca1.TestReadObjects

	
	mWDSelector config.InjectionSelector
	mLASelector config.InjectionSelector

}

func (inst * comFactory4pComTestReadObjects) init() application.ComponentFactory {

	
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)
	inst.mLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestReadObjects) newObject() * testcmds0x82bca1.TestReadObjects {
	return & testcmds0x82bca1.TestReadObjects {}
}

func (inst * comFactory4pComTestReadObjects) castObject(instance application.ComponentInstance) * testcmds0x82bca1.TestReadObjects {
	return instance.Get().(*testcmds0x82bca1.TestReadObjects)
}

func (inst * comFactory4pComTestReadObjects) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestReadObjects) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestReadObjects) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestReadObjects) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestReadObjects) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestReadObjects) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.WD = inst.getterForFieldWDSelector(context)
	obj.LA = inst.getterForFieldLASelector(context)
	return context.LastError()
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestReadObjects) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}

//getterForFieldLASelector
func (inst * comFactory4pComTestReadObjects) getterForFieldLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-testcmds0x82bca1.TestReadObjects")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestReadPackIdx : the factory of component: com2-testcmds0x82bca1.TestReadPackIdx
type comFactory4pComTestReadPackIdx struct {

    mPrototype * testcmds0x82bca1.TestReadPackIdx

	
	mWDSelector config.InjectionSelector
	mLASelector config.InjectionSelector

}

func (inst * comFactory4pComTestReadPackIdx) init() application.ComponentFactory {

	
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)
	inst.mLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestReadPackIdx) newObject() * testcmds0x82bca1.TestReadPackIdx {
	return & testcmds0x82bca1.TestReadPackIdx {}
}

func (inst * comFactory4pComTestReadPackIdx) castObject(instance application.ComponentInstance) * testcmds0x82bca1.TestReadPackIdx {
	return instance.Get().(*testcmds0x82bca1.TestReadPackIdx)
}

func (inst * comFactory4pComTestReadPackIdx) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestReadPackIdx) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestReadPackIdx) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestReadPackIdx) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestReadPackIdx) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestReadPackIdx) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.WD = inst.getterForFieldWDSelector(context)
	obj.LA = inst.getterForFieldLASelector(context)
	return context.LastError()
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestReadPackIdx) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}

//getterForFieldLASelector
func (inst * comFactory4pComTestReadPackIdx) getterForFieldLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com2-testcmds0x82bca1.TestReadPackIdx")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint : the factory of component: com3-demo0x52dcb1.TestPoint
type comFactory4pComTestPoint struct {

    mPrototype * demo0x52dcb1.TestPoint

	
	mAgentSelector config.InjectionSelector
	mCommandSelector config.InjectionSelector
	mWDSelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mCommandSelector = config.NewInjectionSelector("${test.gitlib.command}",nil)
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestPoint) newObject() * demo0x52dcb1.TestPoint {
	return & demo0x52dcb1.TestPoint {}
}

func (inst * comFactory4pComTestPoint) castObject(instance application.ComponentInstance) * demo0x52dcb1.TestPoint {
	return instance.Get().(*demo0x52dcb1.TestPoint)
}

func (inst * comFactory4pComTestPoint) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestPoint) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestPoint) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestPoint) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.Command = inst.getterForFieldCommandSelector(context)
	obj.WD = inst.getterForFieldWDSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComTestPoint) getterForFieldAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com3-demo0x52dcb1.TestPoint")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldCommandSelector
func (inst * comFactory4pComTestPoint) getterForFieldCommandSelector (context application.InstanceContext) string {
    return inst.mCommandSelector.GetString(context)
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestPoint) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}




