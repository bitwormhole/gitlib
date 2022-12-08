package gitlib

import (
	"context"
	"fmt"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter"
)

var theLib store.Lib

// GetLib 函数用来取Lib对象
func GetLib() store.Lib {
	lib := theLib
	if lib != nil {
		return lib
	}
	lib, err := createNewLib()
	if err != nil {
		panic(err)
	}
	theLib = lib
	return lib
}

func createNewLib() (store.Lib, error) {
	mod := Module()
	i := starter.InitApp()
	i.UseMain(mod)
	rt, err := i.RunEx()
	if err != nil {
		panic(err)
	}
	o, err := rt.Context().GetComponent("#git-lib-agent")
	if err != nil {
		return nil, err
	}
	agent, ok := o.(store.LibAgent)
	if !ok {
		return nil, fmt.Errorf("it's not a store.LibAgent")
	}
	return agent.GetLib()
}

// Bind ...
func Bind(cc context.Context) context.Context {
	return store.Bind(cc)
}
