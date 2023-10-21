package numcmp

import (
	"errors"
	"regexp"
	"strings"
)

type value interface {
	compare(n value) int
	appendRune(r rune) (value, error)
}

type Number struct {
	neg     int
	numType NumberType
	value
}

func (n *Number) GetNumType() NumberType {
	return n.numType
}

// AppendRune is a function which equivalent to adding one digit to the end of the current number
func (n *Number) AppendRune(r rune) (*Number, error) {
	v, err := n.appendRune(r)
	if err != nil {
		return n, err
	}
	return &Number{
		neg:     n.neg,
		numType: getValueNumType(v),
		value:   v,
	}, nil
}

func getValueNumType(v value) NumberType {
	switch v.(type) {
	case *integer:
		return IntNumber
	case *float:
		return FloatNumber
	}
	return InvalidNumber
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

type NumberType int

const (
	IntNumber = NumberType(1) << iota
	FloatNumber
	InvalidNumber

	// 所有有效的数字类型
	AllTypeNumber = IntNumber | FloatNumber
)

func NewNumber(str string) (*Number, error) {
	if len(str) == 0 {
		return nil, errors.New("input can not be null")
	}
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
		sb := &strings.Builder{}
		sb.WriteString(str)
		return &Number{neg, numType, &integer{sb}}, nil
	case FloatNumber:
		str = strings.Trim(str, "0")
		split := strings.Split(str, ".")
		hasDot := strings.Contains(str, ".")
		if split[0] == "" {
			split[0] = "0"
		}
		sb0 := &strings.Builder{}
		sb0.WriteString(split[0])
		sb1 := &strings.Builder{}
		sb1.WriteString(split[1])
		return &Number{neg, numType, &float{i: sb0, decimal: sb1, hasDot: hasDot}}, nil
	default:
		return nil, errors.New("unexpect input str: " + str)
	}
}

func judgeNumber(str string) NumberType {
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
