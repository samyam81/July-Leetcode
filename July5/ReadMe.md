# Find the Minimum and Maximum Number of Nodes Between Critical Points: Ju;y 5

This Go code defines a function `nodesBetweenCriticalPoints` that operates on a singly linked list (`ListNode`). The function identifies and computes distances between critical points in the list where the node values exhibit a local extremum (either a local minimum or maximum).

### ListNode Structure
```go
type ListNode struct {
	Val  int
	Next *ListNode
}
```
- `ListNode` represents a node in a singly linked list, containing an integer value (`Val`) and a pointer to the next node (`Next`).

## Function Signature
```go
func nodesBetweenCriticalPoints(head *ListNode) []int
```
- `head`: Pointer to the head of the linked list.
- Returns: A slice of integers containing two values:
  - The shortest distance between consecutive critical points.
  - The total number of nodes between the first and the last critical points.

## Algorithm Explanation
1. **Initialization**: 
   - Initialize `answer` as `[]int{-1, -1}` to store the result.
   - Check if the list is empty or has less than three nodes (`head == nil || head.Next == nil || head.Next.Next == nil`), return the initialized `answer`.

2. **Traversal**:
   - Use two pointers (`pre` and `curr`) to traverse the list.
   - Maintain `position` to track the current node's position.
   - `prePosition` and `curPosition` track positions of previous and current critical points, respectively.
   - `firstPosition` records the position of the first critical point encountered.

3. **Identifying Critical Points**:
   - Iterate through the list while `curr` and `curr.Next` are not `nil`.
   - Determine if `curr` is a critical point:
     - It's a local minimum if `curr.Val` is less than both `pre.Val` and `next.Val`.
     - It's a local maximum if `curr.Val` is greater than both `pre.Val` and `next.Val`.
   - Update `prePosition`, `curPosition`, and `firstPosition` accordingly.

4. **Calculating Distances**:
   - After identifying a pair of critical points (`prePosition` and `curPosition`):
     - Calculate the distance between them and update `answer[0]`.
     - Update `answer[1]` with the total number of nodes between the first and current critical points.

5. **Returning Result**:
   - After traversing the list, `answer` contains the required distances or remains `[-1, -1]` if no critical points are found.

## Helper Function
```go
func min(a, b int) int
```
- Returns the minimum of two integers `a` and `b`.

## Conclusion
This function `nodesBetweenCriticalPoints` efficiently identifies critical points (local minima and maxima) in a singly linked list and computes the distances between them, providing insights into the structure of the list based on its values.