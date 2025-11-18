package repositories

import "github.com/starter-go/v1/afs"

type FileNode interface {
	GetOwnerRepository() Repository

	GetFile() afs.File
}

type DirNode interface {
	GetOwnerRepository() Repository

	GetDir() afs.Directory
}
