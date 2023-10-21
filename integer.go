package numcmp

import (
	"errors"
	"strings"
)

type integer struct {
	i *strings.Builder
}

func (n1 *integer) appendRune(r rune) (value, error) {
	if r == '.' {
		return &float{
			hasDot:  true,
			i:       n1.i,
			decimal: &strings.Builder{},
		}, nil
	}
	if r < '0' || r > '9' {
		return n1, errors.New("r should between '0' and '9'")
	}
	if strings.TrimLeft(n1.i.String(), "0") == "" {
		n1.i.Reset()
	}
	n1.i.WriteRune(r)
	return n1, nil
}

func (n1 *integer) compare(n2 value) int {

	switch n := n2.(type) {
	case *float:
		return n1.compareFloat(n)
	case *integer:
		return n1.compareInteger(n)
	default:
		panic("unhandle value type")
	}
}

func (n1 *integer) compareFloat(n2 *float) int {
	return -n2.compareInteger(n1)
}

func (n1 *integer) compareInteger(n2 *integer) int {
	return cmpStringOfInteger(n1.i.String(), n2.i.String())
}
