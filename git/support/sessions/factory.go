package sessions

import "github.com/bitwormhole/gitlib/git/store"

// Factory ...
type Factory struct {
}

func (inst *Factory) _Impl() store.SessionFactory {
	return inst
}

// OpenSession ...
func (inst *Factory) OpenSession(profile store.RepositoryProfile) (store.Session, error) {

	se := &sessionImpl{profile: profile}

	return se, nil
}
