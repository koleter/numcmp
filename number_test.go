package numcmp

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewNumber(t *testing.T) {
	cases := []struct {
		name   string
		num    string
		expect bool
	}{
		{
			"No.1", "--10", true,
		},
		{
			"No.2", "0", false,
		},
		{
			"No.3", "++10", true,
		},
		{
			"No.4", "10$", true,
		},
		{
			"No.5", ".", false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewNumber(tt.num)
			assert.Equal(t, tt.expect, err != nil)
		})
	}
}
