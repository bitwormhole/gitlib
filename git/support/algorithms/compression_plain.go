package algorithms

import "github.com/bitwormhole/gitlib/git/store"

// CompressionPlain  ...
type CompressionPlain struct {
}

func (inst *CompressionPlain) _Impl() (store.AlgorithmRegistry, store.Compression) {
	return inst, inst
}

// ListRegistrations ...
func (inst *CompressionPlain) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *CompressionPlain) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "plain",
		Type:     store.AlgorithmCompression,
		Provider: inst,
	}
}
