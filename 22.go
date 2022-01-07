package main

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	implStack(n, 0, "", &ret)
	return ret
}

func implStack(out, in int, temp string, result *[]string) {
	if in == 0 && out == 0 {
		*result = append(*result, temp)
		return
	}
	// 一层只干一层的事情！写xx的for loop呢
	if out > 0 {
		implStack(out-1, in+1, temp+"(", result)
	}
	if in > 0 {
		implStack(out, in-1, temp+")", result)
	}

}

/*
func main() {
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))
}

*/
