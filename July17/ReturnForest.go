package main

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    ret, r := dfs(root, to_delete)
    if r != nil {
        ret = append(ret, r)
    }
    return ret
}

func dfs(root *TreeNode, to_delete []int) ([]*TreeNode, *TreeNode) {
    if root == nil {
        return []*TreeNode{}, nil
    }
    roots := make([]*TreeNode, 0)
    leftRoots, left := dfs(root.Left, to_delete)
    rightRoots, right := dfs(root.Right, to_delete)
    root.Left = left
    root.Right = right

    if len(leftRoots) != 0 {
        roots = append(roots, leftRoots...)
    }
    if len(rightRoots) != 0 {
        roots = append(roots, rightRoots...)
    }
    if ok := findValue(root, to_delete); ok {
        if root.Left != nil {
            roots = append(roots, root.Left)
        }
        if root.Right != nil {
            roots = append(roots, root.Right)
        }
        return roots, nil
    }
    return roots, root
}

func findValue(root *TreeNode, to_delete []int) bool {
    for _, v := range to_delete {
        if v == root.Val {
            return true
        }
    }
    return false
}