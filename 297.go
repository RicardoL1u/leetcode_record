package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := ""
	if root.Val != -2000 {
		ans += (fmt.Sprint(root.Val))
		if root.Left != nil {
			ans += ("," + this.serialize(root.Left))
		}
		if root.Right != nil {
			ans += ("," + this.serialize(root.Right))
		}
		return ans
	} else {
		return "null"
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	ptr := 0
	return this.formTree(nodes, &ptr)
}

func (this *Codec) formTree(node []string, ptr *int) *TreeNode {
	if len(node) == *ptr {
		return nil
	}
	if node[*ptr] == "null" {
		(*ptr)++
		return &TreeNode{Val: -2000}
	} else {
		val, _ := strconv.Atoi(node[*ptr])
		(*ptr)++
		return &TreeNode{
			Val:   val,
			Left:  this.formTree(node, ptr),
			Right: this.formTree(node, ptr),
		}
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	testStr := "1,2,3,null,null,4,5"
	ser := Constructor()
	deser := Constructor()
	// fmt.Println(deser.deserialize(testStr))
	root := deser.deserialize(testStr)

	testRoot := TreeNode{Val: 1}
	testRoot.Left = &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: -2000},
		Right: &TreeNode{Val: -2000},
	}
	testRoot.Right = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println(testRoot.Left)
	fmt.Println(ser.serialize(&testRoot))
	fmt.Println(ser.serialize(root))
}
