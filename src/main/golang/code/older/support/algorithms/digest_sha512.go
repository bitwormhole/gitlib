package algorithms

import (
	"crypto/sha512"
	"hash"

	"github.com/bitwormhole/gitlib/git"
)

// DigestSHA512 ...
type DigestSHA512 struct {
}

func (inst *DigestSHA512) _Impl() (git.AlgorithmRegistry, git.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA512) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA512) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "sha512",
		Type:     git.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestSHA512) New() hash.Hash {
	return sha512.New()
}

// Size ...
func (inst *DigestSHA512) Size() git.HashSize {
	const size = 512 / 8
	return git.HashSizeInBytes(size)
}
