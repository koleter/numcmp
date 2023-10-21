package numcmp

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFloat_Compare_Float(t *testing.T) {
	cases := []struct {
		name   string
		n1, n2 string
		expect int
	}{
		{
			"No.1", "10.34", "10.340", 0,
		},
		{
			"No.2", "10.15", "10.1", 1,
		},
		{
			"No.3", "10.132", "10.13", 1,
		},
		{
			"No.4", "10.132", "10.12", 1,
		},
		{
			"No.5", "-10.4", "-8.3", -1,
		},
		{
			"No.6", "-10.986", "-13.7", 1,
		},
		{
			"No.7", "-10.986", "1.7", -1,
		},
		{
			"No.8", "1.1", "3.12", -1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewNumber(tt.n1)
			num2, _ := NewNumber(tt.n2)
			assert.Equal(t, tt.expect, num1.Cmp(num2))
		})
	}
}

func Test_float_appendRune(t *testing.T) {
	cases := []struct {
		name     string
		n1       string
		n1Append rune
		n2       string
		expect   int
	}{
		{
			"No.1", "1", '2', "10", 1,
		},
		{
			"No.2", "1.2", '2', "2", -1,
		},
		{
			"No.3", "0", '2', "5", -1,
		},
		{
			"No.4", ".", '2', "1", -1,
		},
		{
			"No.5", "0.0", '2', "0.200", 0,
		},
		{
			"No.6", "0.1", '2', "0.11", 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewNumber(tt.n1)
			num2, _ := NewNumber(tt.n2)
			num1.AppendRune(tt.n1Append)
			assert.Equal(t, tt.expect, num1.Cmp(num2))
		})
	}
}
