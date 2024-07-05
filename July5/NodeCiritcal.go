package main

type ListNode struct{
	Val int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
    var answer []int = []int{-1, -1}

    if head == nil || head.Next == nil || head.Next.Next == nil {
        return answer
    }

    var prePosition, curPosition, firstPosition, position int = -1, -1, -1, 1

    pre := head
    curr := head.Next

    for curr != nil && curr.Next != nil {
        next := curr.Next

        if (curr.Val < pre.Val && curr.Val < next.Val) ||
           (curr.Val > pre.Val && curr.Val > next.Val) {

            prePosition = curPosition
            curPosition = position

            if firstPosition == -1 {
                firstPosition = position
            }

            if prePosition != -1 { 
                if answer[0] == -1 {
                    answer[0] = curPosition - prePosition
                } else {
                    answer[0] = min(answer[0], curPosition - prePosition)
                }
                answer[1] = position - firstPosition
            }
        }

        pre = curr
        curr = next
        position++
    }

    return answer
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
