package element

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/tests"
)

const (
	keyZipFile = "zipfile"
	keyPwd     = "pwd"
)

type TestCaseRepoInfo struct {
	// properties
	ZipFilePath string `${repo.name.zipfile}` // path in res
	PwdPath     string `${repo.name.pwd}`     // path in fs

	// output
	Name string
	PWD  fs.Path
}

type TestingRepositoriesLoader struct {
	app application.Context
	t   tests.TestContext
}

func (inst *TestingRepositoriesLoader) Init(ac application.Context, t tests.TestContext) {
	inst.app = ac
	inst.t = t
}

func (inst *TestingRepositoriesLoader) Load(config string) ([]*TestCaseRepoInfo, error) {
	srclist, err := inst.loadRepoInfoList(config)
	if err != nil {
		return nil, err
	}

	for _, src := range srclist {
		err := inst.expandZipFile(src)
		if err != nil {
			return nil, err
		}

	}
	return srclist, nil
}

func (inst *TestingRepositoriesLoader) expandZipFile(info *TestCaseRepoInfo) error {

	res := inst.app.GetResources()
	loader := &tests.TestDirectoryLoader{}
	loader.Init(res, inst.t)

	targetPath, err := loader.LoadFromZipFile(info.ZipFilePath)
	if err != nil {
		return err
	}

	info.PWD = targetPath.GetChild("./" + info.PwdPath)
	return nil
}

func (inst *TestingRepositoriesLoader) loadRepoInfoList(config string) ([]*TestCaseRepoInfo, error) {

	text, err := inst.app.GetResources().GetText(config)
	if err != nil {
		return nil, err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return nil, err
	}

	results := make([]*TestCaseRepoInfo, 0)
	table := props.Export(nil)
	const prefix = "repo."
	const suffix = ".zipfile"

	for key := range table {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			name := inst.parseGroupName(key, prefix, suffix)
			item := &TestCaseRepoInfo{}
			err := inst.parseRepoInfo(name, props, item)
			if err != nil {
				return nil, err
			}
			item.Name = name
			results = append(results, item)
		}
	}

	return results, nil
}

func (inst *TestingRepositoriesLoader) parseRepoInfo(name string, props collection.Properties, info *TestCaseRepoInfo) error {

	prefix := "repo." + name + "."
	zip, err1 := props.GetPropertyRequired(prefix + keyZipFile)
	pwd, err3 := props.GetPropertyRequired(prefix + keyPwd)
	errs := []error{err1, err3}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	info.PwdPath = pwd
	info.ZipFilePath = zip
	return nil
}

func (inst *TestingRepositoriesLoader) parseGroupName(key, prefix, suffix string) string {
	len1 := len(prefix)
	len2 := len(suffix)
	len9 := len(key)
	return key[len1 : len9-len2]
}
