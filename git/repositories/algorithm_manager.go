package repositories

import (
	"github.com/bitwormhole/gitlib/git"
)

// AlgorithmManager 是用来管理各种算法的对象
type AlgorithmManager interface {
	Find(name string) (git.Algorithm, error)

	FindCompression(name string) (git.Compression, error)

	FindDigest(name string) (git.Digest, error)

	FindPathMapping(name string) (git.PathMapping, error)
}
