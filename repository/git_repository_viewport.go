package repository

import "io"

type GitRepositoryViewport interface {
	io.Closer
	GitContext
	GitContextClient
}
