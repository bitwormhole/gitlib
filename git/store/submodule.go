package store

// Submodule ...
type Submodule interface {
	Name() string
	URL() string
	Path() string
	IsActive() bool

	Workspace() Workspace
	Exists() bool
}

// Submodules ...
type Submodules interface {
	Get(name string) Submodule
	List() []Submodule
}
