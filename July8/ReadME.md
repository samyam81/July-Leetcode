# CircularGame : July 8

The `findTheWinner` function simulates a game where `n` players are standing in a circle and every `k-th` player is eliminated until only one player remains. The function determines the position of the last remaining player.

## Function Description

The function `findTheWinner(n int, k int) int` takes two integers:
- `n`: Represents the number of players standing in a circle.
- `k`: Represents the step size for eliminating players.

### Operation:

The function simulates the elimination process:
1. Start with player `i` initialized to `1`.
2. Iterate through the players from `2` to `n`.
3. Compute the new position `i` after each elimination:
   - Increase `i` by `k`.
   - Compute the new position using the formula `(i - 1) % j + 1`, where `j` is the current number of players.
4. Return the final position of the last remaining player.

### Parameters:

- `n`: Number of players initially standing in a circle.
- `k`: Step size for eliminating players.

### Return Value:

- An integer representing the position of the last remaining player.

### Complexity

- **Time Complexity**: O(n)
  - The function iterates through `n-1` players once, adjusting the position based on the elimination step `k`.

- **Space Complexity**: O(1)
  - The function uses a constant amount of extra space for variables regardless of the input size.

