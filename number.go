package numcmp

import (
	"errors"
	"regexp"
	"strings"
)

type value interface {
	compare(n value) int
}

type Number struct {
	neg     int
	numType int
	value
}

func (n *Number) GetNumType() int {
	return n.numType
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

func (n1 *Number) Cmp(n2 *Number) int {
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

func (n1 *Number) cmpSign(n2 *Number) int {
	if n1.neg > n2.neg {
		return 1
	} else if n1.neg < n2.neg {
		return -1
	}
	return 0
}

const (
	IntNumber = 1 << iota
	FloatNumber
	InvalidNumber

	// 所有有效的数字类型
	AllTypeNumber = IntNumber | FloatNumber
)

func NewNumber(str string) (*Number, error) {
	neg := 1
	if str[0] == '-' {
		neg = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	numType := judgeNumber(str)

	switch numType {
	case IntNumber:
		str = strings.TrimLeft(str, "0")
		if str == "" {
			str = "0"
		}
		return &Number{neg, numType, &integer{str}}, nil
	case FloatNumber:
		str = strings.Trim(str, "0")
		split := strings.Split(str, ".")
		if split[0] == "" {
			split[0] = "0"
		}
		return &Number{neg, numType, &float{i: split[0], decimal: split[1]}}, nil
	default:
		return nil, errors.New("unexpect input str: " + str)
	}
}

func judgeNumber(str string) int {
	compile := regexp.MustCompile("^\\d+$")
	if compile.MatchString(str) {
		return IntNumber
	}
	compile = regexp.MustCompile("^\\d*\\.\\d*$")
	if compile.MatchString(str) {
		return FloatNumber
	}
	return InvalidNumber
}
