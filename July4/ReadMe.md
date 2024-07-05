# Merge Nodes in Between Zeros: July 4

This Go code defines a function `mergeNodes` that processes a singly linked list (`ListNode`) and merges consecutive nodes with non-zero values into a new linked list. Here’s an explanation of how the function works:

### ListNode Structure
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```
- `ListNode` represents a node in a singly linked list, containing an integer value (`Val`) and a pointer to the next node (`Next`).

### Function Signature
```go
func mergeNodes(head *ListNode) *ListNode
```
- `head`: Pointer to the head of the input linked list.
- Returns: Pointer to the head of a new linked list where consecutive nodes with non-zero values from the input list are merged.

#### Algorithm Explanation
1. **Initialization**: 
   - Create a `dummy` node to simplify the handling of the result linked list.
   - `modify` pointer is used to modify the `Next` field of nodes in the result linked list.

2. **Traversal**: 
   - Use `current` pointer to traverse through the input linked list starting from `head`.

3. **Merging Process**:
   - While `current` is not `nil`, sum up consecutive nodes with non-zero values:
     - Initialize `sum` to accumulate values.
     - Traverse through nodes with non-zero values (`current.Val != 0`), adding their values to `sum`.
     - Create a new node with `sum` as its value (`newNode := &ListNode{Val: sum}`).
     - Append `newNode` to the result linked list using `modify.Next` and move `modify` to `newNode`.
   
4. **Skipping Zeroes**:
   - After summing up non-zero values, skip over nodes with zero values (`current.Val == 0`) to proceed to the next sequence of non-zero values.

5. **Return Result**:
   - Return `dummy.Next`, which points to the head of the merged linked list excluding the initial dummy node.

### Edge Cases:
- If the input linked list is empty (`head == nil`), the function returns `nil`.
- Handles sequences of zeroes by skipping over them effectively.

## Conclusion
The `mergeNodes` function efficiently merges consecutive nodes with non-zero values from the input linked list into a new linked list. It utilizes a dummy node for ease of manipulation and ensures that the result linked list excludes nodes with zero values, effectively condensing sequences of non-zero values into single nodes with their sums.