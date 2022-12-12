// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibcfg

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x4595be "github.com/bitwormhole/gitlib/etc/lib"
	git0x229c8a "github.com/bitwormhole/gitlib/git"
	clients0x465781 "github.com/bitwormhole/gitlib/git/network/clients"
	http0xe484f4 "github.com/bitwormhole/gitlib/git/network/http"
	pktline0xd37953 "github.com/bitwormhole/gitlib/git/network/pktline"
	servers0xb5845d "github.com/bitwormhole/gitlib/git/network/servers"
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	http4git0x537f63 "github.com/bitwormhole/gitlib/git/support/net/http4git"
	services0x887aef "github.com/bitwormhole/gitlib/git/support/services"
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

	// component: git-main-client
	cominfobuilder.Next()
	cominfobuilder.ID("git-main-client").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainClientImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-http0xe484f4.GitClient
	cominfobuilder.Next()
	cominfobuilder.ID("com1-http0xe484f4.GitClient").Class("git-client-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitClient{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-main-server
	cominfobuilder.Next()
	cominfobuilder.ID("git-main-server").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainServerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-http4git0x537f63.HTTPGitConnector
	cominfobuilder.Next()
	cominfobuilder.ID("com3-http4git0x537f63.HTTPGitConnector").Class("pktline-connector-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPGitConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-services0x887aef.GitFetchService
	cominfobuilder.Next()
	cominfobuilder.ID("com4-services0x887aef.GitFetchService").Class("git-instruction-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitFetchService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-services0x887aef.GitPushService
	cominfobuilder.Next()
	cominfobuilder.ID("com5-services0x887aef.GitPushService").Class("git-instruction-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitPushService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-lib0x4595be.ConfigAlgorithms
	cominfobuilder.Next()
	cominfobuilder.ID("com6-lib0x4595be.ConfigAlgorithms").Class("git-context-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigAlgorithms{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com7-lib0x4595be.TheSHA1
	cominfobuilder.Next()
	cominfobuilder.ID("com7-lib0x4595be.TheSHA1").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheSHA1{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-lib0x4595be.TheSHA256
	cominfobuilder.Next()
	cominfobuilder.ID("com8-lib0x4595be.TheSHA256").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheSHA256{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com9-lib0x4595be.TheSHA512
	cominfobuilder.Next()
	cominfobuilder.ID("com9-lib0x4595be.TheSHA512").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheSHA512{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com10-lib0x4595be.TheMD5
	cominfobuilder.Next()
	cominfobuilder.ID("com10-lib0x4595be.TheMD5").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheMD5{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com11-lib0x4595be.TheDeflate
	cominfobuilder.Next()
	cominfobuilder.ID("com11-lib0x4595be.TheDeflate").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheDeflate{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com12-lib0x4595be.ThePlain
	cominfobuilder.Next()
	cominfobuilder.ID("com12-lib0x4595be.ThePlain").Class("git-algorithm-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComThePlain{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com13-lib0x4595be.ConfigCommands
	cominfobuilder.Next()
	cominfobuilder.ID("com13-lib0x4595be.ConfigCommands").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigCommands{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com14-lib0x4595be.ConfigConnectors
	cominfobuilder.Next()
	cominfobuilder.ID("com14-lib0x4595be.ConfigConnectors").Class("git-context-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigConnectors{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com15-lib0x4595be.HTTPGitConnectorReg
	cominfobuilder.Next()
	cominfobuilder.ID("com15-lib0x4595be.HTTPGitConnectorReg").Class("pktline-connector-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPGitConnectorReg{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com16-lib0x4595be.ConfigContextBase
	cominfobuilder.Next()
	cominfobuilder.ID("com16-lib0x4595be.ConfigContextBase").Class("git-context-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigContextBase{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com17-lib0x4595be.ConfigContextWithInstructions
	cominfobuilder.Next()
	cominfobuilder.ID("com17-lib0x4595be.ConfigContextWithInstructions").Class("git-context-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigContextWithInstructions{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com18-lib0x4595be.ConfigCore
	cominfobuilder.Next()
	cominfobuilder.ID("com18-lib0x4595be.ConfigCore").Class("git-core-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigCore{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com19-lib0x4595be.ConfigInstructions
	cominfobuilder.Next()
	cominfobuilder.ID("com19-lib0x4595be.ConfigInstructions").Class("git-instruction-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigInstructions{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: git-lib-agent
	cominfobuilder.Next()
	cominfobuilder.ID("git-lib-agent").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGitlibAgent{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainClientImpl : the factory of component: git-main-client
type comFactory4pComMainClientImpl struct {

    mPrototype * clients0x465781.MainClientImpl

	
	mClientRegistryListSelector config.InjectionSelector

}

func (inst * comFactory4pComMainClientImpl) init() application.ComponentFactory {

	
	inst.mClientRegistryListSelector = config.NewInjectionSelector(".git-client-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainClientImpl) newObject() * clients0x465781.MainClientImpl {
	return & clients0x465781.MainClientImpl {}
}

func (inst * comFactory4pComMainClientImpl) castObject(instance application.ComponentInstance) * clients0x465781.MainClientImpl {
	return instance.Get().(*clients0x465781.MainClientImpl)
}

func (inst * comFactory4pComMainClientImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainClientImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainClientImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainClientImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainClientImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainClientImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ClientRegistryList = inst.getterForFieldClientRegistryListSelector(context)
	return context.LastError()
}

//getterForFieldClientRegistryListSelector
func (inst * comFactory4pComMainClientImpl) getterForFieldClientRegistryListSelector (context application.InstanceContext) []clients0x465781.ClientRegistry {
	list1 := inst.mClientRegistryListSelector.GetList(context)
	list2 := make([]clients0x465781.ClientRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(clients0x465781.ClientRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitClient : the factory of component: com1-http0xe484f4.GitClient
type comFactory4pComGitClient struct {

    mPrototype * http0xe484f4.GitClient

	

}

func (inst * comFactory4pComGitClient) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitClient) newObject() * http0xe484f4.GitClient {
	return & http0xe484f4.GitClient {}
}

func (inst * comFactory4pComGitClient) castObject(instance application.ComponentInstance) * http0xe484f4.GitClient {
	return instance.Get().(*http0xe484f4.GitClient)
}

func (inst * comFactory4pComGitClient) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitClient) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitClient) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitClient) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitClient) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitClient) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainServerImpl : the factory of component: git-main-server
type comFactory4pComMainServerImpl struct {

    mPrototype * servers0xb5845d.MainServerImpl

	
	mServerRegistryListSelector config.InjectionSelector

}

func (inst * comFactory4pComMainServerImpl) init() application.ComponentFactory {

	
	inst.mServerRegistryListSelector = config.NewInjectionSelector(".git-server-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainServerImpl) newObject() * servers0xb5845d.MainServerImpl {
	return & servers0xb5845d.MainServerImpl {}
}

func (inst * comFactory4pComMainServerImpl) castObject(instance application.ComponentInstance) * servers0xb5845d.MainServerImpl {
	return instance.Get().(*servers0xb5845d.MainServerImpl)
}

func (inst * comFactory4pComMainServerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainServerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainServerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainServerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainServerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainServerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ServerRegistryList = inst.getterForFieldServerRegistryListSelector(context)
	return context.LastError()
}

//getterForFieldServerRegistryListSelector
func (inst * comFactory4pComMainServerImpl) getterForFieldServerRegistryListSelector (context application.InstanceContext) []servers0xb5845d.ServerRegistry {
	list1 := inst.mServerRegistryListSelector.GetList(context)
	list2 := make([]servers0xb5845d.ServerRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(servers0xb5845d.ServerRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPGitConnector : the factory of component: com3-http4git0x537f63.HTTPGitConnector
type comFactory4pComHTTPGitConnector struct {

    mPrototype * http4git0x537f63.HTTPGitConnector

	

}

func (inst * comFactory4pComHTTPGitConnector) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPGitConnector) newObject() * http4git0x537f63.HTTPGitConnector {
	return & http4git0x537f63.HTTPGitConnector {}
}

func (inst * comFactory4pComHTTPGitConnector) castObject(instance application.ComponentInstance) * http4git0x537f63.HTTPGitConnector {
	return instance.Get().(*http4git0x537f63.HTTPGitConnector)
}

func (inst * comFactory4pComHTTPGitConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPGitConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPGitConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPGitConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGitConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGitConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitFetchService : the factory of component: com4-services0x887aef.GitFetchService
type comFactory4pComGitFetchService struct {

    mPrototype * services0x887aef.GitFetchService

	
	mMainClientSelector config.InjectionSelector

}

func (inst * comFactory4pComGitFetchService) init() application.ComponentFactory {

	
	inst.mMainClientSelector = config.NewInjectionSelector("#git-main-client",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitFetchService) newObject() * services0x887aef.GitFetchService {
	return & services0x887aef.GitFetchService {}
}

func (inst * comFactory4pComGitFetchService) castObject(instance application.ComponentInstance) * services0x887aef.GitFetchService {
	return instance.Get().(*services0x887aef.GitFetchService)
}

func (inst * comFactory4pComGitFetchService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitFetchService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitFetchService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitFetchService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitFetchService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitFetchService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MainClient = inst.getterForFieldMainClientSelector(context)
	return context.LastError()
}

//getterForFieldMainClientSelector
func (inst * comFactory4pComGitFetchService) getterForFieldMainClientSelector (context application.InstanceContext) clients0x465781.MainClient {

	o1 := inst.mMainClientSelector.GetOne(context)
	o2, ok := o1.(clients0x465781.MainClient)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-services0x887aef.GitFetchService")
		eb.Set("field", "MainClient")
		eb.Set("type1", "?")
		eb.Set("type2", "clients0x465781.MainClient")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitPushService : the factory of component: com5-services0x887aef.GitPushService
type comFactory4pComGitPushService struct {

    mPrototype * services0x887aef.GitPushService

	
	mMainClientSelector config.InjectionSelector

}

func (inst * comFactory4pComGitPushService) init() application.ComponentFactory {

	
	inst.mMainClientSelector = config.NewInjectionSelector("#git-main-client",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitPushService) newObject() * services0x887aef.GitPushService {
	return & services0x887aef.GitPushService {}
}

func (inst * comFactory4pComGitPushService) castObject(instance application.ComponentInstance) * services0x887aef.GitPushService {
	return instance.Get().(*services0x887aef.GitPushService)
}

func (inst * comFactory4pComGitPushService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitPushService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitPushService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitPushService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitPushService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitPushService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MainClient = inst.getterForFieldMainClientSelector(context)
	return context.LastError()
}

//getterForFieldMainClientSelector
func (inst * comFactory4pComGitPushService) getterForFieldMainClientSelector (context application.InstanceContext) clients0x465781.MainClient {

	o1 := inst.mMainClientSelector.GetOne(context)
	o2, ok := o1.(clients0x465781.MainClient)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-services0x887aef.GitPushService")
		eb.Set("field", "MainClient")
		eb.Set("type1", "?")
		eb.Set("type2", "clients0x465781.MainClient")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigAlgorithms : the factory of component: com6-lib0x4595be.ConfigAlgorithms
type comFactory4pComConfigAlgorithms struct {

    mPrototype * lib0x4595be.ConfigAlgorithms

	
	mAlgorithmsSelector config.InjectionSelector

}

func (inst * comFactory4pComConfigAlgorithms) init() application.ComponentFactory {

	
	inst.mAlgorithmsSelector = config.NewInjectionSelector(".git-algorithm-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigAlgorithms) newObject() * lib0x4595be.ConfigAlgorithms {
	return & lib0x4595be.ConfigAlgorithms {}
}

func (inst * comFactory4pComConfigAlgorithms) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigAlgorithms {
	return instance.Get().(*lib0x4595be.ConfigAlgorithms)
}

func (inst * comFactory4pComConfigAlgorithms) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigAlgorithms) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigAlgorithms) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigAlgorithms) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigAlgorithms) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigAlgorithms) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Algorithms = inst.getterForFieldAlgorithmsSelector(context)
	return context.LastError()
}

//getterForFieldAlgorithmsSelector
func (inst * comFactory4pComConfigAlgorithms) getterForFieldAlgorithmsSelector (context application.InstanceContext) []git0x229c8a.AlgorithmRegistry {
	list1 := inst.mAlgorithmsSelector.GetList(context)
	list2 := make([]git0x229c8a.AlgorithmRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(git0x229c8a.AlgorithmRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheSHA1 : the factory of component: com7-lib0x4595be.TheSHA1
type comFactory4pComTheSHA1 struct {

    mPrototype * lib0x4595be.TheSHA1

	

}

func (inst * comFactory4pComTheSHA1) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheSHA1) newObject() * lib0x4595be.TheSHA1 {
	return & lib0x4595be.TheSHA1 {}
}

func (inst * comFactory4pComTheSHA1) castObject(instance application.ComponentInstance) * lib0x4595be.TheSHA1 {
	return instance.Get().(*lib0x4595be.TheSHA1)
}

func (inst * comFactory4pComTheSHA1) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheSHA1) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheSHA1) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheSHA1) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA1) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA1) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheSHA256 : the factory of component: com8-lib0x4595be.TheSHA256
type comFactory4pComTheSHA256 struct {

    mPrototype * lib0x4595be.TheSHA256

	

}

func (inst * comFactory4pComTheSHA256) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheSHA256) newObject() * lib0x4595be.TheSHA256 {
	return & lib0x4595be.TheSHA256 {}
}

func (inst * comFactory4pComTheSHA256) castObject(instance application.ComponentInstance) * lib0x4595be.TheSHA256 {
	return instance.Get().(*lib0x4595be.TheSHA256)
}

func (inst * comFactory4pComTheSHA256) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheSHA256) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheSHA256) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheSHA256) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA256) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA256) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheSHA512 : the factory of component: com9-lib0x4595be.TheSHA512
type comFactory4pComTheSHA512 struct {

    mPrototype * lib0x4595be.TheSHA512

	

}

func (inst * comFactory4pComTheSHA512) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheSHA512) newObject() * lib0x4595be.TheSHA512 {
	return & lib0x4595be.TheSHA512 {}
}

func (inst * comFactory4pComTheSHA512) castObject(instance application.ComponentInstance) * lib0x4595be.TheSHA512 {
	return instance.Get().(*lib0x4595be.TheSHA512)
}

func (inst * comFactory4pComTheSHA512) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheSHA512) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheSHA512) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheSHA512) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA512) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheSHA512) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheMD5 : the factory of component: com10-lib0x4595be.TheMD5
type comFactory4pComTheMD5 struct {

    mPrototype * lib0x4595be.TheMD5

	

}

func (inst * comFactory4pComTheMD5) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheMD5) newObject() * lib0x4595be.TheMD5 {
	return & lib0x4595be.TheMD5 {}
}

func (inst * comFactory4pComTheMD5) castObject(instance application.ComponentInstance) * lib0x4595be.TheMD5 {
	return instance.Get().(*lib0x4595be.TheMD5)
}

func (inst * comFactory4pComTheMD5) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheMD5) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheMD5) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheMD5) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheMD5) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheMD5) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheDeflate : the factory of component: com11-lib0x4595be.TheDeflate
type comFactory4pComTheDeflate struct {

    mPrototype * lib0x4595be.TheDeflate

	

}

