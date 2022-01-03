package main

import "fmt"

type state struct {
	r, l, p int
}

type queue struct {
	e []state
}

func (p *queue) Push(now state) {
	p.e = append(p.e, now)
}

func (p *queue) Pop() state {
	temp := p.e[0]
	p.e = p.e[1:]
	return temp

}

func (p *queue) IsEmpty() bool {
	return len(p.e) == 0
}

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	// cntr, cntl := cntTheUnmatched(s)

	return ans

}

func cntTheUnmatched(s string) (int, int) {
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return l, r
}

func main() {
	s := "()())()"
	fmt.Println(removeInvalidParentheses(s))
	testQueue := queue{e: make([]state, 0)}
	testQueue.Push(state{l: 0, r: 0, p: 12})
	fmt.Println(testQueue.IsEmpty())
	fmt.Println(testQueue.Pop())
	fmt.Println(testQueue.IsEmpty())
}
