package store

// RepositoryProfile 表示一个存在的git仓库的视图
type RepositoryProfile interface {
	Layout() RepositoryLayout

	Config() Config

	HEAD() HEAD

	Index() Index

	Refs() Refs

	Objects() Objects
}

// RepositoryProfileFactory ...
type RepositoryProfileFactory interface {
	Create(l RepositoryLayout) (RepositoryProfile, error)
}
