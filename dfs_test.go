package algor_lib

import (
	"testing"
)

func TestCreateTree(t *testing.T) {
	arr := []int{11, 2, 23, 34, 55}

	root, err := CreateTree(arr)
	if err != nil {
		t.Fatalf("error:%s", err)
	}

	if root.Val != arr[0] {
		t.Fatalf("wrong root val:%d,should:%d", root.Val, arr[0])
	}

	if root.Right.Val != arr[2] {
		t.Fatalf("wrong root.right val:%d,should:%d", root.Right.Val, arr[2])
	}

	if root.Left.Right.Val != arr[4] {
		t.Fatalf("wrong root.left.right val:%d,should:%d", root.Left.Right.Val, arr[4])
	}
}

func TestDfs(t *testing.T) {
	arr := []int{11, 2, 23, 34, 55}
	root, _ := CreateTree(arr)

	result := make([]int, 0, 10)
	Dfs(root, &result)

	t.Log(result)
	if result[0] != 11 {
		t.Fatalf("result[0]:%d,should:%d", result[0], 11)
	}
	if result[2] != 34 {
		t.Fatalf("result[2]:%d,should:%d", result[2], 34)
	}
	if result[4] != 23 {
		t.Fatalf("result[4]:%d,should:%d", result[4], 23)
	}
}

func TestBfs(t *testing.T) {
	arr := []int{11, 2, 23, 34, 55}
	root, _ := CreateTree(arr)

	result := make([]int, 0, 10)
	Bfs(root, &result)

	t.Log(result)
	if result[0] != 11 {
		t.Fatalf("result[0]:%d,should:%d", result[0], 11)
	}
	if result[2] != 34 {
		t.Fatalf("result[2]:%d,should:%d", result[2], 34)
	}
	if result[4] != 23 {
		t.Fatalf("result[4]:%d,should:%d", result[4], 23)
	}
}
