package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	codec := Constructor()
	s := codec.serialize(root)
	fmt.Println(s)
	node := codec.deserialize(s)
	fmt.Println(node)
}

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
	var sb = &strings.Builder{}
	var build func(node *TreeNode)
	build = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("nil,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		build(node.Left)
		build(node.Right)
	}
	build(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var arr = strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if arr[0] == "nil" {
			arr = arr[1:]
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		arr = arr[1:]
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