func (inst * comFactory4pComTheDeflate) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheDeflate) newObject() * lib0x4595be.TheDeflate {
	return & lib0x4595be.TheDeflate {}
}

func (inst * comFactory4pComTheDeflate) castObject(instance application.ComponentInstance) * lib0x4595be.TheDeflate {
	return instance.Get().(*lib0x4595be.TheDeflate)
}

func (inst * comFactory4pComTheDeflate) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheDeflate) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheDeflate) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheDeflate) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheDeflate) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheDeflate) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComThePlain : the factory of component: com12-lib0x4595be.ThePlain
type comFactory4pComThePlain struct {

    mPrototype * lib0x4595be.ThePlain

	

}

func (inst * comFactory4pComThePlain) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComThePlain) newObject() * lib0x4595be.ThePlain {
	return & lib0x4595be.ThePlain {}
}

func (inst * comFactory4pComThePlain) castObject(instance application.ComponentInstance) * lib0x4595be.ThePlain {
	return instance.Get().(*lib0x4595be.ThePlain)
}

func (inst * comFactory4pComThePlain) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComThePlain) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComThePlain) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComThePlain) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComThePlain) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComThePlain) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigCommands : the factory of component: com13-lib0x4595be.ConfigCommands
type comFactory4pComConfigCommands struct {

    mPrototype * lib0x4595be.ConfigCommands

	

}

