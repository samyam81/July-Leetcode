# Three Consecutive Odds: July 1

### Approach:

The main goal of the function `threeConsecutiveOdds` is to determine if an array of integers contains three consecutive odd numbers. The approach taken can be summarized as follows:

1. **Edge Case Handling**: 
   - If the length of the array `arr` is less than or equal to 2, it immediately returns `false` because it's impossible to have three consecutive numbers in such a short array.

2. **Iteration through Array**: 
   - The function iterates through each element of the array using a `for` loop.
   - For each element, it checks if the number is odd (`arr[i] % 2 == 1`).
   - If the number is odd, it increments a `count` variable that keeps track of consecutive odd numbers encountered.
   - If `count` reaches 3 at any point during the iteration, it means three consecutive odd numbers have been found, and the function returns `true`.

3. **Resetting Counter**: 
   - If an even number is encountered during the iteration (`arr[i] % 2 == 0`), it resets the `count` variable to 0 because consecutive sequence is broken.

4. **Final Check**: 
   - If the loop completes without finding three consecutive odd numbers, the function returns `false`.

### Complexity:

- **Time Complexity**: The time complexity of the function is `O(n)`, where `n` is the number of elements in the input array `arr`. This is because the function iterates through the array once (linear time complexity).

- **Space Complexity**: The space complexity is `O(1)` (constant space), as the function only uses a fixed amount of extra space regardless of the size of the input array. Specifically, it uses a single variable (`count`) to keep track of the consecutive odd numbers.

### Working Procedure:

1. **Initialization**: 
   - The function starts by initializing a variable `count` to 0.

2. **Edge Case Check**: 
   - If the length of the input array `arr` is 0, 1, or 2, the function immediately returns `false` because it's impossible to find three consecutive odd numbers.

3. **Iterating through the Array**: 
   - It loops through each element of the array `arr`.
   - For each element, it checks if it's odd.
     - If it's odd, it increments the `count` variable.
     - If `count` reaches 3, the function returns `true` because it has found three consecutive odd numbers.
     - If it's even, it resets `count` to 0 because the sequence of consecutive odds is broken.

4. **Completion**: 
   - If the loop completes without finding three consecutive odd numbers, the function returns `false`.

### Summary:

The function efficiently checks for the presence of three consecutive odd numbers in an array using a straightforward linear scan approach. It optimally uses constant extra space and operates in linear time relative to the size of the input array. This makes it both efficient and straightforward for the given task.