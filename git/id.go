package git

// ID 表示git的一个 hash 值
type ID string

func (id ID) String() string {
	return string(id)
}

// Bytes ...
func (id ID) Bytes() []byte {
	return nil
}

// Size ...
func (id ID) Size() int {
	cb := len(id)
	return cb * 4
}

////////////////////////////////////////////////////////////////////////////////

// ObjectID 表示git的对象ID
type ObjectID ID

// PackID 表示git的包ID
type PackID ID
