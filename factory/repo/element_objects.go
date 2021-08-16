package repo

type GitObjectsElementFactory struct{}

func (inst *GitObjectsElementFactory) _Impl() ElementFactory {
	return inst
}

func (inst *GitObjectsElementFactory) Make(ctx *ViewportContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type objectsElement struct{}
