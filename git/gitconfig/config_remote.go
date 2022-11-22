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

// Remote 表示 config["remote.<name>.*"]
type Remote struct {
	Name   string
	Exists bool

	URL   string // ='git@github.com:example/example.git'
	Fetch string // ='+refs/heads/*:refs/remotes/origin/*'
}

////////////////////////////////////////////////////////////////////////////////

// Remotes 表示一组 remote
type Remotes struct {
	Name string

	RemoteNames []string
}

// Group aka for Remotes
type Group Remotes

////////////////////////////////////////////////////////////////////////////////

// RemoteAndBranch  表示一组相关的 remote-branches
type RemoteAndBranch struct {
	RemoteName string
	Remote     *Remote
	Branch     *Branch
}

////////////////////////////////////////////////////////////////////////////////
