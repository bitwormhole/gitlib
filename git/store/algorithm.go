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
	Name() string
	Type() AlgorithmType
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
	Map(base afs.Path, id dxo.ObjectID) afs.Path
}
