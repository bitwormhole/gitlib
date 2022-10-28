package gitconfig

// ConfigName 表示仓库配置名称
type ConfigName string

// 定义仓库配置名称
const (
	CoreLogAllRefUpdates        ConfigName = "core.logallrefupdates"
	CoreSymlinks                ConfigName = "core.symlinks"
	CoreIgnoreCase              ConfigName = "core.ignorecase"
	CoreBare                    ConfigName = "core.bare"
	CoreRepositoryFormatVersion ConfigName = "core.repositoryformatversion"
	CoreFileMode                ConfigName = "core.filemode"
)
