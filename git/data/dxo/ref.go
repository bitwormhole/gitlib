package dxo

// ReferenceName is the name for .git/refs/*
type ReferenceName string

func (v ReferenceName) String() string {
	return string(v)
}

// Normalize ...
func (v ReferenceName) Normalize() ReferenceName {
	return "todo..."
}
