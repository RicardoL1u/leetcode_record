package stack

import (
	"log"
	"testing"
)

type stack struct {
	data []*TreeNode
	p    int
}

func (s *stack) Pop() *TreeNode {
	s.p--
	return s.data[s.p]
}

func (s *stack) Push(e *TreeNode) {
	if s.p == len(s.data) {
		s.data = append(s.data, e)
	} else {
		s.data[s.p] = e
	}
	s.p++
}
func (s *stack) isEmpty() bool {
	return s.p == 0
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	s := stack{
		data: make([]*TreeNode, 0),
		p:    0,
	}
	ans := make([]int, 0)
	now := root
	for now != nil || !s.isEmpty() {
		for now != nil { // 模仿深搜的不断深入
			s.Push(now)    // 确实由于我们在深搜的时候也从来没吧 左右点区别对待
			now = now.Left // 所以我们都是push now
		}
		now = s.Pop()
		ans = append(ans, now.Val)
		now = now.Right
	}
	return ans
}

func TestInorder(t *testing.T) {
	node2 := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: nil,
	}
	node1 := TreeNode{
		Val:   2,
		Left:  &node2,
		Right: nil,
	}
	log.Println(inorderTraversal(&node1))
}
