package algorithms

import (
	"hash"
	"hash/crc32"

	"github.com/bitwormhole/gitlib/git"
)

// DigestCRC32 ...
type DigestCRC32 struct {
}

func (inst *DigestCRC32) _Impl() (git.AlgorithmRegistry, git.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestCRC32) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestCRC32) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "crc32",
		Type:     git.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestCRC32) New() hash.Hash {
	t := crc32.MakeTable(crc32.IEEE)
	return crc32.New(t)
}

// Size ...
func (inst *DigestCRC32) Size() git.HashSize {
	const size = 4 // *8 = 32
	return git.HashSizeInBytes(size)
}
