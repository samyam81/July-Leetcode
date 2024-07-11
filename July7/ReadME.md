# WaterBottle: July 7

The `numWaterBottles` function calculates the maximum number of water bottles you can drink, given `numBottles` full water bottles initially, and `numExchange` empty bottles required for exchanging for a full bottle.

## Function Description

The function `numWaterBottles(numBottles int, numExchange int) int` takes two integers:
- `numBottles`: Represents the initial number of full water bottles.
- `numExchange`: Represents the number of empty bottles required to exchange for one full water bottle.

### Operation:

The function simulates the drinking process:
1. Start with `numBottles` full bottles.
2. Exchange empty bottles until you can no longer exchange (`numBottles < numExchange`).
3. Each exchange consumes `numExchange` empty bottles and gives you one additional full bottle.
4. Return the total number of bottles consumed, including initial and additional bottles obtained through exchanges.

### Complexity

- **Time Complexity**: O(log(numBottles) * numExchange)
  - The time complexity is primarily determined by how many times we can perform the exchange operation (`numBottles / numExchange` times) until the number of bottles is less than `numExchange`. This is logarithmic with respect to `numBottles`.

- **Space Complexity**: O(1)
  - The function uses a constant amount of extra space for variables regardless of the input size.

