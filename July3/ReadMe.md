# Minimum Difference Between Largest and Smallest Value in Three Moves: July 3

This Go code defines a function `minDifference` that calculates the minimum difference between the largest and smallest elements after removing exactly three elements from an array of integers (`nums`). Here’s how the function works:

#### Function Signature
```go
func minDifference(nums []int) int
```
- `nums`: A slice of integers.
- Returns: The minimum difference between the largest and smallest elements after removing exactly three elements from `nums`.

## Algorithm Explanation
1. **Sorting**:
   - Sort the array `nums` in non-decreasing order using `sort.Ints(nums)`.

2. **Calculating Differences**:
   - Calculate the potential minimum differences by considering removing exactly three elements:
     - `nums[len(nums)-4] - nums[0]`: Remove the first three elements.
     - `nums[len(nums)-1] - nums[3]`: Remove the last three elements.
     - `nums[len(nums)-2] - nums[2]`: Remove the first, second, and last elements.
     - `nums[len(nums)-3] - nums[1]`: Remove the first, last, and second-to-last elements.
   - Update `minVal` with the minimum of these calculated differences using the `min` function.

3. **Edge Case**:
   - If the length of `nums` is less than or equal to 4 (`len(nums) <= 4`), return `0` because it's not possible to remove three elements.

4. **Return Result**:
   - Return `minVal`, which holds the minimum difference after considering all possible removals.

## Helper Function
```go
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```
- Returns the minimum of two integers `a` and `b`.

## Conclusion
The `minDifference` function efficiently computes the minimum difference between the largest and smallest elements of an integer array after removing exactly three elements. By leveraging sorting and calculating differences in a systematic way, it ensures that the result is accurate and efficient, adhering to the problem requirements effectively.