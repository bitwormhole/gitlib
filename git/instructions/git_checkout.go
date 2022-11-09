package instructions

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
)

// Checkout 表示一条git命令
type Checkout struct {
	store.Meta

	Service CheckoutService

	// Path string
}

// Run ...
func (inst *Checkout) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Checkout) GetMeta() *store.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// CheckoutService 。。。
type CheckoutService interface {
	store.Service
	Run(task *Checkout) error
}

////////////////////////////////////////////////////////////////////////////////

// NewCheckout ...
func NewCheckout(c context.Context) *Checkout {
	cmd := &Checkout{}
	cmd.Context = c
	cmd.Name = GitCheckout
	cmd.Service = findService(&cmd.Meta).(CheckoutService)
	return cmd
}
