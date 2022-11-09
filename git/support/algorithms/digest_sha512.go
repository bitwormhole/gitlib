package algorithms

import (
	"crypto/sha512"
	"hash"

	"github.com/bitwormhole/gitlib/git/store"
)

// DigestSHA512 ...
type DigestSHA512 struct {
}

func (inst *DigestSHA512) _Impl() (store.AlgorithmRegistry, store.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA512) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA512) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "sha512",
		Type:     store.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestSHA512) New() hash.Hash {
	return sha512.New()
}
