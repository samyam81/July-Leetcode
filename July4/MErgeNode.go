package main


 type ListNode struct {
     Val int
     Next *ListNode
 }

func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    modify := dummy

    current := head

    for current != nil {
        sum := 0
        for current != nil && current.Val != 0 {
            sum += current.Val
            current = current.Next
        }

        if sum != 0 { 
            newNode := &ListNode{Val: sum}
            modify.Next = newNode
            modify = modify.Next
        }

        for current != nil && current.Val == 0 {
            current = current.Next
        }
    }

    return dummy.Next
}