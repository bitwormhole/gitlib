package gitconfig

// 定义仓库配置名称 (remote.*)
const (
	RemoteNameURL     KeyTemplate = "remote.<name>.url"
	RemoteNamePushURL KeyTemplate = "remote.<name>.pushurl"
	RemoteNameProxy   KeyTemplate = "remote.<name>.proxy"
	RemoteNameFetch   KeyTemplate = "remote.<name>.fetch"
	RemoteNamePush    KeyTemplate = "remote.<name>.push"
	RemoteNameMirror  KeyTemplate = "remote.<name>.mirror"
)
