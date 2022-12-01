package algorithms

import (
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git"
)

// PathPatternMapping ...
type PathPatternMapping struct {
	pattern         string
	parts           []int // like [2,2,2,-1], 表示路径的每一节长度，最后一个小于0表示整个剩余部分
	minStringLength int   // 最小字符串长度，低于这个值的不予处理
}

func (inst *PathPatternMapping) _Impl() (git.AlgorithmRegistry, git.PathMapping) {
	return inst, inst
}

// ListRegistrations ...
func (inst *PathPatternMapping) ListRegistrations() []*git.AlgorithmRegistration {
	ar := inst.GetInfo()
	return []*git.AlgorithmRegistration{ar}
}

// GetInfo ...
func (inst *PathPatternMapping) GetInfo() *git.AlgorithmRegistration {
	return &git.AlgorithmRegistration{
		Name:     "pattern",
		Type:     git.AlgorithmPathMapping,
		Provider: inst,
	}
}

// WithPattern ...
func (inst *PathPatternMapping) WithPattern(pattern string) git.PathMapping {
	old := inst.pattern
	if old == pattern {
		return inst
	}
	next := &PathPatternMapping{
		pattern: pattern,
	}
	return next
}

func (inst *PathPatternMapping) loadparts() ([]int, int) {
	total := 0
	src := strings.Split(inst.pattern, "/")
	end := len(src) - 1
	dst := make([]int, 0)
	for i, part := range src {
		part = strings.TrimSpace(part)
		size := len(part)
		if size == 0 {
			continue
		}
		if i < end {
			dst = append(dst, size)
			total += size
		} else {
			dst = append(dst, -1)
			break
		}
	}
	minStrLen := total + 1
	return dst, minStrLen
}

func (inst *PathPatternMapping) getparts() ([]int, int) {
	parts := inst.parts
	msl := inst.minStringLength
	if parts == nil {
		parts, msl = inst.loadparts()
		inst.minStringLength = msl
		inst.parts = parts
	}
	return parts, msl
}

// Map ...
func (inst *PathPatternMapping) Map(base afs.Path, id git.ObjectID) afs.Path {

	parts, minStrLen := inst.getparts()
	str := id.String()
	const padding = 5
	if len(str) < minStrLen+padding {
		return base.GetChild(str)
	}

	builder := strings.Builder{}
	sep := ""
	i1 := 0
	for _, pLen := range parts {
		if pLen < 0 {
			part := str[i1:]
			builder.WriteString(sep)
			builder.WriteString(part)
			break
		} else if pLen == 0 {
			continue
		}
		i2 := i1 + pLen
		part := str[i1:i2]
		builder.WriteString(sep)
		builder.WriteString(part)
		i1 = i2
		sep = "/"
	}
	return base.GetChild(builder.String())
}
