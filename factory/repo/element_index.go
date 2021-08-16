package repo

type GitIndexElementFactory struct {
}

func (inst *GitIndexElementFactory) _Impl() ElementFactory {
	return inst
}

func (inst *GitIndexElementFactory) Make(ctx *ViewportContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type indexElement struct{}
