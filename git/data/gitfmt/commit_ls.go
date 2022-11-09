package gitfmt

import "github.com/bitwormhole/gitlib/git"

// FormatCommit ...
func FormatCommit(commit *git.Commit) (string, error) {
	cvt := convertor{}
	hm, err := cvt.fromCommitToMessage(commit)
	if err != nil {
		return "", err
	}
	return FormatHyperMessage(hm)
}

// ParseCommit ...
func ParseCommit(text string) (*git.Commit, error) {
	hm, err := ParseHyperMessage(text)
	if err != nil {
		return nil, err
	}
	cvt := convertor{}
	return cvt.fromMessageToCommit(hm)
}

////////////////////////////////////////////////////////////////////////////////

// CommitBuilder ...
type CommitBuilder struct {
	data git.Commit
}

// SetCommitter ...
func (inst *CommitBuilder) SetCommitter(value *git.HyperHeader) error {
	src := value.Values
	for _, str := range src {
		op, err := ParseOperator(str)
		if err != nil {
			return err
		}
		inst.data.Committer = op
	}
	return nil
}

// SetAuthor ...
func (inst *CommitBuilder) SetAuthor(value *git.HyperHeader) error {
	src := value.Values
	for _, str := range src {
		op, err := ParseOperator(str)
		if err != nil {
			return err
		}
		inst.data.Author = op
	}
	return nil
}

// SetTree ...
func (inst *CommitBuilder) SetTree(value *git.HyperHeader) error {
	src := value.Values
	for _, str := range src {
		oid, err := ParseObjectID(str)
		if err != nil {
			return err
		}
		inst.data.Tree = oid
	}
	return nil
}

// AddExt ...
func (inst *CommitBuilder) AddExt(src *git.HyperHeader) error {
	table := inst.data.Ext
	if table == nil {
		table = make(map[string]*git.HyperHeader)
		inst.data.Ext = table
	}
	name := src.Name
	dst := table[name]
	if dst == nil {
		dst = &git.HyperHeader{Name: name}
		table[name] = dst
	}
	dst.AddValues(src.Values)
	return nil
}

// SetParent ...
func (inst *CommitBuilder) SetParent(value *git.HyperHeader) error {
	src := value.Values
	dst := inst.data.Parents
	for _, str := range src {
		oid, err := ParseObjectID(str)
		if err != nil {
			return err
		}
		dst = append(dst, oid)
	}
	inst.data.Parents = dst
	return nil
}

// Add ...
func (inst *CommitBuilder) Add(value *git.HyperHeader) error {
	const (
		keyAuthor    = "author"
		keyCommitter = "committer"
		keyParent    = "parent"
		keyTree      = "tree"
	)
	if value == nil {
		return nil
	}
	name := value.Name
	switch name {
	case keyAuthor:
		return inst.SetAuthor(value)
	case keyCommitter:
		return inst.SetCommitter(value)
	case keyParent:
		return inst.SetParent(value)
	case keyTree:
		return inst.SetTree(value)
	default:
		break
	}
	return inst.AddExt(value)
}

// Create ...
func (inst *CommitBuilder) Create() *git.Commit {
	dst := &git.Commit{}
	*dst = inst.data
	return dst
}

// SetContent ...
func (inst *CommitBuilder) SetContent(content string) {
	inst.data.Content = content
}
