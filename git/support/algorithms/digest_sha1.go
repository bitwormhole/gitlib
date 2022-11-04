package algorithms

import "github.com/bitwormhole/gitlib/git/store"

// DigestSHA1  ...
type DigestSHA1 struct {
}

func (inst *DigestSHA1) _Impl() (store.AlgorithmRegistry, store.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA1) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA1) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "sha1",
		Type:     store.AlgorithmDigest,
		Provider: inst,
	}
}
