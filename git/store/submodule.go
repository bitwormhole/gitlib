package store

// Submodule ...
type Submodule interface {
	Name() string
	Workspace() Workspace
	Exists() bool
}

// Submodules ...
type Submodules interface {
	Get(name string) Submodule
	List() []Submodule
}
