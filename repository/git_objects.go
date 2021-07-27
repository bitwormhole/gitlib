package repository

import "github.com/bitwormhole/gitlib/util"

type ObjectsPack interface {
	util.LocalDirectory
}

type ObjectsInfo interface {
	util.LocalDirectory
}

type GitObjects interface {
	util.LocalDirectory

	Info() ObjectsInfo
	Pack() ObjectsPack
}
