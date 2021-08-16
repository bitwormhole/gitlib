package util

import (
	"errors"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

// LoadProperties 从文件加载属性
func LoadProperties(file fs.Path, dst collection.Properties) (collection.Properties, error) {
	if file == nil {
		return nil, errors.New("file==nil")
	}
	if dst == nil {
		dst = collection.CreateProperties()
	}
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, dst)
}

// SaveProperties 保存属性到文件
func SaveProperties(p collection.Properties, file fs.Path, mkdirs bool) error {
	if p == nil {
		return errors.New("properties==nil")
	}
	if file == nil {
		return errors.New("file==nil")
	}
	text := collection.FormatPropertiesWithSegment(p)
	return file.GetIO().WriteText(text, nil, mkdirs)
}
