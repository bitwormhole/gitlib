package services

import (
	"errors"
	"io/fs"
	"strconv"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/services"
)

// GitInitService ...
type GitInitService struct {
}

func (inst *GitInitService) _Impl() (git.InitService, services.ServiceRegistry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GitInitService) ListRegistrations() []*services.ServiceRegistration {
	name := inst.Name()
	reg := &services.ServiceRegistration{
		Name:    name,
		Service: inst,
	}
	return []*services.ServiceRegistration{reg}
}

// Name ...
func (inst *GitInitService) Name() string {
	return services.GitInit
}

// Run ...
func (inst *GitInitService) Run(task *git.Init) error {

	dotgit, err := inst.getDotGitDir(task)
	if err != nil {
		return err
	}

	builder := gitInitServiceRepoBuilder{dotgit: dotgit, task: task}
	builder.addDir("hooks")
	builder.addDir("info")
	builder.addDir("objects")
	builder.addDir("objects/info")
	builder.addDir("objects/pack")
	builder.addDir("refs")
	builder.addDir("refs/heads")
	builder.addDir("refs/tags")

	fileConf := builder.addFile("config")
	fileDesc := builder.addFile("description")
	fileHead := builder.addFile("HEAD")

	err = builder.mkdirs()
	if err != nil {
		return err
	}

	err = builder.mkfiles()
	if err != nil {
		return err
	}

	err = builder.initFileConfig(fileConf)
	if err != nil {
		return err
	}

	err = builder.initFileDescription(fileDesc)
	if err != nil {
		return err
	}

	err = builder.initFileHead(fileHead)
	if err != nil {
		return err
	}

	return nil
}

func (inst *GitInitService) getDotGitDir(task *git.Init) (afs.Path, error) {

	wd := task.WD
	bare := task.Bare
	dirName := task.Directory
	dotgit := wd

	if !bare {
		dotgit = wd.GetChild(dirName + "/.git")
	} else {
		dotgit = wd.GetChild(dirName)
	}

	if !wd.IsDirectory() {
		path := wd.GetPath()
		return nil, errors.New("the working directory is NOT exists, path=" + path)
	}

	if dotgit.Exists() {
		path := dotgit.GetPath()
		return nil, errors.New("the repository directory is exists, path=" + path)
	}

	return dotgit, nil
}

////////////////////////////////////////////////////////////////////////////////

type gitInitServiceRepoBuilder struct {
	task   *git.Init
	dotgit afs.Path
	files  []afs.Path
	dirs   []afs.Path
}

func (inst *gitInitServiceRepoBuilder) addFile(path string) afs.Path {
	p := inst.dotgit.GetChild(path)
	inst.files = append(inst.files, p)
	return p
}

func (inst *gitInitServiceRepoBuilder) addDir(path string) afs.Path {
	p := inst.dotgit.GetChild(path)
	inst.dirs = append(inst.dirs, p)
	return p
}

func (inst *gitInitServiceRepoBuilder) mkdirs() error {
	opt := &afs.Options{Create: true, Mkdirs: true}
	list := inst.dirs
	for _, dir := range list {
		err := dir.Mkdirs(opt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *gitInitServiceRepoBuilder) mkfiles() error {
	text := ""
	opt := &afs.Options{
		Create:     true,
		Mkdirs:     true,
		Permission: fs.ModePerm,
	}
	list := inst.files
	for _, file := range list {
		err := file.GetIO().WriteText(text, opt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *gitInitServiceRepoBuilder) initFileConfig(path afs.Path) error {

	const nl = "\n"
	builder := strings.Builder{}
	task := inst.task
	bare := strconv.FormatBool(task.Bare)

	builder.WriteString("[core]" + nl)
	builder.WriteString("	repositoryformatversion = 0" + nl)
	builder.WriteString("	filemode = false" + nl)
	builder.WriteString("	logallrefupdates = true" + nl)
	builder.WriteString("	symlinks = false" + nl)
	builder.WriteString("	ignorecase = true" + nl)
	builder.WriteString("	bare = " + bare + nl)

	text := builder.String()
	return path.GetIO().WriteText(text, nil)
}

func (inst *gitInitServiceRepoBuilder) initFileDescription(path afs.Path) error {
	text := "Unnamed repository; edit this file 'description' to name the repository.\n"
	return path.GetIO().WriteText(text, nil)
}

func (inst *gitInitServiceRepoBuilder) initFileHead(path afs.Path) error {
	text := "ref: refs/heads/master\n"
	return path.GetIO().WriteText(text, nil)
}