func (inst * comFactory4pComConfigCommands) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigCommands) newObject() * lib0x4595be.ConfigCommands {
	return & lib0x4595be.ConfigCommands {}
}

func (inst * comFactory4pComConfigCommands) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigCommands {
	return instance.Get().(*lib0x4595be.ConfigCommands)
}

func (inst * comFactory4pComConfigCommands) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigCommands) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigCommands) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigCommands) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigCommands) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigCommands) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigConnectors : the factory of component: com14-lib0x4595be.ConfigConnectors
type comFactory4pComConfigConnectors struct {

    mPrototype * lib0x4595be.ConfigConnectors

	
	mConnectorsSelector config.InjectionSelector

}

func (inst * comFactory4pComConfigConnectors) init() application.ComponentFactory {

	
	inst.mConnectorsSelector = config.NewInjectionSelector(".pktline-connector-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigConnectors) newObject() * lib0x4595be.ConfigConnectors {
	return & lib0x4595be.ConfigConnectors {}
}

func (inst * comFactory4pComConfigConnectors) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigConnectors {
	return instance.Get().(*lib0x4595be.ConfigConnectors)
}

func (inst * comFactory4pComConfigConnectors) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigConnectors) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigConnectors) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigConnectors) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigConnectors) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigConnectors) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Connectors = inst.getterForFieldConnectorsSelector(context)
	return context.LastError()
}

