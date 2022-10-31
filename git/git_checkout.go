package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/instructions"
)

// Checkout 表示一条git命令
type Checkout struct {
	instructions.Meta

	Service CheckoutService

	// Path string
}

// Run ...
func (inst *Checkout) Run() error {
	return inst.Service.Run(inst)
}

// GetMeta ...
func (inst *Checkout) GetMeta() *instructions.Meta {
	return &inst.Meta
}

////////////////////////////////////////////////////////////////////////////////

// CheckoutService 。。。
type CheckoutService interface {
	Run(task *Checkout) error
}

////////////////////////////////////////////////////////////////////////////////

// NewCheckout ...
func NewCheckout(c context.Context) *Checkout {
	cmd := &Checkout{}
	cmd.Context = c
	cmd.Name = instructions.GitCheckout
	cmd.Service = findService(&cmd.Meta).(CheckoutService)
	return cmd
}
