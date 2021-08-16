// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfg

import(
	errors "errors"
	repository_d59845 "github.com/bitwormhole/gitlib/repository"
	element_d197cc "github.com/bitwormhole/gitlib/src/test/golang/element"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theGitRepoLayoutTester
	cominfobuilder.Reset()
	cominfobuilder.ID("theGitRepoLayoutTester").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &element_d197cc.RepoLayoutTester{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*element_d197cc.RepoLayoutTester).Start()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theGitRepoLayoutTester{}
		adapter.instance = o.(*element_d197cc.RepoLayoutTester)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theGitRepoLocatorTester
	cominfobuilder.Reset()
	cominfobuilder.ID("theGitRepoLocatorTester").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &element_d197cc.RepoLocatorTester{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*element_d197cc.RepoLocatorTester).Start()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theGitRepoLocatorTester{}
		adapter.instance = o.(*element_d197cc.RepoLocatorTester)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theGitRepoTester
	cominfobuilder.Reset()
	cominfobuilder.ID("theGitRepoTester").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &element_d197cc.GitRepoTester{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*element_d197cc.GitRepoTester).Start()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theGitRepoTester{}
		adapter.instance = o.(*element_d197cc.GitRepoTester)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }


	return nil
}


////////////////////////////////////////////////////////////////////////////////
// type theGitRepoLayoutTester struct

func (inst *theGitRepoLayoutTester) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.AppContext=inst.__get_AppContext__(injection, "context")
	inst.RM=inst.__get_RM__(injection, "#git-repository-manager")


	// to instance
	instance.AppContext=inst.AppContext
	instance.RM=inst.RM


	// invoke custom inject method


	return injection.Close()
}

func (inst * theGitRepoLayoutTester) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theGitRepoLayoutTester) __get_RM__(injection application.Injection,selector string) repository_d59845.Manager {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(repository_d59845.Manager)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Manager")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theGitRepoLocatorTester struct

func (inst *theGitRepoLocatorTester) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.AppContext=inst.__get_AppContext__(injection, "context")
	inst.RM=inst.__get_RM__(injection, "#git-repository-manager")


	// to instance
	instance.AppContext=inst.AppContext
	instance.RM=inst.RM


	// invoke custom inject method


	return injection.Close()
}

func (inst * theGitRepoLocatorTester) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theGitRepoLocatorTester) __get_RM__(injection application.Injection,selector string) repository_d59845.Manager {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(repository_d59845.Manager)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Manager")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theGitRepoTester struct

func (inst *theGitRepoTester) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.Path=inst.__get_Path__(injection, "${test.repo.path}")
	inst.RM=inst.__get_RM__(injection, "#git-repository-manager")


	// to instance
	instance.Path=inst.Path
	instance.RM=inst.RM


	// invoke custom inject method


	return injection.Close()
}

func (inst * theGitRepoTester) __get_Path__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theGitRepoTester) __get_RM__(injection application.Injection,selector string) repository_d59845.Manager {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(repository_d59845.Manager)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Manager")
		injection.OnError(err)
		return nil
	}

	return o2

}