//getterForFieldConnectorsSelector
func (inst * comFactory4pComConfigConnectors) getterForFieldConnectorsSelector (context application.InstanceContext) []pktline0xd37953.ConnectorRegistry {
	list1 := inst.mConnectorsSelector.GetList(context)
	list2 := make([]pktline0xd37953.ConnectorRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(pktline0xd37953.ConnectorRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPGitConnectorReg : the factory of component: com15-lib0x4595be.HTTPGitConnectorReg
type comFactory4pComHTTPGitConnectorReg struct {

    mPrototype * lib0x4595be.HTTPGitConnectorReg

	

}

func (inst * comFactory4pComHTTPGitConnectorReg) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPGitConnectorReg) newObject() * lib0x4595be.HTTPGitConnectorReg {
	return & lib0x4595be.HTTPGitConnectorReg {}
}

func (inst * comFactory4pComHTTPGitConnectorReg) castObject(instance application.ComponentInstance) * lib0x4595be.HTTPGitConnectorReg {
	return instance.Get().(*lib0x4595be.HTTPGitConnectorReg)
}

func (inst * comFactory4pComHTTPGitConnectorReg) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPGitConnectorReg) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPGitConnectorReg) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPGitConnectorReg) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGitConnectorReg) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGitConnectorReg) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigContextBase : the factory of component: com16-lib0x4595be.ConfigContextBase
type comFactory4pComConfigContextBase struct {

    mPrototype * lib0x4595be.ConfigContextBase

	

}

