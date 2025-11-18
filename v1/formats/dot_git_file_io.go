package formats

import (
	"fmt"
	"strings"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/vlog"
)

type DotGitFile struct {
	Text   string   // .git 的文本内容
	Target afs.Path // .git (file) 指向的目标位置
	File   afs.File // .git 本身
}

func (inst *DotGitFile) Read() error {

	file := inst.File
	io := file.GetIO()

	text, err := io.ReadText(file, nil)
	if err != nil {
		return err
	}

	path, err := inst.parsePath(text)
	if err != nil {
		return err
	}

	inst.Target = path
	inst.Text = text
	return nil
}

func (inst *DotGitFile) Write() error {

	file := inst.File
	io := file.GetIO()
	opt := new(afs.Options)

	text := inst.formatPath(inst.Target)
	inst.Text = text

	return io.WriteText(file, text, opt)
}

func (inst *DotGitFile) formatPath(path afs.Path) string {
	const prefix = "gitdir:"
	p2, err := path.Normalize()
	if err != nil {
		vlog.Warn("DotGitFile.formatPath() : %s", err.Error())
		p2 = path
	}
	return fmt.Sprint(prefix, p2.String(), "\n")
}

func (inst *DotGitFile) parsePath(str string) (afs.Path, error) {
	const prefix = "gitdir:"
	str = strings.TrimSpace(str)
	if strings.HasPrefix(str, prefix) {
		path := str[len(prefix):]
		path = strings.TrimSpace(path)
		return afs.Path(path), nil
	}
	return "", fmt.Errorf("bad format of .git file, content = [%s]", str)
}
