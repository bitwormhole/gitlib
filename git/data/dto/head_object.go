package dto

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
)

// HEAD 表示".git/HEAD"文件的内容
type HEAD struct {
	value string // like 'refs/heads/main'
}

// LoadFrom 从文件加载
func (inst *HEAD) LoadFrom(file fs.Path) error {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	return inst.Parse(text)
}

// SaveTo 保存到文件
func (inst *HEAD) SaveTo(file fs.Path) error {
	text := inst.Format()
	return file.GetIO().WriteText(text, nil, true)
}

// Parse 解析文本中的值
func (inst *HEAD) Parse(s string) error {
	const prefix = "ref:"
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, prefix) {
		s2 := s[len(prefix):]
		inst.value = strings.TrimSpace(s2)
		return nil
	}
	return errors.New("bad format of HEAD")
}

// Format 格式化为文本
func (inst *HEAD) Format() string {
	return "ref: " + inst.value + "\n"
}
