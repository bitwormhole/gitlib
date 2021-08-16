// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etcgitlib

import(
	errors "errors"
	repo_bc84d6 "github.com/bitwormhole/gitlib/factory/repo"
	repository_d59845 "github.com/bitwormhole/gitlib/repository"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theElementConfig
	cominfobuilder.Reset()
	cominfobuilder.ID("theElementConfig").Class("default-git-repository-element").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.GitConfigElementFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theElementConfig{}
		adapter.instance = o.(*repo_bc84d6.GitConfigElementFactory)
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

	// theElementHead
	cominfobuilder.Reset()
	cominfobuilder.ID("theElementHead").Class("default-git-repository-element").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.GitHeadElementFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theElementHead{}
		adapter.instance = o.(*repo_bc84d6.GitHeadElementFactory)
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

	// theElementIndex
	cominfobuilder.Reset()
	cominfobuilder.ID("theElementIndex").Class("default-git-repository-element").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.GitIndexElementFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theElementIndex{}
		adapter.instance = o.(*repo_bc84d6.GitIndexElementFactory)
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

	// theElementObjects
	cominfobuilder.Reset()
	cominfobuilder.ID("theElementObjects").Class("default-git-repository-element").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.GitObjectsElementFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theElementObjects{}
		adapter.instance = o.(*repo_bc84d6.GitObjectsElementFactory)
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

	// theElementRefs
	cominfobuilder.Reset()
	cominfobuilder.ID("theElementRefs").Class("default-git-repository-element").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.GitRefsElementFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theElementRefs{}
		adapter.instance = o.(*repo_bc84d6.GitRefsElementFactory)
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

	// theFileRepoFactory
	cominfobuilder.Reset()
	cominfobuilder.ID("file-git-repository-factory").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.FileRepositoryFactory{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theFileRepoFactory{}
		adapter.instance = o.(*repo_bc84d6.FileRepositoryFactory)
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

	// theFileRepoLocator
	cominfobuilder.Reset()
	cominfobuilder.ID("file-git-repository-locator").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.FileRepositoryLocator{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theFileRepoLocator{}
		adapter.instance = o.(*repo_bc84d6.FileRepositoryLocator)
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

	// theRepoDriverDefault
	cominfobuilder.Reset()
	cominfobuilder.ID("file-git-repository-driver").Class("git-repository-driver").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.DefaultRepositoryDriver{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRepoDriverDefault{}
		adapter.instance = o.(*repo_bc84d6.DefaultRepositoryDriver)
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

	// theRepoManager
	cominfobuilder.Reset()
	cominfobuilder.ID("git-repository-manager").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &repo_bc84d6.DefaultRepositoryManager{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRepoManager{}
		adapter.instance = o.(*repo_bc84d6.DefaultRepositoryManager)
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
// type theElementConfig struct

func (inst *theElementConfig) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theElementHead struct

func (inst *theElementHead) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theElementIndex struct

func (inst *theElementIndex) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theElementObjects struct

func (inst *theElementObjects) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theElementRefs struct

func (inst *theElementRefs) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theFileRepoFactory struct

func (inst *theFileRepoFactory) __inject__(context application.Context) error {

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
	inst.Pipeline=inst.__get_Pipeline__(injection, ".default-git-repository-element")


	// to instance
	instance.Pipeline=inst.Pipeline


	// invoke custom inject method


	return injection.Close()
}

func (inst * theFileRepoFactory) __get_Pipeline__(injection application.Injection,selector string) []repo_bc84d6.ElementFactory {
	list := make([]repo_bc84d6.ElementFactory, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(repo_bc84d6.ElementFactory)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

////////////////////////////////////////////////////////////////////////////////
// type theFileRepoLocator struct

func (inst *theFileRepoLocator) __inject__(context application.Context) error {

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


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theRepoDriverDefault struct

func (inst *theRepoDriverDefault) __inject__(context application.Context) error {

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
	inst.MyFactory=inst.__get_MyFactory__(injection, "#file-git-repository-factory")
	inst.MyLocator=inst.__get_MyLocator__(injection, "#file-git-repository-locator")


	// to instance
	instance.MyFactory=inst.MyFactory
	instance.MyLocator=inst.MyLocator


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRepoDriverDefault) __get_MyFactory__(injection application.Injection,selector string) repository_d59845.Factory {

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

	o2, ok := o1.(repository_d59845.Factory)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Factory")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theRepoDriverDefault) __get_MyLocator__(injection application.Injection,selector string) repository_d59845.Locator {

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

	o2, ok := o1.(repository_d59845.Locator)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Locator")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theRepoManager struct

func (inst *theRepoManager) __inject__(context application.Context) error {

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
	inst.Drivers=inst.__get_Drivers__(injection, ".git-repository-driver")


	// to instance
	instance.Drivers=inst.Drivers


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRepoManager) __get_Drivers__(injection application.Injection,selector string) []repository_d59845.Driver {
	list := make([]repository_d59845.Driver, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(repository_d59845.Driver)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

