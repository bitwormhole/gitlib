package gitfmt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitwormhole/gitlib/git"
	"github.com/starter-go/base/lang"
)

// FormatOperator ...
func FormatOperator(o *git.Operator) (string, error) {

	return "", fmt.Errorf("no impl")
}

// ParseOperator ...
func ParseOperator(text string) (*git.Operator, error) {

	text = strings.TrimSpace(text)
	i1 := strings.LastIndexByte(text, '<')
	i2 := strings.LastIndexByte(text, '>')
	i3 := strings.LastIndexByte(text, ' ')

	if 0 <= i1 && i1 < i2 && i2 < i3 {
		// continue
	} else {
		return nil, fmt.Errorf("bad git.Operator string: " + text)
	}

	p1 := strings.TrimSpace(text[0:i1])
	p2 := strings.TrimSpace(text[i1+1 : i2])
	p3 := strings.TrimSpace(text[i2+1 : i3])
	p4 := strings.TrimSpace(text[i3+1:])

	timeNum1, err := strconv.ParseInt(p3, 10, 0)
	if err != nil {
		return nil, err
	}
	tt := lang.Time(timeNum1 * 1000).Time()

	result := &git.Operator{
		Name:  p1,
		Email: p2,
		Time:  tt,
		Zone:  p4,
	}
	return result, nil
}
