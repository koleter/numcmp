package numcmp

import (
	"errors"
	"strings"
)

type float struct {
	hasDot bool
	// part of integer
	i       *strings.Builder
	decimal *strings.Builder
}

func (n1 *float) appendRune(r rune) (value, error) {
	if r == '.' {
		if n1.hasDot {
			return n1, errors.New("invalid byte r")
		}
		n1.hasDot = true
		return n1, nil
	}
	if r < '0' || r > '9' {
		return n1, errors.New("invalid byte r")
	}
	if n1.hasDot {
		n1.decimal.WriteRune(r)
	} else {
		if strings.TrimLeft(n1.i.String(), "0") == "" {
			n1.i.Reset()
		}
		n1.i.WriteRune(r)
	}
	return n1, nil
}

func (n1 *float) compare(n2 value) int {
	switch n := n2.(type) {
	case *float:
		return n1.compareFloat(n)
	case *integer:
		return n1.compareInteger(n)
	default:
		panic("unhandle value type")
	}
}

func (n1 *float) compareFloat(n2 *float) int {
	resulti := cmpStringOfInteger(n1.i.String(), n2.i.String())
	if resulti == 1 || resulti == 0 && n1.decimal.String() > n2.decimal.String() {
		return 1
	} else if resulti == 0 && n1.decimal.String() == n2.decimal.String() {
		return 0
	}
	return -1
}

func (n1 *float) compareInteger(n2 *integer) int {
	resulti := cmpStringOfInteger(n1.i.String(), n2.i.String())
	if resulti == 1 || resulti == 0 && n1.decimal.String() != "" {
		return 1
	} else if resulti == 0 && n1.decimal.String() == "" {
		return 0
	}
	return -1
}
