package main

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(v *TreeNode, dist int, ans *int ) []int {
    if v == nil {
        return []int{}
    }
    if v.Left == nil && v.Right == nil {
        return []int{1}
    }

    left, right := dfs(v.Left, dist, ans), dfs(v.Right, dist, ans)
    ret := make([]int, 0)
    for _, l := range left {
        for _, r := range right {
            if l + r <= dist {
                *ans++
            } 
        }

         if l + 1 <  dist {
            ret = append(ret, l+1)
        }
    }

    for _, r := range right {
        if r + 1 < dist {
            ret = append(ret, r+1)
        }
    }

    return ret
} 

func countPairs(root *TreeNode, distance int) int {
    ans := 0
    dfs(root, distance, &ans)
    return ans
}