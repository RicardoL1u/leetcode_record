package stack

import (
	"log"
	"strconv"
	"testing"
)

type stackStr struct {
	data []string
}

func (s *stackStr) Push(e string) {
	s.data = append(s.data, e)
}
func (s *stackStr) Pop() string {
	ans := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ans
}

func calculate(op string, a, b int) string {
	ans := 0
	switch op {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	}
	log.Printf("The ans is %d\n", ans)
	return strconv.Itoa(ans)
}

func evalRPN(tokens []string) int {
	// op := stack{data:make([]string,0)}
	opr := stackStr{data: make([]string, 0)}
	// opr2 := stack{data:make([]rune,0)}
	for _, token := range tokens {
		if IsDigital(token) {
			opr.Push(token)
		} else {
			op2, _ := strconv.Atoi(opr.Pop())
			op1, _ := strconv.Atoi(opr.Pop())
			opr.Push(calculate(token, op1, op2))
		}
		log.Println(opr.data)
	}
	ans, _ := strconv.Atoi(opr.Pop())
	return ans
}

func IsDigital(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func TestEval(t *testing.T) {
	// test :=
	log.Println(evalRPN([]string{"4", "13", "5", "/", "+"}) == 6)
	log.Println(evalRPN([]string{"2", "1", "+", "3", "*"}) == 9)
	log.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}) == 22)
}
