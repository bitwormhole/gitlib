package store

import "bitwormhole.com/starter/afs"

// Workspace 表示工作区目录
type Workspace interface {

	// alias of WorkingDirectory
	Path() afs.Path

	// return the working directory
	WorkingDirectory() afs.Path

	// return the '.git' file or directory
	DotGit() afs.Path
}
