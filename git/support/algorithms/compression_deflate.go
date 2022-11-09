package algorithms

import (
	"compress/zlib"
	"io"

	"github.com/bitwormhole/gitlib/git/store"
)

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

// NewReader ...
func (inst *CompressionDeflate) NewReader(r io.Reader) (io.ReadCloser, error) {

	// 奇怪，这里用 zlib.NewReader 而不是 flate.NewReader ?

	// r2 := flate.NewReader(r)
	// return r2, nil

	return zlib.NewReader(r)
}

// NewWriter ...
func (inst *CompressionDeflate) NewWriter(w io.Writer) (io.WriteCloser, error) {

	// 奇怪，这里用 zlib.NewReader 而不是 flate.NewReader ?

	// return flate.NewWriter(w, 0)

	w2 := zlib.NewWriter(w)
	return w2, nil
}
