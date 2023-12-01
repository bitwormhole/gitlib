package algorithms

import (
	"compress/zlib"
	"io"

	"github.com/bitwormhole/gitlib/git"
)

// CompressionDeflate  ...
type CompressionDeflate struct {
}

func (inst *CompressionDeflate) _Impl() (git.AlgorithmRegistry, git.Compression) {
	return inst, inst
}

// ListRegistrations ...
func (inst *CompressionDeflate) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *CompressionDeflate) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "deflate",
		Type:     git.AlgorithmCompression,
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
