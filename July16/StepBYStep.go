package main

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	s, _ := find(root, startValue, []byte{})
	t, _ := find(root, destValue, []byte{})
	for len(s) > 0 && len(t) > 0 && s[0] == t[0] {
		s = s[1:]
		t = t[1:]
	}
	return strings.Repeat("U", len(s)) + string(t)
}

func find(root *TreeNode, value int, pre []byte) ([]byte, bool) {
	if root == nil {
		return nil, false
	}
	if root.Val == value {
		return pre, true
	}
	if res, ok := find(root.Left, value, append(pre, 'L')); ok {
		return res, true
	}
	return find(root.Right, value, append(pre, 'R'))
}