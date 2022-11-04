package store

import (
	"bitwormhole.com/starter/afs"

	"github.com/bitwormhole/gitlib/git/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

// AlgorithmType 表示算法的类型
type AlgorithmType int

// 定义几种算法的类型
const (
	AlgorithmCompression AlgorithmType = 1
	AlgorithmDigest      AlgorithmType = 2
	AlgorithmPathMapping AlgorithmType = 3
)

////////////////////////////////////////////////////////////////////////////////

// Algorithm 表示抽象的算法
type Algorithm interface {
	GetInfo() *AlgorithmRegistration
}

// AlgorithmRegistration  ...
type AlgorithmRegistration struct {
	Name     string
	Type     AlgorithmType
	Provider Algorithm
}

// AlgorithmRegistry ...
// [inject:".git-algorithm-registry"]
type AlgorithmRegistry interface {
	ListRegistrations() []*AlgorithmRegistration
}

// AlgorithmManager 是用来管理各种算法的对象
type AlgorithmManager interface {
	Find(name string) (Algorithm, error)

	FindCompression(name string) (Compression, error)

	FindDigest(name string) (Digest, error)

	FindPathMapping(name string) (PathMapping, error)
}

////////////////////////////////////////////////////////////////////////////////

// Compression 压缩算法
type Compression interface {
	Algorithm
}

// Digest 摘要算法
type Digest interface {
	Algorithm
}

// PathMapping 路径映射算法
type PathMapping interface {
	Algorithm

	WithPattern(pattern string) PathMapping

	Map(base afs.Path, id dxo.ObjectID) afs.Path
}
