package git

import (
	"hash"
	"io"

	"bitwormhole.com/starter/afs"
)

// AlgorithmType 表示算法的类型
type AlgorithmType int

// 定义几种算法的类型
const (
	AlgorithmCompression AlgorithmType = 1
	AlgorithmDigest      AlgorithmType = 2
	AlgorithmPathMapping AlgorithmType = 3
)

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

////////////////////////////////////////////////////////////////////////////////

// Compression 压缩算法
type Compression interface {
	Algorithm

	NewReader(r io.Reader) (io.ReadCloser, error)
	NewWriter(w io.Writer) (io.WriteCloser, error)
}

// Digest 摘要算法
type Digest interface {
	Algorithm

	Size() HashSize

	New() hash.Hash
}

// PathMapping 路径映射算法
type PathMapping interface {
	Algorithm

	WithPattern(pattern string) PathMapping

	Map(base afs.Path, id ObjectID) afs.Path
}
