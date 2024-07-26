package slice

import (
	"github.com/ExJuser/LeetCodeUtils/common"
	"math/rand/v2"
)

func GenerateIntSlice(bound, length int) (res []int) {
	for i := 0; i < length; i++ {
		res = append(res, rand.IntN(bound))
	}
	return
}

func SlicesSum[T common.Number](nums []T) (sum T) {
	for _, num := range nums {
		sum += num
	}
	return
}

func PrefixSum(nums []int) []int {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return prefix
}
