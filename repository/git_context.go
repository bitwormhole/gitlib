package repository

type GitContext interface {
	PWD() GitPWD
	Repository() GitRepository
	Worktree() GitWorktree
}

type GitContextClient interface {
	Context() GitContext
}
