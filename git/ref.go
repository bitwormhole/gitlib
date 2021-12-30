package git

// ReferenceName is the name for .git/refs/*
type ReferenceName interface {
	String() string

	Normalize() ReferenceName
}

// Ref  is the key-value for .git/refs/*
type Ref interface {
	Name() ReferenceName

	Exists() bool

	GetValue() (ObjectID, error)

	SetValue(id ObjectID) error
}
