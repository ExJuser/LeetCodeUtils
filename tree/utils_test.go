package tree

import (
	"fmt"
	"github.com/alecthomas/assert/v2"
	"math/rand/v2"
	"testing"
)

func TestGenerateTreeFromLevelOrder(t *testing.T) {
	for i := 0; i < 1000; i++ {
		randomOrders := make([]int, 1000)
		for j := 0; j < len(randomOrders); j++ {
			randomOrders[j] = rand.IntN(11) - 1
		}
		assert.NotPanics(t, func() {
			GenerateTreeFromLevelOrder(randomOrders)
		})
	}
}

func TestPrintAllPathsToLeaves(t *testing.T) {
	randomOrders := make([]int, 10)
	for j := 0; j < len(randomOrders); j++ {
		randomOrders[j] = rand.IntN(11) - 1
	}
	root := GenerateTreeFromLevelOrder(randomOrders)
	fmt.Println(randomOrders)
	assert.NotPanics(t, func() {
		PrintAllPathsToLeaves(root)
	})
}

func TestLevelOrder(t *testing.T) {
	randomOrders := make([]int, 15)
	for j := 0; j < len(randomOrders); j++ {
		randomOrders[j] = rand.IntN(11) - 1
	}
	root := GenerateTreeFromLevelOrder(randomOrders)
	fmt.Println(LevelOrder(root))
	PrintAllPathsToLeaves(root)
}