func (inst * comFactory4pComConfigContextBase) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigContextBase) newObject() * lib0x4595be.ConfigContextBase {
	return & lib0x4595be.ConfigContextBase {}
}

func (inst * comFactory4pComConfigContextBase) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigContextBase {
	return instance.Get().(*lib0x4595be.ConfigContextBase)
}

func (inst * comFactory4pComConfigContextBase) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigContextBase) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigContextBase) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigContextBase) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigContextBase) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigContextBase) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigContextWithInstructions : the factory of component: com17-lib0x4595be.ConfigContextWithInstructions
type comFactory4pComConfigContextWithInstructions struct {

    mPrototype * lib0x4595be.ConfigContextWithInstructions

	
	mInstructionsSelector config.InjectionSelector

}

func (inst * comFactory4pComConfigContextWithInstructions) init() application.ComponentFactory {

	
	inst.mInstructionsSelector = config.NewInjectionSelector(".git-instruction-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigContextWithInstructions) newObject() * lib0x4595be.ConfigContextWithInstructions {
	return & lib0x4595be.ConfigContextWithInstructions {}
}

func (inst * comFactory4pComConfigContextWithInstructions) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigContextWithInstructions {
	return instance.Get().(*lib0x4595be.ConfigContextWithInstructions)
}

func (inst * comFactory4pComConfigContextWithInstructions) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigContextWithInstructions) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigContextWithInstructions) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigContextWithInstructions) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigContextWithInstructions) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigContextWithInstructions) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Instructions = inst.getterForFieldInstructionsSelector(context)
	return context.LastError()
}

