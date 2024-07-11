# Average Waiting Time: July 9

The `averageWaitingTime` function calculates the average waiting time for customers based on their arrival and service times.

## Function Description

The function `averageWaitingTime(customers [][]int) float64` takes a 2D slice `customers` where each inner slice represents a customer with two integers:
- `arrival`: The time when the customer arrives.
- `service`: The time required to serve the customer.

### Operation:

The function simulates the service process:
1. Initialize `totalWaitingTime` to zero and `currentTime` to zero.
2. Iterate through each customer in `customers`.
3. For each customer:
   - Retrieve their arrival time (`arrival`) and service time (`service`).
   - If `currentTime` is less than `arrival`, update `currentTime` to `arrival` (wait until the customer arrives if necessary).
   - Calculate the waiting time as `currentTime - arrival`.
   - Add the waiting time and service time to `totalWaitingTime`.
   - Update `currentTime` to reflect when the customer finishes being served (`currentTime + service`).
4. Compute the average waiting time by dividing `totalWaitingTime` by the number of customers (`len(customers)`).

### Parameters:

- `customers`: A 2D slice where each inner slice contains two integers representing a customer's arrival and service times.

### Return Value:

- A float64 value representing the average waiting time for all customers.

### Complexity

- **Time Complexity**: O(n)
  - The function iterates through each customer exactly once, where `n` is the number of customers.

- **Space Complexity**: O(1)
  - The function uses a constant amount of extra space for variables regardless of the input size.

