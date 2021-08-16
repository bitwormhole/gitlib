package repo

type GitRefsElementFactory struct{}

func (inst *GitRefsElementFactory) _Impl() ElementFactory {
	return inst
}

func (inst *GitRefsElementFactory) Make(ctx *ViewportContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type refsElement struct{}
