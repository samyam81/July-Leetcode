# Pass the Pillow: July 6

The `passThePillow` function simulates a game where a pillow is passed among `n` participants in a circular manner. It determines which participant is holding the pillow after `time` seconds have passed.

## Function Signature

```go
func passThePillow(n int, time int) int
```

### Parameters

- `n` (int): Number of participants.
- `time` (int): Time in seconds after which the pillow's position is queried.

### Returns

- (int): The participant number (1-based index) who is holding the pillow after `time` seconds.

## Explanation

1. **Total Turns Calculation (`N`)**:
   - The total number of turns (`N`) the pillow can make before it returns to the starting participant is `2*n - 2`. This accounts for each participant except the starting one twice (once going forward and once coming back).

2. **Position Calculation (`x`)**:
   - Compute `x = time % N` to find out the exact position within the circular sequence.

3. **Determining Participant**:
   - If `x < n`, the participant holding the pillow is `1 + x`.
   - If `x >= n`, the participant holding the pillow is `1 + (N - x)`.

   This ensures that the result is always a valid participant number (1 through `n`).

## Time Complexity

- The time complexity of the `passThePillow` function is O(1). This is because the function performs a fixed number of arithmetic operations and conditional checks, independent of the size of `n` or `time`.

## Space Complexity

- The space complexity of the `passThePillow` function is O(1). The function only uses a constant amount of extra space for storing variables `N`, `x`, and the return value, regardless of the input values of `n` and `time`.