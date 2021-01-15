package algor_lib

import (
	"errors"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(arr []int) (tree *TreeNode, err error) {
	arrLen := len(arr)
	if arrLen < 1 {
		return nil, errors.New("empty arr")
	}
	maxDep := int(math.Ceil(math.Log2(float64(arrLen - 1))))

	treeNodeMap := make(map[int]*TreeNode)
	for dep := 0; dep <= maxDep; dep++ {
		for i := int(math.Pow(float64(2), float64(dep))) - 1; i < int(math.Pow(float64(2), float64(dep+1)))-1; i++ {
			if i >= arrLen {
				break
			}

			node := TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			treeNodeMap[i] = &node

			if i == 0 {
				continue
			}

			parentIdx := 0
			if i%2 == 1 {
				parentIdx = int((i - 1) / 2)
				treeNodeMap[parentIdx].Left = &node
			} else {
				parentIdx = int((i - 2) / 2)
				treeNodeMap[parentIdx].Right = &node
			}
		}
	}

	return treeNodeMap[0], nil
}

func Dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	Dfs(root.Left, result)
	Dfs(root.Right, result)
}

func Bfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	stack := make([]TreeNode, 0, 10)
	stack = append(stack, *root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*result = append(*result, node.Val)
		if node.Right != nil {
			stack = append(stack, *node.Right)
		}
		if node.Left != nil {
			stack = append(stack, *node.Left)
		}
	}
}
