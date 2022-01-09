package main

import (
	"fmt"

	"sort"
)

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
	//QuickSort(&data, 0, len(data))
	fmt.Printf("%d", data[0])
	for i := 1; i < num; i++ {
		fmt.Printf(" %d", data[i])
	}
	// return
}
