package main

import (
	"fmt"
	"sort"
)

type sortByte []byte

func (temp sortByte) Len() int           { return len(temp) }
func (temp sortByte) Swap(i, j int)      { temp[i], temp[j] = temp[j], temp[i] }
func (temp sortByte) Less(i, j int) bool { return temp[i] < temp[j] }
func groupAnagrams(strs []string) [][]string {
	// rec := make([][]int, len(strs))
	ans := make([][]string, 0)

	rec := make(map[string]int)
	cnt := 0
	for _, item := range strs {
		temp := sortByte(item)
		sort.Sort(temp)
		if k, ok := rec[string(temp)]; ok {
			ans[k] = append(ans[k], item)
		} else {
			rec[string(temp)] = cnt
			ans = append(ans, make([]string, 0))
			ans[cnt] = append(ans[cnt], item)
			cnt++
		}
	}
	return ans
}
func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}
