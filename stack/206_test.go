package stack

import (
	"log"
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	p := head.Next
	head.Next = nil
	for p.Next != nil {
		realNext := p.Next
		p.Next = prev
		prev = p
		p = realNext
	}
	p.Next = prev
	return p
}

func printList(head *ListNode) {
	p := head
	cnt := 0
	for p != nil {
		log.Println(p)
		p = p.Next
		cnt++
		if cnt > 10 {
			break
		}
	}
}
func TestReverseList(t *testing.T) {
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
	printList(&node1)
	log.Println()
	printList(reverseList(&node1))

}
