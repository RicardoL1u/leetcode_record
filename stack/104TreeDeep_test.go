package stack

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type stackTree struct {
	data []*TreeNode
}

func (s *stackTree) Push(node *TreeNode) {
	s.data = append(s.data, node)
}
func (s *stackTree) Pop() *TreeNode {
	ans := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ans
}
func (s *stackTree) IsEmpty() bool {
	return len(s.data) == 0
}

func maxDepth(root *TreeNode) int {
	ans := 0
	nowDeep := 0
	s := stackTree{data: make([]*TreeNode, 0)}
	for root != nil || !s.IsEmpty() {
		for root != nil {
			nowDeep++
			ans = max(ans, nowDeep)
			s.Push(root)
			root = root.Left
		}
		root = s.data[len(s.data)-1]
		root = root.Right
		for root != nil {
			nowDeep++
			ans = max(ans, nowDeep)
			s.Push(root)
			root = root.Left
		}
		s.Pop()
		nowDeep--
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
