package store

import (
	"context"
	"errors"
)

////////////////////////////////////////////////////////////////////////////////

// contextValueKey 用来表示上下文中键值对的 key 类型
type tagBindingKey string

const theBindingKey tagBindingKey = "github.com/bitwormhole/gitlib/git/repository/Binding#binding"

// Binding ...
type Binding interface {
	SetLib(lib Lib) error
	GetLib() (Lib, error)
}

////////////////////////////////////////////////////////////////////////////////

// Bind ...
func Bind(cc context.Context) context.Context {
	if cc == nil {
		cc = context.Background()
	}
	const key = theBindingKey
	o1, ok := cc.Value(key).(*binding)
	if ok && o1 != nil {
		return cc
	}
	o1 = &binding{}
	return context.WithValue(cc, key, o1)
}

// GetBinding ...
func GetBinding(cc context.Context) (Binding, error) {
	if cc == nil {
		cc = context.Background()
	}
	const key = theBindingKey
	o1, ok := cc.Value(key).(*binding)
	if ok && o1 != nil {
		return o1, nil
	}
	return nil, errors.New("")
}

////////////////////////////////////////////////////////////////////////////////

// binding ...
type binding struct {
	lib Lib
}

func (inst *binding) _Impl() Binding {
	return inst
}

func (inst *binding) SetLib(l Lib) error {
	if l == nil {
		return errors.New("the param 'Lib' is nil")
	}
	inst.lib = l
	return nil
}

func (inst *binding) GetLib() (Lib, error) {
	l := inst.lib
	if l == nil {
		return nil, errors.New("no Lib in this binding")
	}
	return l, nil
}

////////////////////////////////////////////////////////////////////////////////
