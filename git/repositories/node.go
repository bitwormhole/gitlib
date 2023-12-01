package repositories

import "github.com/starter-go/afs"

// NodeType 表示Node的类型
type NodeType string

// 定义各种 Node 的类型
const (
	NodeConfig    NodeType = ".git/config"
	NodeObjects   NodeType = ".git/objects"
	NodeDotGit    NodeType = ".git"
	NodeWorkspace NodeType = ".git/.."
	NodeIndex     NodeType = ".git/index"
	NodeHEAD      NodeType = ".git/HEAD"
)

// NodeInfo 表示 git 中的一个 dir 或者 file
type NodeInfo struct {
	Name string
	Path afs.FS
	Type NodeType
}

// Node  提供节点的路径
type Node interface {
	Path() afs.Path
	NodeType() NodeType
}
