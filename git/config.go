package git

// Config 表示基本的配置文件
type Config interface {
	Parent() Config

	Import(src map[string]string)
	Export() map[string]string

	GetProperty(name string) string
	SetProperty(name, value string)

	GetHybridConfig() HybridConfig

	Save() error
	Reload() error
}

// SystemConfig 表示系统的配置文件
type SystemConfig interface {
	Config
}

// UserConfig 表示用户的配置文件
type UserConfig interface {
	Config
}

// RepositoryConfig 表示仓库的配置文件
type RepositoryConfig interface {
	Config
}

// HybridConfig 表示混合在一起的配置
type HybridConfig interface {
	Config
}
