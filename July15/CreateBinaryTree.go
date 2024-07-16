package main

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func createBinaryTree(desc [][]int) *TreeNode {
    mp := make(map[int]*TreeNode)
    vis := make(map[int]bool)
    
    for _, vals := range desc {
        parent, child, isLeft := vals[0], vals[1], vals[2]

        if _, ok := mp[parent]; !ok {
            mp[parent] = &TreeNode{Val: parent}
        }

        if _, ok := mp[child]; !ok {
            mp[child] = &TreeNode{Val: child}
        }

        if isLeft == 1 {
            mp[parent].Left = mp[child]
        } else {
            mp[parent].Right = mp[child]
        }

        vis[child] = true
    }

    for _, vals := range desc {
        if !vis[vals[0]] {
            return mp[vals[0]]
        }
    }

    return nil
}