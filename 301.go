package main

import (
	"fmt"
)

type state struct {
	l, r, p, unmatchedL int
	ans                 string
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

func searchAns(s string, ch chan []string) {
	ans := make([]string, 0)
	strSet := make(map[string]bool)
	cntl, cntr := cntTheUnmatched(s)
	q := queue{e: make([]state, 0)}
	q.Push(state{r: cntr, l: cntl, p: 0, unmatchedL: 0, ans: ""})
	for !q.IsEmpty() {
		now := q.Pop()
		if (now.l == now.r && now.l == 0) || now.p == len(s) {
			// fmt.Println(now)
			temp := now.ans
			if now.p < len(s) {
				temp = temp + s[now.p:]
			}
			tr, tl := cntTheUnmatched(temp)
			if tr == 0 && tl == 0 {
				strSet[temp] = true
			}
			continue
		}
		// next := now
		switch s[now.p] {
		case '(':
			if now.l > 0 {
				q.Push(state{r: now.r, l: now.l - 1, p: now.p + 1, ans: now.ans})
			}
			q.Push(state{r: now.r, l: now.l, p: now.p + 1, unmatchedL: now.unmatchedL + 1, ans: now.ans + "("})
		case ')':
			if now.r > 0 {
				q.Push(state{r: now.r - 1, l: now.l, p: now.p + 1, ans: now.ans})
			}
			if now.unmatchedL > 0 {
				q.Push(state{r: now.r, l: now.l, p: now.p + 1, unmatchedL: now.unmatchedL - 1, ans: now.ans + ")"})
			}
			if now.r == 0 && now.unmatchedL == 0 {
				continue
			}
		default:
			q.Push(state{r: now.r, l: now.l, p: now.p + 1, unmatchedL: now.unmatchedL, ans: now.ans + string(s[now.p])})
		}

	}
	for k, _ := range strSet {
		ans = append(ans, k)
	}
	ch <- ans
}

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	ch1 := make(chan []string)
	ch2 := make(chan []string)
	go searchAns(s, ch1)
	go searchAns(reversePar(s), ch2)
	select {
	case ans = <-ch1:
		return ans
	case ans = <-ch2:
		return ans
	}
	return ans
}

func reversePar(s string) string {
	ans := ""
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			ans = ans + ")"
		case ')':
			ans = ans + "("
		default:
			ans = ans + string(s[i])
		}
	}
	return ans
}

func cntTheUnmatched(s string) (int, int) {
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
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
	// s := "(a)())()"
	// s := ")("
	// fmt.Println(cntTheUnmatched(s))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")s(asd)"))
	fmt.Println(removeInvalidParentheses("a"))
	fmt.Println(removeInvalidParentheses("("))
	fmt.Println(removeInvalidParentheses("(a)())((((()"))

}
