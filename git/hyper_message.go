package git

import "strings"

// HyperHeader ...
type HyperHeader struct {
	Name   string
	Values []string
}

// AddValues ....
func (inst *HyperHeader) AddValues(src []string) {
	for _, v := range src {
		inst.AddValue(v)
	}
}

// AddValue ....
func (inst *HyperHeader) AddValue(v string) {
	v = strings.TrimSpace(v)
	if v == "" {
		return
	}
	inst.Values = append(inst.Values, v)
}

////////////////////////////////////////////////////////////////////////////////

// HyperMessage ...
type HyperMessage struct {
	Headers map[string]*HyperHeader
	Content string
}
