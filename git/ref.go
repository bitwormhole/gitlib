package git

// ReferenceName is the name for .git/refs/*
type ReferenceName string

func (v ReferenceName) String() string {
	return string(v)
}

func (v ReferenceName) Normalize() ReferenceName {
	return "todo..."
}

////////////////////////////////////////////////////////////////////////////////

// Ref  is the key-value for .git/refs/*
type Ref interface {
	Name() ReferenceName

	Exists() bool

	GetValue() (ObjectID, error)

	SetValue(id ObjectID) error
}

// Ref  is the key-value for .git/refs/*
type Refs interface {
	GetRef(name ReferenceName) Ref
}
