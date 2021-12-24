package git

// ID 表示git的一个 hash 值
type ID interface {
	Bytes() []byte

	GetBytes(buffer []byte) []byte

	GetFactory() IdentityFactory

	String() string
}

// ObjectID 表示git的对象ID
type ObjectID interface {
	ID
}

// PackID 表示git的包ID
type PackID interface {
	ID
}

// IdentityFactory 表示一个创建git-ID的工厂
type IdentityFactory interface {
	Algorithm() string

	// Size in bits
	Size() int

	Parse(s string) (ID, error)
	Create(b []byte) (ID, error)
}
