# Intersection of Arrays: July 2

## Overview

The `intersect` function finds the intersection of two integer arrays, `nums1` and `nums2`. The intersection of two arrays is defined as the set containing distinct elements common to both arrays.

## Function Signature

```go
func intersect(nums1 []int, nums2 []int) []int
```

- **Input:**
  - `nums1`: An integer array.
  - `nums2`: Another integer array.

- **Output:**
  - Returns an integer array containing elements that are present in both `nums1` and `nums2`.

## Edge Cases

- If either `nums1` or `nums2` is empty, the function returns an empty array because an intersection cannot exist.

## Approach

1. **Frequency Map:** 
   - Create a frequency map (`freqMap`) to count occurrences of each number in `nums1`.

2. **Intersection Calculation:**
   - Iterate through `nums2`. For each number in `nums2`, check if it exists in `freqMap`.
   - If found and the count in `freqMap` is greater than zero, add the number to the `intersection` list and decrement its count in `freqMap`.

3. **Efficiency:**
   - The function uses a single pass through `nums1` to build the frequency map, which takes O(m) time where m is the length of `nums1`.
   - Checking against `nums2` and constructing the `intersection` list also takes O(n) time where n is the length of `nums2`.
   - Overall time complexity is O(m + n), where m and n are the lengths of `nums1` and `nums2`, respectively.
   - Space complexity is O(min(m, n)) due to the space used by the frequency map and the `intersection` list.

## Conclusion

The `intersect` function efficiently computes the intersection of two arrays using a frequency map to track element occurrences. It handles various edge cases and provides a straightforward approach to finding common elements between arrays.
