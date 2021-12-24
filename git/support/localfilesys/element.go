package localfilesys

import "github.com/bitwormhole/gitlib/git/files"

// Element  仓库元素，这个接口是可选的，有些元素可能不提供此接口
type Element interface {
	InitElement() error
}

// ElementFactory 仓库元素工厂
type ElementFactory interface {
	CreateElement(r *Repo, v *files.RepositoryView) error
}
