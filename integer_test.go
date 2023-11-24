package numcmp

import "testing"
import "github.com/stretchr/testify/assert"

func TestInteger_Compare_Integer(t *testing.T) {
	cases := []struct {
		name   string
		n1, n2 string
		expect int
	}{
		{
			"No.1", "10", "10", 0,
		},
		{
			"No.2", "10", "8", 1,
		},
		{
			"No.3", "10", "13", -1,
		},
		{
			"No.4", "10", "-13", 1,
		},
		{
			"No.5", "-10", "-8", -1,
		},
		{
			"No.6", "-10", "-13", 1,
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

func TestInteger_Compare_Float(t *testing.T) {
	cases := []struct {
		name   string
		n1, n2 string
		expect int
	}{
		{
			"No.1", "10", "10.0", 0,
		},
		{
			"No.2", "10", "8.3", 1,
		},
		{
			"No.3", "10", "13.7", -1,
		},
		{
			"No.4", "10", "10.2", -1,
		},
		{
			"No.5", "-10", "-8.3", -1,
		},
		{
			"No.6", "-10", "-13.7", 1,
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

func Test_integer_appendRune(t *testing.T) {
	cases := []struct {
		name     string
		n1       string
		n1Append string
		n2       string
		expect   int
	}{
		{
			"No.1", "0", "2", "10", -1,
		},
		{
			"No.2", "0", "00002", "10", -1,
		},
		{
			"No.3", "2", "00002", "10", 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			num1, _ := NewNumber(tt.n1)
			num2, _ := NewNumber(tt.n2)
			appendRunes := []rune(tt.n1Append)
			for _, appendRune := range appendRunes {
				num1.AppendRune(appendRune)
			}
			assert.Equal(t, tt.expect, num1.Cmp(num2))
		})
	}
}
