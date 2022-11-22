package gitconfig

// 定义仓库配置名称 (branch.*)
const (
	BranchSort             KeyTemplate = "branch.sort"
	BranchNameRemote       KeyTemplate = "branch.<name>.remote"
	BranchNamePushRemote   KeyTemplate = "branch.<name>.pushRemote"
	BranchNameMerge        KeyTemplate = "branch.<name>.merge"
	BranchNameMergeOptions KeyTemplate = "branch.<name>.mergeOptions"
	BranchNameRebase       KeyTemplate = "branch.<name>.rebase"
	BranchNameDescription  KeyTemplate = "branch.<name>.description"
)

// Branch 表示 config["branch.<name>.*"]
type Branch struct {
	Name   string
	Exists bool

	Merge       string // ='refs/heads/main'
	Remote      string // the remote name
	PushRemote  string
	Description string
}
