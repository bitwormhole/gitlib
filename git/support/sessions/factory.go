package sessions

import "github.com/bitwormhole/gitlib/git/store"

// Factory ...
type Factory struct {
}

func (inst *Factory) _Impl() store.SessionFactory {
	return inst
}

// OpenSession ...
func (inst *Factory) OpenSession(profile store.Repository) (store.Session, error) {
	se := &sessionImpl{repo: profile}
	err := se.open()
	if err != nil {
		return nil, err
	}
	return se, nil
}
