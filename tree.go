package main

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pop(list *lists.List) (value interface{}, ok bool) {
	value, ok = (*list).Get((*list).Size() - 1)
	(*list).Remove((*list).Size() - 1)
	return
}

func stringToTreeNode(input string) *TreeNode {
	input = strings.TrimSpace(input)
	input = input[1 : len(input)-1]
	if len(input) == 0 {
		return nil
	}
	parts := strings.Split(input, ",")
	item := parts[0]
	v, _ := strconv.Atoi(item)
	root := TreeNode{Val: v}

	var nodeQueue lists.List = arraylist.New()
	nodeQueue.Add(&root)
	index := 1
	for !nodeQueue.Empty() {
		nodeNotCasted, _ := pop(&nodeQueue)
		node := nodeNotCasted.(*TreeNode)

		if index == len(parts){
			break
		}
		item = parts[index]
		index++
		item = strings.TrimSpace(item)

		if item != "null" {
			leftNumber, _ := strconv.Atoi(item)
			node.Left = &TreeNode{Val: leftNumber}
			nodeQueue.Add(node.Left)
		}

		if index == len(parts){
			break
		}
		item = parts[index]
		index++
		item = strings.TrimSpace(item)

		if item != "null" {
			rightNumber, _ := strconv.Atoi(item)
			node.Right = &TreeNode{Val: rightNumber}
			nodeQueue.Add(node.Right)
		}
	}
	return &root
}
