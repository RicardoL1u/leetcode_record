package binarysearch

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
)

func majorityElement(nums []int) int {
	suffledData := make([]int, len(nums))
	perm := rand.Perm(len(nums)) // if some one say quick is not stable
	// show this and fuck off
	for i, v := range perm {
		suffledData[v] = nums[i]
	}
	p := findK(&suffledData, 0, len(nums), len(nums)/2)
	return p
}

func findK(nums *[]int, s, e, k int) int { // 我这里的k就是从大到小排序后的下标
	i, j := s, e-1
	p := (*nums)[s]
	for i != j {
		for i < j && (*nums)[j] >= p {
			j--
		}
		(*nums)[i] = (*nums)[j]
		for i < j && (*nums)[i] < p {
			i++
		}
		(*nums)[j] = (*nums)[i]
	}
	(*nums)[i] = p
	if k == i {
		return p
	} else if k > i {
		return findK(nums, i+1, e, k)
	} else {
		return findK(nums, s, i, k)
	}
}

func TestMajority(t *testing.T) {
}

func QuickSort(nums *[]int, s, e int) {
	if e-s <= 1 { // 由于是左闭右开惯例，e-s<=1返回
		return
	}
	p := (*nums)[s] // 记录枢纽
	i, j := s, e-1  // i 从头 j 从尾巴
	for i != j {
		for i < j && (*nums)[j] >= p {
			j--
		}
		(*nums)[i] = (*nums)[j]
		for i < j && (*nums)[i] < p {
			i++
		}
		(*nums)[j] = (*nums)[i]
	}
	(*nums)[i] = p
	// log.Printf("now i = %d,j = %d, nums = %v\n", i, j, *nums)
	QuickSort(nums, s, i)
	QuickSort(nums, i+1, e)
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	num := 0
	fmt.Scanf("%d\n", &num)
	data := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &data[i])
	}
	sort.Sort(SortBy(data))
	QuickSort(&data, 0, len(data))
	fmt.Printf("%d", data[0])
	for i := 1; i < num; i++ {
		fmt.Print(" %d", data[i])
	}
	// return
}

func TestQuickSort(t *testing.T) {
	data := []int{3, 2, 1, 1, 2, 3, 4, 5}
	QuickSort(&data, 0, len(data))
	log.Println(data)
}
