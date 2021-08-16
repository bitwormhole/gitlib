package repo

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/gitlib/util"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
)

// FileRepositoryLocator 是默认的仓库定位器
type FileRepositoryLocator struct {
}

func (inst *FileRepositoryLocator) _Impl() repository.Locator {
	return inst
}

// Locate 根据给出的路径，定位仓库的准确位置
func (inst *FileRepositoryLocator) Locate(uri lang.URI) (*repository.Location, error) {

	path, err := fs.Default().GetPathByURI(uri)
	if err != nil {
		return nil, err
	}

	location := &repository.Location{}
	err = inst.findPwd(path, location)
	if err != nil {
		return nil, err
	}

	err = inst.findDotGit(location)
	if err != nil {
		// 如果找不到.git, 尝试作为bare库搜索
		core2, err2 := inst.tryFindCoreDirAsBare(location.PWD)
		if err2 == nil {
			location.CoreDirectory = core2
			return location, nil
		}
		return nil, err
	}

	err = inst.findWorkingDir(location)
	if err != nil {
		return nil, err
	}

	err = inst.findShellDir(location)
	if err != nil {
		return nil, err
	}

	err = inst.findCoreDir(location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// Accept 判断是否支持给定的URI
func (inst *FileRepositoryLocator) Accept(uri lang.URI) bool {
	if uri == nil {
		return false
	}
	return uri.Scheme() == "file"
}

func (inst *FileRepositoryLocator) findPwd(path fs.Path, location *repository.Location) error {
	p0 := path
	for timeout := git.MaxPathDepth; path != nil; path = path.Parent() {
		if timeout > 0 {
			timeout--
		} else {
			return errors.New("The path is too deep, path=" + p0.Path())
		}
		if path.Exists() && path.IsDir() {
			location.PWD = path
			return nil
		}
		path = path.Parent()
	}
	return errors.New("No dir in path: " + p0.Path())
}

func (inst *FileRepositoryLocator) findDotGit(location *repository.Location) error {
	p0 := location.PWD
	path := p0
	for timeout := git.MaxPathDepth; path != nil; path = path.Parent() {
		if timeout > 0 {
			timeout--
		} else {
			return errors.New("The path is too deep, path=" + p0.Path())
		}
		dotgit := path.GetChild(".git")
		if dotgit.Exists() {
			location.DotGit = dotgit
			return nil
		}
	}
	return errors.New("No .git in path: " + p0.Path())
}

func (inst *FileRepositoryLocator) findWorkingDir(location *repository.Location) error {
	location.WorkingDirectory = location.DotGit.Parent()
	return nil
}

func (inst *FileRepositoryLocator) findShellDir(location *repository.Location) error {
	dotgit := location.DotGit
	if dotgit.IsDir() {
		location.WorktreeDirectory = nil
		location.SubmoduleDirectory = nil
		location.CoreDirectory = dotgit
		return nil
	} else if dotgit.IsFile() {
		shell, err := inst.parseDotGitFile(dotgit)
		if err != nil {
			return err
		}
		if inst.isInSubmodule(shell) {
			location.SubmoduleDirectory = shell
		} else if inst.isInWorktree(shell) {
			location.WorktreeDirectory = shell
		} else {
			return errors.New("bad .git node, not a submodule nor worktree, node=" + dotgit.Path())
		}
		return nil
	}
	return errors.New("bad .git node:" + dotgit.Path())
}

func (inst *FileRepositoryLocator) findCoreDir(location *repository.Location) error {
	if location.CoreDirectory != nil {
		return nil
	}
	from := location.PWD
	if location.SubmoduleDirectory != nil {
		from = location.SubmoduleDirectory.Parent()
	} else if location.WorktreeDirectory != nil {
		from = location.WorktreeDirectory.Parent()
	}
	core, err := inst.tryFindCoreDirAsBare(from)
	if err != nil {
		return err
	}
	location.CoreDirectory = core
	return nil
}

func (inst *FileRepositoryLocator) tryFindCoreDirAsBare(path fs.Path) (fs.Path, error) {
	p0 := path
	for timeout := git.MaxPathDepth; path != nil; path = path.Parent() {
		if timeout > 0 {
			timeout--
		} else {
			return nil, errors.New("The path is too deep, path=" + p0.Path())
		}
		if inst.isLayoutAsCoreDir(path) {
			return path, nil
		}
	}
	return nil, errors.New("No repo-core in path: " + p0.Path())
}

func (inst *FileRepositoryLocator) isLayoutAsCoreDir(dir fs.Path) bool {

	const minGoodValue = 8
	good := 0

	// item as file
	good += inst.checkItemInDir(dir, "config", false, 2)
	good += inst.checkItemInDir(dir, "index", false, 2)
	good += inst.checkItemInDir(dir, "HEAD", false, 2)

	// item as dir
	good += inst.checkItemInDir(dir, "refs", true, 2)
	good += inst.checkItemInDir(dir, "objects", true, 2)
	good += inst.checkItemInDir(dir, "logs", true, 1)
	good += inst.checkItemInDir(dir, "hooks", true, 1)

	if good < minGoodValue {
		return false
	}

	// check the 'config' file
	// todo ...

	return true
}

func (inst *FileRepositoryLocator) checkItemInDir(dir fs.Path, name string, isdir bool, weight int) int {
	child := dir.GetChild(name)
	if child.Exists() {
		if child.IsDir() && isdir {
			return weight
		} else if child.IsFile() && !isdir {
			return weight
		}
	}
	return 0
}

func (inst *FileRepositoryLocator) parseDotGitFile(file fs.Path) (fs.Path, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	text = strings.TrimSpace(text)
	const prefix = "gitdir:"
	if !strings.HasPrefix(text, prefix) {
		return nil, errors.New("bad .git file: " + file.Path())
	}
	gitdir := strings.TrimSpace(text[len(prefix):])
	if strings.HasPrefix(gitdir, "./") || strings.HasPrefix(gitdir, "../") {
		// 相对路径
		target := file.Parent().GetChild(gitdir)
		return target, nil
	}
	// 绝对路径
	target := file.FileSystem().GetPath(gitdir)
	return target, nil
}

func (inst *FileRepositoryLocator) isInWorktree(path fs.Path) bool {

	file1 := path.GetChild("gitdir")
	file2 := path.GetChild("commondir")
	file3 := path.GetChild("HEAD")
	file4 := path.GetChild("index")
	checkpoints := []fs.Path{file1, file2, file3, file4}

	for _, cp := range checkpoints {
		if !cp.Exists() {
			return false
		}
	}

	const token = "/worktrees/"
	return strings.Contains(path.Path(), token)
}

func (inst *FileRepositoryLocator) isInSubmodule(path fs.Path) bool {

	// check: file [config] contains key [core.worktree]
	config := path.GetChild("config")

	props, err := util.LoadProperties(config, nil)
	if err != nil {
		return false
	}

	_, err = props.GetPropertyRequired("core.worktree")
	if err != nil {
		return false
	}

	const token = "/modules/"
	return strings.Contains(path.Path(), token)
}
