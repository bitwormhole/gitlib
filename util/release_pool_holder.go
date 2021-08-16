package util

import "github.com/bitwormhole/starter/lang"

// ReleasePoolHolder 是 ReleasePool 的控制对象
type ReleasePoolHolder struct {
	pool lang.ReleasePool
}

func (inst *ReleasePoolHolder) Init(pool lang.ReleasePool) {
	inst.pool = pool
}

func (inst *ReleasePoolHolder) Disconnect() {
	inst.pool = nil
}

func (inst *ReleasePoolHolder) Release() error {
	pool := inst.pool
	if pool == nil {
		return nil
	}
	inst.pool = nil
	return pool.Release()
}
