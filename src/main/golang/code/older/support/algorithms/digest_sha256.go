package algorithms

import (
	"crypto/sha256"
	"hash"

	"github.com/bitwormhole/gitlib/git"
)

// DigestSHA256 ...
type DigestSHA256 struct {
}

func (inst *DigestSHA256) _Impl() (git.AlgorithmRegistry, git.Digest) {
	return inst, inst
}

// ListRegistrations ...
func (inst *DigestSHA256) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *DigestSHA256) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "sha256",
		Type:     git.AlgorithmDigest,
		Provider: inst,
	}
}

// New ...
func (inst *DigestSHA256) New() hash.Hash {
	return sha256.New()
}

// Size ...
func (inst *DigestSHA256) Size() git.HashSize {
	const size = 256 / 8
	return git.HashSizeInBytes(size)
}
