package gitfmt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/gitlib/git"
)

// FormatHyperMessage ...
func FormatHyperMessage(msg *git.HyperMessage) (string, error) {
	const nl = '\n'
	builder := strings.Builder{}
	hds := msg.Headers
	keys := make([]string, 0)
	for key := range hds {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		hh := hds[key]
		values := hh.Values
		for _, v := range values {
			builder.WriteString(hh.Name)
			builder.WriteRune(' ')
			builder.WriteString(v)
			builder.WriteRune(nl)
		}
	}
	builder.WriteRune(nl)
	builder.WriteString(msg.Content)
	return builder.String(), nil
}

// ParseHyperMessage ...
func ParseHyperMessage(text string) (*git.HyperMessage, error) {
	const (
		markEndOfHead = "\n\n"
		headerSep     = "\n"
	)
	// head & body
	parts := strings.SplitN(text, markEndOfHead, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("parts.length != 2")
	}
	head := parts[0]
	body := parts[1]
	// parse head
	parts = strings.Split(head, headerSep)
	builder := hyperMessageHeadBuilder{}
	for _, row := range parts {
		err := builder.addRow(row)
		if err != nil {
			return nil, err
		}
	}
	// result
	headers := builder.create()
	msg := &git.HyperMessage{
		Headers: headers,
		Content: body,
	}
	return msg, nil
}

////////////////////////////////////////////////////////////////////////////////

type hyperMessageHeadBuilder struct {
	table map[string]*git.HyperHeader
}

func (inst *hyperMessageHeadBuilder) getTab() map[string]*git.HyperHeader {
	t := inst.table
	if t != nil {
		return t
	}
	t = make(map[string]*git.HyperHeader)
	inst.table = t
	return t
}

func (inst *hyperMessageHeadBuilder) create() map[string]*git.HyperHeader {
	return inst.getTab()
}

func (inst *hyperMessageHeadBuilder) addRow(row string) error {
	t := inst.getTab()
	header, err := inst.parseRow(row)
	if err != nil {
		return err
	}
	key := header.Name
	old := t[key]
	if old == nil {
		t[key] = header
	} else {
		old.Values = append(old.Values, header.Values...)
	}
	return nil
}

func (inst *hyperMessageHeadBuilder) parseRow(row string) (*git.HyperHeader, error) {
	const (
		n     = 2
		sep   = " "
		empty = ""
	)
	parts := strings.SplitN(row, sep, n)
	if len(parts) != n {
		return nil, fmt.Errorf("parts.length != %v", n)
	}
	k := strings.TrimSpace(parts[0])
	v := strings.TrimSpace(parts[1])
	if k == empty || v == empty {
		return nil, fmt.Errorf("bad hyper-header row: " + row)
	}
	hh := &git.HyperHeader{
		Name:   k,
		Values: []string{v},
	}
	return hh, nil
}
