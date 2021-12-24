package localfilesys

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/files"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// LocalRepoLocator 本地仓库定位器
type LocalRepoLocator struct {
	markup.Component `id:"git-local-repository-locator"`
}

func (inst *LocalRepoLocator) _Impl() files.RepositoryLocator {
	return inst
}

// Locate 根据给出的路径确定仓库位置
func (inst *LocalRepoLocator) Locate(path fs.Path) (*files.RepositoryLocation, error) {

	location := &files.RepositoryLocation{}

	err := inst.findDotGit(location, path)
	if err != nil {
		// 可能是bare库，尝试找config文件
		err = inst.findConfigFile(location, path)
		if err != nil {
			return nil, err
		}
		location.Current = path
		return location, nil
	}

	dotgit := location.DotGit
	if dotgit.IsDir() {
		err = inst.findShellByDotGitDir(location, dotgit)
	} else if dotgit.IsFile() {
		err = inst.findShellByDotGitFile(location, dotgit)
	} else {
		return nil, errors.New("bad .git at " + dotgit.String())
	}

	if err != nil {
		return nil, err
	}
	return location, nil
}

func (inst *LocalRepoLocator) findShellByDotGitDir(location *files.RepositoryLocation, dotgit fs.Path) error {
	config := dotgit.GetChild("config")
	if !config.IsFile() {
		return errors.New("no file " + config.Path())
	}
	location.ShellDirectory = dotgit
	location.CoreDirectory = dotgit
	location.ConfigFile = config
	return nil
}

func (inst *LocalRepoLocator) findShellByDotGitFile(location *files.RepositoryLocation, dotgit fs.Path) error {

	shell, err := inst.parseDotGitFile(dotgit)
	if err != nil {
		return err
	}

	err = inst.findConfigFile(location, shell)
	if err != nil {
		return err
	}

	location.ShellDirectory = shell
	return nil
}

// parseDotGitFile (dotgit) return (shell_dir,error)
func (inst *LocalRepoLocator) parseDotGitFile(dotgit fs.Path) (fs.Path, error) {

	// in worktree  : gitdir: /abs/path/to/dir
	// in submodule : gitdir: ../.git/modules/demo

	text, err := dotgit.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	const prefix = "gitdir:"
	text = strings.TrimSpace(text)
	path := ""

	if strings.HasPrefix(text, prefix) {
		path = strings.TrimSpace(text[len(prefix):])
	} else {
		return nil, errors.New("bad .git file format: " + dotgit.Path())
	}

	var path2 fs.Path
	if strings.HasPrefix(path, ".") {
		// 相对路径
		path2 = dotgit.GetHref(path)
	} else {
		// 绝对路径
		path2 = dotgit.FileSystem().GetPath(path)
	}
	if !path2.IsDir() {
		return nil, errors.New("no directory at " + path2.Path())
	}

	return path2, nil
}

func (inst *LocalRepoLocator) findDotGit(location *files.RepositoryLocation, pwd fs.Path) error {

	const target = ".git"
	var dotgit fs.Path = nil
	p := pwd

	for ttl := git.MaxPathDepth; ttl > 0; ttl-- {
		if p == nil {
			break
		}
		t := p.GetChild(target)
		if t.Exists() {
			dotgit = t
			break
		}
		p = p.Parent()
	}

	if dotgit == nil {
		return errors.New("not a git repository (or any of the parent directories): .git")
	}

	location.Current = pwd
	location.DotGit = dotgit
	location.WorkingDirectory = dotgit.Parent()
	return nil
}

func (inst *LocalRepoLocator) findConfigFile(location *files.RepositoryLocation, path fs.Path) error {

	const target = "config"
	var config fs.Path = nil
	p := path

	for ttl := git.MaxPathDepth; ttl > 0; ttl-- {
		t := p.GetChild(target)
		if t.IsFile() {
			config = t
			break
		}
		p = p.Parent()
	}

	if config == nil {
		return errors.New("not a git repository (or any of the parent directories): .git/config")
	}

	location.ConfigFile = config
	location.CoreDirectory = config.Parent()
	return nil
}
