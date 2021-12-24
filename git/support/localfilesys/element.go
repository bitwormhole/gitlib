package localfilesys

import "github.com/bitwormhole/gitlib/git/files"

// ElementFactory 仓库元素工厂
type ElementFactory interface {
	InitElement(r *Repo, v *files.RepositoryView) error
}
