package algorithms

import (
	"crypto/md5"
	"hash"

	"github.com/bitwormhole/gitlib/git/store"
)

// DigestMD5 ...
type DigestMD5 struct {
}

func (inst *DigestMD5) _Impl() (store.AlgorithmRegistry, store.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestMD5) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestMD5) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "md5",
		Type:     store.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestMD5) New() hash.Hash {
	return md5.New()
}
