package algorithms

import (
	"io"

	"github.com/bitwormhole/gitlib/git/store"
)

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

// NewReader ...
func (inst *CompressionPlain) NewReader(r io.Reader) (io.ReadCloser, error) {
	return r.(io.ReadCloser), nil
}

// NewWriter ...
func (inst *CompressionPlain) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return w.(io.WriteCloser), nil
}
