package store

import (
	"context"
)

////////////////////////////////////////////////////////////////////////////////

// SetupBinding ...
func SetupBinding(c context.Context) context.Context {
	if c == nil {
		c = context.Background()
	}
	const key contextValueKey = keyBindingBinding
	o1 := c.Value(key)
	if o1 != nil {
		holder, ok := o1.(*Binding)
		if ok && holder != nil {
			return c
		}
	}
	binding := &Binding{}
	return context.WithValue(c, key, binding)
}

// GetBinding ...
func GetBinding(c context.Context) *Binding {
	if c == nil {
		panic("context is nil")
	}
	const key contextValueKey = keyBindingBinding
	o1 := c.Value(key)
	if o1 != nil {
		holder, ok := o1.(*Binding)
		if ok && holder != nil {
			return holder
		}
	}
	panic("need store.SetupBinding() first")
}

// GetLib ...
func GetLib(ctx context.Context) Lib {
	binding := GetBinding(ctx)
	return binding.GetLib()
}

////////////////////////////////////////////////////////////////////////////////

// contextValueKey 用来表示上下文中键值对的 key 类型
type contextValueKey string

const keyBindingBinding contextValueKey = "github.com/bitwormhole/gitlib/git/repository/Binding#binding"

////////////////////////////////////////////////////////////////////////////////

// Binding ...
type Binding struct {
	context *Context
	factory ContextFactory
	config  *ContextConfiguration
}

// Config ...
func (inst *Binding) Config(cfg *ContextConfiguration) {
	inst.factory = cfg.Factory
	inst.config = cfg
}

// GetLib ...
func (inst *Binding) GetLib() Lib {
	l := inst.getContext().Lib
	l.FS() // panic if nil
	return l
}

func (inst *Binding) getContext() *Context {
	ctx := inst.context
	if ctx != nil {
		return ctx
	}
	factory := inst.factory
	if factory == nil {
		panic("no store.ContextFactory, call (store.Binding).SetContextFactory first")
	}
	ctx = factory.Create(inst.config)
	inst.context = ctx
	return ctx
}

////////////////////////////////////////////////////////////////////////////////
