package algorithms

import (
	"crypto/sha1"
	"hash"

	"github.com/bitwormhole/gitlib/git"
)

// DigestSHA1  ...
type DigestSHA1 struct {
}

func (inst *DigestSHA1) _Impl() (git.AlgorithmRegistry, git.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA1) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA1) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "sha1",
		Type:     git.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestSHA1) New() hash.Hash {
	return sha1.New()
}

// Size ...
func (inst *DigestSHA1) Size() git.HashSize {
	const size = 160 / 8
	return git.HashSizeInBytes(size)
}
