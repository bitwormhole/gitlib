package algorithms

import "github.com/bitwormhole/gitlib/git/store"

// CompressionDeflate  ...
type CompressionDeflate struct {
}

func (inst *CompressionDeflate) _Impl() (store.AlgorithmRegistry, store.Compression) {
	return inst, inst
}

// ListRegistrations ...
func (inst *CompressionDeflate) ListRegistrations() []*store.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*store.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *CompressionDeflate) GetInfo() *store.AlgorithmRegistration {
	return &store.AlgorithmRegistration{
		Name:     "deflate",
		Type:     store.AlgorithmCompression,
		Provider: inst,
	}
}
