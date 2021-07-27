package git

type GitId interface {
	String() string
	Bytes() []byte
}

type ObjectId interface {
	GitId
}

type PackId interface {
	GitId
}
