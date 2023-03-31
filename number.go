package numcmp

import (
	"errors"
	"regexp"
	"strings"
)

type value interface {
	compare(n value) int
}

type number struct {
	neg int
	value
}

func cmpStringOfInteger(n1, n2 string) int {
	l1 := len(n1)
	l2 := len(n2)
	if l1 > l2 {
		return 1
	} else if l1 < l2 {
		return -1
	}
	if n1 > n2 {
		return 1
	} else if n1 < n2 {
		return -1
	}
	return 0
}

func (n1 *number) Cmp(n2 *number) int {
	sign := n1.cmpSign(n2)
	if sign != 0 {
		return sign
	}
	compare := n1.value.compare(n2.value)
	if n1.neg == -1 {
		compare = -compare
	}
	return compare
}

func (n1 *number) cmpSign(n2 *number) int {
	if n1.neg > n2.neg {
		return 1
	} else if n1.neg < n2.neg {
		return -1
	}
	return 0
}

const (
	intNumber = iota
	floatNumber
	invalidNumber
)

func NewNumber(str string) (*number, error) {
	neg := 1
	if str[0] == '-' {
		neg = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	numType := judgeNumber(str)

	switch numType {
	case intNumber:
		str = strings.TrimLeft(str, "0")
		if str == "" {
			str = "0"
		}
		return &number{neg, &integer{str}}, nil
	case floatNumber:
		str = strings.Trim(str, "0")
		split := strings.Split(str, ".")
		if split[0] == "" {
			split[0] = "0"
		}
		return &number{neg, &float{i: split[0], decimal: split[1]}}, nil
	default:
		return nil, errors.New("unexpect input str: " + str)
	}
}

func judgeNumber(str string) int {
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	compile := regexp.MustCompile("^\\d+$")
	if compile.MatchString(str) {
		return intNumber
	}
	compile = regexp.MustCompile("^\\d*\\.\\d*$")
	if compile.MatchString(str) {
		return floatNumber
	}
	return invalidNumber
}
