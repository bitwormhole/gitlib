package gitfmt

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git"
)

type convertor struct {
}

func (inst *convertor) fromCommitToMessage(src *git.Commit) (*git.HyperMessage, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *convertor) fromTagToMessage(src *git.Tag) (*git.HyperMessage, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *convertor) fromMessageToCommit(src *git.HyperMessage) (*git.Commit, error) {
	builder := CommitBuilder{}
	for _, h := range src.Headers {
		err := builder.Add(h)
		if err != nil {
			return nil, err
		}
	}
	builder.SetContent(src.Content)
	dst := builder.Create()
	return dst, nil
}

func (inst *convertor) fromMessageToTag(src *git.HyperMessage) (*git.Tag, error) {

	return nil, fmt.Errorf("no impl")
}
