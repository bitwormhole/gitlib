package refs

import (
	"sort"
	"strings"

	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/afs"
)

type finderNode struct {
	parent *finderNode
	root   afs.Path // the 'refs' dir
	path   afs.Path // this node
	name   string
}

func (inst *finderNode) init(refs afs.Path) *finderNode {
	inst.parent = nil
	inst.root = refs
	inst.path = refs
	inst.name = refs.GetName()
	return inst
}

func (inst *finderNode) items() []string {
	return inst.path.ListNames()
}

func (inst *finderNode) isFile() bool {
	return inst.path.IsFile()
}

func (inst *finderNode) isDir() bool {
	return inst.path.IsDirectory()
}

func (inst *finderNode) getChild(name string) *finderNode {
	child := inst.path.GetChild(name)
	return &finderNode{
		parent: inst,
		root:   inst.root,
		path:   child,
		name:   name,
	}
}

////////////////////////////////////////////////////////////////////////////////

type finder struct {
	results []string
}

func (inst *finder) find(refs afs.Path) []git.ReferenceName {
	root := &finderNode{}
	root = root.init(refs)
	inst.walk(root, 9)

	dst := make([]git.ReferenceName, 0)
	src := inst.results
	sort.Strings(src)

	for _, str := range src {
		name := git.ReferenceName(str)
		dst = append(dst, name)
	}

	return dst
}

func (inst *finder) walk(node *finderNode, depthLimit int) {

	if depthLimit < 1 {
		return
	}

	if node.isDir() {
		items := node.items()
		for _, name := range items {
			child := node.getChild(name)
			inst.walk(child, depthLimit-1)
		}
	} else if node.isFile() {
		inst.handleFile(node)
	}
}

func (inst *finder) handleFile(node *finderNode) {
	builder := strings.Builder{}
	inst.buildRefName(node, &builder)
	name := builder.String()
	inst.results = append(inst.results, name)
}

func (inst *finder) buildRefName(p *finderNode, builder *strings.Builder) {
	if p == nil {
		return
	}
	parent := p.parent
	inst.buildRefName(parent, builder)
	if parent != nil {
		builder.WriteString("/")
	}
	builder.WriteString(p.name)
}
