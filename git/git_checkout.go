package git

import (
	"context"

	"github.com/bitwormhole/gitlib/git/services"
)

// Checkout 表示一条git命令
type Checkout struct {
	services.Command

	Service CheckoutService

	// Path string
}

// Run ...
func (inst *Checkout) Run() error {
	return inst.Service.Run(inst)
}

// GetCommand ...
func (inst *Checkout) GetCommand() *services.Command {
	return &inst.Command
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
	cmd.Name = services.GitCheckout
	cmd.Service = findServiceForCommand(&cmd.Command).(CheckoutService)
	return cmd
}
