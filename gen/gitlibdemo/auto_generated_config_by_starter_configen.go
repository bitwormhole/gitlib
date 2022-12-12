// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibdemo

import (
	demo0x52dcb1 "github.com/bitwormhole/gitlib/etc/demo"
	testcmds0x82bca1 "github.com/bitwormhole/gitlib/etc/demo/testcmds"
	servers0xb5845d "github.com/bitwormhole/gitlib/git/network/servers"
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

	// component: com1-testcmds0x82bca1.TestPackDeltaObjects
	cominfobuilder.Next()
	cominfobuilder.ID("com1-testcmds0x82bca1.TestPackDeltaObjects").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPackDeltaObjects{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-testcmds0x82bca1.TestReadObjects
	cominfobuilder.Next()
	cominfobuilder.ID("com2-testcmds0x82bca1.TestReadObjects").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestReadObjects{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-testcmds0x82bca1.TestReadPackIdx
	cominfobuilder.Next()
	cominfobuilder.ID("com3-testcmds0x82bca1.TestReadPackIdx").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestReadPackIdx{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-testcmds0x82bca1.TestServerAPI
	cominfobuilder.Next()
	cominfobuilder.ID("com4-testcmds0x82bca1.TestServerAPI").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestServerAPI{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-demo0x52dcb1.TestPoint
	cominfobuilder.Next()
	cominfobuilder.ID("com5-demo0x52dcb1.TestPoint").Class("life").Aliases("").Scope("")
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

// comFactory4pComTestPackDeltaObjects : the factory of component: com1-testcmds0x82bca1.TestPackDeltaObjects
type comFactory4pComTestPackDeltaObjects struct {

    mPrototype * testcmds0x82bca1.TestPackDeltaObjects

	
	mWDSelector config.InjectionSelector
	mLASelector config.InjectionSelector

}

func (inst * comFactory4pComTestPackDeltaObjects) init() application.ComponentFactory {

	
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)
	inst.mLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestPackDeltaObjects) newObject() * testcmds0x82bca1.TestPackDeltaObjects {
	return & testcmds0x82bca1.TestPackDeltaObjects {}
}

func (inst * comFactory4pComTestPackDeltaObjects) castObject(instance application.ComponentInstance) * testcmds0x82bca1.TestPackDeltaObjects {
	return instance.Get().(*testcmds0x82bca1.TestPackDeltaObjects)
}

func (inst * comFactory4pComTestPackDeltaObjects) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestPackDeltaObjects) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestPackDeltaObjects) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestPackDeltaObjects) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPackDeltaObjects) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPackDeltaObjects) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.WD = inst.getterForFieldWDSelector(context)
	obj.LA = inst.getterForFieldLASelector(context)
	return context.LastError()
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestPackDeltaObjects) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}

//getterForFieldLASelector
func (inst * comFactory4pComTestPackDeltaObjects) getterForFieldLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-testcmds0x82bca1.TestPackDeltaObjects")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestReadObjects : the factory of component: com2-testcmds0x82bca1.TestReadObjects
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
		eb.Set("com", "com2-testcmds0x82bca1.TestReadObjects")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestReadPackIdx : the factory of component: com3-testcmds0x82bca1.TestReadPackIdx
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
		eb.Set("com", "com3-testcmds0x82bca1.TestReadPackIdx")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestServerAPI : the factory of component: com4-testcmds0x82bca1.TestServerAPI
type comFactory4pComTestServerAPI struct {

    mPrototype * testcmds0x82bca1.TestServerAPI

	
	mWDSelector config.InjectionSelector
	mLASelector config.InjectionSelector
	mMainServerSelector config.InjectionSelector

}

func (inst * comFactory4pComTestServerAPI) init() application.ComponentFactory {

	
	inst.mWDSelector = config.NewInjectionSelector("${test.repo.path}",nil)
	inst.mLASelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mMainServerSelector = config.NewInjectionSelector("#git-main-server",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestServerAPI) newObject() * testcmds0x82bca1.TestServerAPI {
	return & testcmds0x82bca1.TestServerAPI {}
}

func (inst * comFactory4pComTestServerAPI) castObject(instance application.ComponentInstance) * testcmds0x82bca1.TestServerAPI {
	return instance.Get().(*testcmds0x82bca1.TestServerAPI)
}

func (inst * comFactory4pComTestServerAPI) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestServerAPI) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestServerAPI) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestServerAPI) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestServerAPI) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestServerAPI) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.WD = inst.getterForFieldWDSelector(context)
	obj.LA = inst.getterForFieldLASelector(context)
	obj.MainServer = inst.getterForFieldMainServerSelector(context)
	return context.LastError()
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestServerAPI) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}

//getterForFieldLASelector
func (inst * comFactory4pComTestServerAPI) getterForFieldLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-testcmds0x82bca1.TestServerAPI")
		eb.Set("field", "LA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldMainServerSelector
func (inst * comFactory4pComTestServerAPI) getterForFieldMainServerSelector (context application.InstanceContext) servers0xb5845d.MainServer {

	o1 := inst.mMainServerSelector.GetOne(context)
	o2, ok := o1.(servers0xb5845d.MainServer)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-testcmds0x82bca1.TestServerAPI")
		eb.Set("field", "MainServer")
		eb.Set("type1", "?")
		eb.Set("type2", "servers0xb5845d.MainServer")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint : the factory of component: com5-demo0x52dcb1.TestPoint
type comFactory4pComTestPoint struct {

    mPrototype * demo0x52dcb1.TestPoint

	
	mContextSelector config.InjectionSelector
	mAgentSelector config.InjectionSelector
	mCmdKeySelector config.InjectionSelector
	mWDSelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mCmdKeySelector = config.NewInjectionSelector("${test.gitlib.command}",nil)
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
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.CmdKey = inst.getterForFieldCmdKeySelector(context)
	obj.WD = inst.getterForFieldWDSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComTestPoint) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComTestPoint) getterForFieldAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-demo0x52dcb1.TestPoint")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldCmdKeySelector
func (inst * comFactory4pComTestPoint) getterForFieldCmdKeySelector (context application.InstanceContext) string {
    return inst.mCmdKeySelector.GetString(context)
}

//getterForFieldWDSelector
func (inst * comFactory4pComTestPoint) getterForFieldWDSelector (context application.InstanceContext) string {
    return inst.mWDSelector.GetString(context)
}




