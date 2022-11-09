package gitfmt

import "github.com/bitwormhole/gitlib/git"

// FormatTag ...
func FormatTag(tag *git.Tag) (string, error) {
	cvt := convertor{}
	hm, err := cvt.fromTagToMessage(tag)
	if err != nil {
		return "", err
	}
	return FormatHyperMessage(hm)
}

// ParseTag ...
func ParseTag(text string) (*git.Tag, error) {
	hm, err := ParseHyperMessage(text)
	if err != nil {
		return nil, err
	}
	cvt := convertor{}
	return cvt.fromMessageToTag(hm)
}
