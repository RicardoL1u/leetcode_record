package main

func findRepeatedDnaSequences(s string) (ans []string) {
	bitmap := map[byte]uint32{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}
	revmap := map[uint32]byte{
		0: 'A', 1: 'C', 2: 'G', 3: 'T',
	}
	repMap := make(map[uint32]uint8)
	var temp uint32
	temp = 0
	var cmpNum uint32 = 1<<20 - 1
	for i := 0; i < len(s); i++ {
		temp = temp<<2 + bitmap[s[i]]
		temp = temp & cmpNum
		if i >= 9 {
			// fmt.Printf("%x\n", temp)
			if repMap[temp] == 0 {
				repMap[temp] = 1
			} else {
				repMap[temp]++
			}
		}
	}

	for k, v := range repMap {
		// fmt.Println(v)
		if v < 2 {
			continue
		}
		retStr := make([]byte, 10)
		for i := 0; i < 10; i++ {
			var revint uint32 = k & 3
			retStr[9-i] = revmap[revint]
			k = k >> 2
		}
		ans = append(ans, string(retStr))
	}
	return
}

/*
func main() {
	test := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	// test := "AAAAAAAAAAA"
	fmt.Printf("%x\n", 1<<10-1)
	fmt.Print(findRepeatedDnaSequences(test))

}

*/
