package slice

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestPrefixSum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{0, 1, 3, 6, 10, 15},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, PrefixSum(tc.input))
		})
	}
}
