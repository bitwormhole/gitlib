package clients

import (
	"github.com/bitwormhole/gitlib/git"
)

// Intent ...
type Intent struct {
	Action Action
	URL    string

	MergeRef  git.ReferenceName // like 'refs/heads/main'
	FetchRef  git.ReferenceName // like 'refs/remotes/origin/main'
	RemoteRef git.ReferenceName // like 'refs/heads/main'
}
