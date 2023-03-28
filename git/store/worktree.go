package store

// Worktree ...
type Worktree interface {
	Name() string
	Workspace() Workspace
	Exists() bool
}

// Worktrees ...
type Worktrees interface {
	Get(name string) Worktree
	List() []Worktree
}
