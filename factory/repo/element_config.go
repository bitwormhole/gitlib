package repo

type GitConfigElementFactory struct{}

func (inst *GitConfigElementFactory) _Impl() ElementFactory {
	return inst
}

func (inst *GitConfigElementFactory) Make(ctx *ViewportContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type configElement struct{}
