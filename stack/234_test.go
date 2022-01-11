package stack

import (
	"log"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// type Stack struct {
// 	data []int
// }

func isPalindrome(head *ListNode) bool {
	p := head
	listLen := 1
	for p.Next != nil {
		p = p.Next
		listLen++
	}
	end := p
	if listLen == 1 {
		return true
	}
	if listLen == 2 {
		return head.Val == head.Next.Val
	}
	half := listLen / 2
	cnt := 0
	p = head
	prev := head
	for {
		tmp := p.Next
		if cnt > half {
			p.Next = prev
		}
		prev = p
		p = tmp
		cnt++
		if p == nil {
			break
		}
	}
	p1 := head
	p2 := end
	for p1 != p2 {
		// log.Println("hi")
		if p1.Val != p2.Val {
			return false
		}
		if p1.Next == p2 {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
		log.Printf("%v %v\n", p1, p2)
	}
	return true
}

func TestISPalindrome(t *testing.T) {
	node4 := ListNode{
		Val: 1,
	}
	node3 := ListNode{
		Val:  2,
		Next: &node4,
	}
	node2 := ListNode{
		Val:  1,
		Next: &node3,
	}
	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}
	log.Println(isPalindrome(&node1))
}
