package algorithms

import "github.com/bitwormhole/gitlib/git/store"

// DigestSHA256 ...
type DigestSHA256 struct {
}

func (inst *DigestSHA256) _Impl() (store.AlgorithmRegistry, store.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA256) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA256) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "sha256",
		Type:     store.AlgorithmDigest,
		Provider: inst,
	}
}
