package bit

import (
	"log"
	"testing"
)

// 异或的性质 a xor b xor a  = b
func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}

func bitSwap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func TestBisSwap(t *testing.T) {
	a, b := 1, 2
	bitSwap(&a, &b)
	log.Println(a, b)
}
