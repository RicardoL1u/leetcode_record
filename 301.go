package main

type state struct {
	rml, rmr, p int
	ul, ur      int
	tol         int
	ans         string
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
	strSet := make(map[string]bool)
	cntl, cntr, ul, ur := cntTheUnmatched(s)
	q := queue{e: make([]state, 0)}
	q.Push(state{rmr: cntr, rml: cntl, p: 0, tol: 0, ul: ul, ur: ur, ans: ""})
	for !q.IsEmpty() {
		now := q.Pop()
		if (now.rml == now.rmr && now.rml == 0) || now.p == len(s) {
			temp := now.ans
			if now.p < len(s) {
				temp = temp + s[now.p:]
			}
			tr, tl, _, _ := cntTheUnmatched(temp)
			if tr == 0 && tl == 0 {
				strSet[temp] = true
			}
			continue
		}
		// next := now
		switch s[now.p] {
		case '(':
			if now.rml > 0 {
				next := now
				next.rml--
				next.p++
				q.Push(next)
			}
			if now.ur > 0 { // 只当后面存在还未被匹配的）才入队左括号(
				next := now
				next.p++
				next.tol++
				next.ans += "("
				q.Push(next)
			}
		case ')':
			if now.rmr > 0 {
				next := now
				next.rmr--
				next.p++
				q.Push(next)
			}
			if now.tol > 0 { // 只当前面还存在未被匹配的(才入队右括号
				next := now
				next.p++
				next.tol--
				next.ur--
				next.ans += ")"
				q.Push(next)
			}
		default:
			next := now
			next.p++
			next.ans += string(s[now.p])
			q.Push(next)
		}

	}
	for k, _ := range strSet {
		ans = append(ans, k)
	}
	return ans
}

func cntTheUnmatched(s string) (int, int, int, int) {
	totl, totr := 0, 0
	rml, rmr := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			totl++
			rml++
		} else if s[i] == ')' {
			totr++
			if rml > 0 {
				rml--
			} else {
				rmr++
			}
		}
	}
	return rml, rmr, totl - rml, totr - rmr
}

/*
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

*/
