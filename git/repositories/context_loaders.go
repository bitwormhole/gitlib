package repositories

import "github.com/starter-go/afs"

// SystemContextLoader ...
type SystemContextLoader interface {
	Load() (*SystemContext, error)
}

// UserContextLoader ...
type UserContextLoader interface {
	Load(parent *SystemContext, home afs.Path) (*UserContext, error)
}

// RepositoryContextLoader ...
type RepositoryContextLoader interface {
	Load(parent *UserContext, layout Layout) (*RepositoryContext, error)
}

// SessionContextLoader ...
type SessionContextLoader interface {
	Load(parent *RepositoryContext) (*SessionContext, error)
}

// WorktreeContextLoader ...
type WorktreeContextLoader interface {
	Load(parent *RepositoryContext, layout Layout) (*WorktreeContext, error)
}

// SubmoduleContextLoader ...
type SubmoduleContextLoader interface {
	Load(parent *RepositoryContext, layout Layout) (*SubmoduleContext, error)
}
