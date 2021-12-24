package dto

import (
	"strings"

	"github.com/bitwormhole/starter/io/fs"
)

// Ref 表示 ".git/refs/.../xxx" 文件的内容
type Ref struct {
	value string
}

// LoadFrom 从文件加载
func (inst *Ref) LoadFrom(file fs.Path) error {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	return inst.Parse(text)
}

// SaveTo 保存到文件
func (inst *Ref) SaveTo(file fs.Path) error {
	text := inst.Format()
	return file.GetIO().WriteText(text, nil, true)
}

// Parse 解析文本中的值
func (inst *Ref) Parse(s string) error {
	inst.value = strings.TrimSpace(s)
	return nil
}

// Format 格式化为文本
func (inst *Ref) Format() string {
	return inst.value + "\n"
}
