package repositories

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/base/safe"
)

////////////////////////////////////////////////////////////////////////////////

// SystemParams ...
type SystemParams struct {
	Mode safe.Mode
}

// SystemContextLoader ...
type SystemContextLoader interface {
	Load(params *SystemParams) (*SystemContext, error)
}

////////////////////////////////////////////////////////////////////////////////

// UserParams ...
type UserParams struct {
	Mode   safe.Mode
	Home   afs.Path
	Parent *SystemContext
}

// UserContextLoader ...
type UserContextLoader interface {
	Load(params *UserParams) (*UserContext, error)
}

////////////////////////////////////////////////////////////////////////////////

// RepositoryParams ...
type RepositoryParams struct {
	Mode   safe.Mode
	Layout Layout
	Parent *UserContext
}

// RepositoryContextLoader ...
type RepositoryContextLoader interface {
	Load(params *RepositoryParams) (*RepositoryContext, error)
}

////////////////////////////////////////////////////////////////////////////////

// SessionParams ...
type SessionParams struct {
	Mode   safe.Mode
	Parent *RepositoryContext
}

// SessionContextLoader ...
type SessionContextLoader interface {
	Load(params *SessionParams) (*SessionContext, error)
}

////////////////////////////////////////////////////////////////////////////////

// WorktreeParams ...
type WorktreeParams struct {
	Mode   safe.Mode
	Layout Layout
	Parent *RepositoryContext
}

// WorktreeContextLoader ...
type WorktreeContextLoader interface {
	Load(params *WorktreeParams) (*WorktreeContext, error)
}

////////////////////////////////////////////////////////////////////////////////

// SubmoduleParams ...
type SubmoduleParams struct {
	Mode   safe.Mode
	Layout Layout
	Parent *RepositoryContext
}

// SubmoduleContextLoader ...
type SubmoduleContextLoader interface {
	Load(params *SubmoduleParams) (*SubmoduleContext, error)
}

////////////////////////////////////////////////////////////////////////////////