//getterForFieldInstructionsSelector
func (inst * comFactory4pComConfigContextWithInstructions) getterForFieldInstructionsSelector (context application.InstanceContext) []store0x8467b3.ServiceRegistry {
	list1 := inst.mInstructionsSelector.GetList(context)
	list2 := make([]store0x8467b3.ServiceRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(store0x8467b3.ServiceRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigCore : the factory of component: com18-lib0x4595be.ConfigCore
type comFactory4pComConfigCore struct {

    mPrototype * lib0x4595be.ConfigCore

	

}

func (inst * comFactory4pComConfigCore) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigCore) newObject() * lib0x4595be.ConfigCore {
	return & lib0x4595be.ConfigCore {}
}

func (inst * comFactory4pComConfigCore) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigCore {
	return instance.Get().(*lib0x4595be.ConfigCore)
}

func (inst * comFactory4pComConfigCore) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigCore) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigCore) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigCore) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigCore) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigCore) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigInstructions : the factory of component: com19-lib0x4595be.ConfigInstructions
type comFactory4pComConfigInstructions struct {

    mPrototype * lib0x4595be.ConfigInstructions

	

}

func (inst * comFactory4pComConfigInstructions) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfigInstructions) newObject() * lib0x4595be.ConfigInstructions {
	return & lib0x4595be.ConfigInstructions {}
}

func (inst * comFactory4pComConfigInstructions) castObject(instance application.ComponentInstance) * lib0x4595be.ConfigInstructions {
	return instance.Get().(*lib0x4595be.ConfigInstructions)
}

func (inst * comFactory4pComConfigInstructions) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfigInstructions) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfigInstructions) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfigInstructions) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigInstructions) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfigInstructions) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGitlibAgent : the factory of component: git-lib-agent
type comFactory4pComGitlibAgent struct {

    mPrototype * lib0x4595be.GitlibAgent

	
	mCLISelector config.InjectionSelector
	mContextConfigurersSelector config.InjectionSelector
	mCoreConfigurersSelector config.InjectionSelector

}

func (inst * comFactory4pComGitlibAgent) init() application.ComponentFactory {

	
	inst.mCLISelector = config.NewInjectionSelector("#cli",nil)
	inst.mContextConfigurersSelector = config.NewInjectionSelector(".git-context-configurer",nil)
	inst.mCoreConfigurersSelector = config.NewInjectionSelector(".git-core-configurer",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGitlibAgent) newObject() * lib0x4595be.GitlibAgent {
	return & lib0x4595be.GitlibAgent {}
}

func (inst * comFactory4pComGitlibAgent) castObject(instance application.ComponentInstance) * lib0x4595be.GitlibAgent {
	return instance.Get().(*lib0x4595be.GitlibAgent)
}

func (inst * comFactory4pComGitlibAgent) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGitlibAgent) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGitlibAgent) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGitlibAgent) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitlibAgent) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGitlibAgent) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.CLI = inst.getterForFieldCLISelector(context)
	obj.ContextConfigurers = inst.getterForFieldContextConfigurersSelector(context)
	obj.CoreConfigurers = inst.getterForFieldCoreConfigurersSelector(context)
	return context.LastError()
}

//getterForFieldCLISelector
func (inst * comFactory4pComGitlibAgent) getterForFieldCLISelector (context application.InstanceContext) cli0xf7c71e.CLI {

	o1 := inst.mCLISelector.GetOne(context)
	o2, ok := o1.(cli0xf7c71e.CLI)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "git-lib-agent")
		eb.Set("field", "CLI")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf7c71e.CLI")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextConfigurersSelector
func (inst * comFactory4pComGitlibAgent) getterForFieldContextConfigurersSelector (context application.InstanceContext) []store0x8467b3.ContextConfigurer {
	list1 := inst.mContextConfigurersSelector.GetList(context)
	list2 := make([]store0x8467b3.ContextConfigurer, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(store0x8467b3.ContextConfigurer)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldCoreConfigurersSelector
func (inst * comFactory4pComGitlibAgent) getterForFieldCoreConfigurersSelector (context application.InstanceContext) []store0x8467b3.CoreConfigurer {
	list1 := inst.mCoreConfigurersSelector.GetList(context)
	list2 := make([]store0x8467b3.CoreConfigurer, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(store0x8467b3.CoreConfigurer)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}




