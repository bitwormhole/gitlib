// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibdemo

import (
	demo0x52dcb1 "github.com/bitwormhole/gitlib/etc/demo"
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

	// component: com0-demo0x52dcb1.TestPoint
	cominfobuilder.Next()
	cominfobuilder.ID("com0-demo0x52dcb1.TestPoint").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPoint{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint : the factory of component: com0-demo0x52dcb1.TestPoint
type comFactory4pComTestPoint struct {

    mPrototype * demo0x52dcb1.TestPoint

	
	mAgentSelector config.InjectionSelector
	mCommandSelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mCommandSelector = config.NewInjectionSelector("${test.gitlib.command}",nil)


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
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComTestPoint) getterForFieldAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-demo0x52dcb1.TestPoint")
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




