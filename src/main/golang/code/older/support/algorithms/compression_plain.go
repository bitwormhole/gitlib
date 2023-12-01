package algorithms

import (
	"io"

	"github.com/bitwormhole/gitlib/git"
)

// CompressionPlain  ...
type CompressionPlain struct {
}

func (inst *CompressionPlain) _Impl() (git.AlgorithmRegistry, git.Compression) {
	return inst, inst
}

// ListRegistrations ...
func (inst *CompressionPlain) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *CompressionPlain) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "plain",
		Type:     git.AlgorithmCompression,
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
