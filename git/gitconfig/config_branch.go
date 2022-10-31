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
