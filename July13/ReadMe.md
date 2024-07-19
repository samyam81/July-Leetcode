# Robot Collision: July 13

This Java class `RobotCollision` simulates a scenario where robots collide based on their positions, directions, and healths, and determines which robots survive after collisions.

## Method: `survivedRobotsHealths`

### Parameters
- `positions`: An array of integers representing initial positions of robots.
- `healths`: An array of integers representing initial healths of robots.
- `directions`: A string representing the movement directions of robots ('L' for left and 'R' for right).

### Returns
- A list of integers representing healths of robots that survived after collisions.

### Explanation

1. **Initialization and Sorting:**
   - Creates a list `list` containing indices of robots.
   - Sorts `list` based on positions of robots using `Collections.sort` with a custom comparator.

2. **Collision Handling:**
   - Uses a stack `stack` to keep track of robots moving right ('R').
   - Iterates through sorted `list`:
     - For robots moving left ('L'):
       - Compares healths of current robot (`healths[l]`) with robots in `stack`.
       - Adjusts healths based on collision rules:
         - If healths are equal, both robots are destroyed.
         - If current robot has higher health, it survives with reduced health of collided robot.
         - If collided robot has higher health, it survives with reduced health of current robot.
     - Pushes robots moving right ('R') onto `stack`.

3. **Survivor List Creation:**
   - Constructs `res` list containing healths of robots that survived (health not equal to 0).

4. **Return:**
   - Returns `res` list containing healths of survived robots.

In this example, with `positions = {3, 1, 5, 7}`, `healths = {4, 6, 2, 5}`, and `directions = "LLRL"`, the method calculates and prints the healths of robots that survive after collisions.

### Considerations

- Ensure that the input arrays `positions` and `healths` are of the same length, and `directions` string length matches the number of robots.
- The method assumes valid inputs and does not handle errors such as invalid directions or unequal array lengths.
- Performance considerations involve sorting (`O(n log n)`) and stack operations (`O(n)`), where `n` is the number of robots.
