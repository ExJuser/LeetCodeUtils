package linkedlist

import (
	"github.com/alecthomas/assert/v2"
	"reflect"
	"testing"
)

func TestLinkedListEqual(t *testing.T) {
	testCase := []struct {
		name     string
		input    [2]*ListNode
		expected bool
	}{
		{
			name: "相等的链表",
			input: [2]*ListNode{
				GenerateLinkedListFromSlice([]int{1, 2, 3, 4}),
				GenerateLinkedListFromSlice([]int{1, 2, 3, 4}),
			},
			expected: true,
		},
		{
			name: "相等的空链表",
			input: [2]*ListNode{
				GenerateLinkedListFromSlice([]int{}),
				GenerateLinkedListFromSlice([]int{}),
			},
			expected: true,
		},
		{
			name: "不等的链表",
			input: [2]*ListNode{
				GenerateLinkedListFromSlice([]int{1, 2, 3, 4}),
				GenerateLinkedListFromSlice([]int{1, 2, 3}),
			},
			expected: false,
		},
		{
			name: "不等的链表，有一个为空",
			input: [2]*ListNode{
				GenerateLinkedListFromSlice([]int{1, 2, 3, 4}),
				GenerateLinkedListFromSlice([]int{}),
			},
			expected: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res1 := CheckLinkedListEqual(tc.input[0], tc.input[1])
			res2 := reflect.DeepEqual(tc.input[0], tc.input[1])
			assert.Equal(t, tc.expected, res1)
			assert.Equal(t, res1, res2)
		})
	}
}
