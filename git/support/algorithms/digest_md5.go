package algorithms

import (
	"crypto/md5"
	"hash"

	"github.com/bitwormhole/gitlib/git"
)

// DigestMD5 ...
type DigestMD5 struct {
}

func (inst *DigestMD5) _Impl() (git.AlgorithmRegistry, git.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestMD5) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestMD5) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "md5",
		Type:     git.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestMD5) New() hash.Hash {
	return md5.New()
}

// Size ...
func (inst *DigestMD5) Size() git.HashSize {
	const size = 128 / 8
	return git.HashSizeInBytes(size)
}
