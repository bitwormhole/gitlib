package repo

type GitHeadElementFactory struct{}

func (inst *GitHeadElementFactory) _Impl() ElementFactory {
	return inst
}

func (inst *GitHeadElementFactory) Make(ctx *ViewportContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type headElement struct{}
