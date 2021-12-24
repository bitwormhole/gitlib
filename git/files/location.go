package files

import "github.com/bitwormhole/starter/io/fs"

// RepositoryLocation 仓库位置
type RepositoryLocation struct {
	Current          fs.Path
	WorkingDirectory fs.Path // the parent of '.git'
	DotGit           fs.Path // the '.git' file or directory
	ConfigFile       fs.Path // the 'config' file
	ShellDirectory   fs.Path // the directory of 'HEAD' file
	CoreDirectory    fs.Path // the parent directory of 'config' file
}
