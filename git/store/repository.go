package store

// RepositoryProfile 表示一个存在的git仓库的视图
type RepositoryProfile interface {
	Layout() RepositoryLayout

	Config() ConfigChain

	HEAD() HEAD

	Index() Index

	Refs() Refs

	Objects() Objects

	Digest() Digest           // default="SHA-1"
	Compression() Compression // default="DEFLATE"
	PathMapping() PathMapping // default="xx/xxxx"

	OpenSession() (Session, error)
}

// Repository  ...
type Repository interface {
	RepositoryProfile
}

// RepositoryLoader ...
type RepositoryLoader interface {
	Load(l RepositoryLayout) (Repository, error)
}
