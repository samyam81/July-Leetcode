# Reverse Substring : July 11

The `reverseParentheses` function reverses substrings enclosed within parentheses in a given string `s`.

## Function Description

The function `reverseParentheses(s string) string` takes a string `s` as input and performs the following operations:
1. Uses a stack to keep track of the indices of opening parentheses '(' encountered in the string.
2. When a closing parenthesis ')' is encountered, retrieves the corresponding opening parenthesis index from the stack.
3. Reverses the substring between the retrieved opening parenthesis index and the current closing parenthesis index.
4. Removes parentheses from the string and returns the modified string where all substrings within parentheses are reversed.

### Parameters:

- `s`: A string containing alphanumeric characters and parentheses.

### Operation:

1. Initialize an empty stack and convert the input string `s` into a byte slice `chars` for mutability.
2. Iterate through each character in `chars`.
   - If an opening parenthesis '(' is encountered, push its index onto the stack.
   - If a closing parenthesis ')' is encountered, pop the stack to get the corresponding opening parenthesis index.
   - Use the `reverse` function to reverse the substring between the retrieved indices.
3. Construct a new byte slice `result` to hold the characters of the modified string without parentheses.
4. Return the string representation of `result`.

## Complexity

- **Time Complexity**: O(n^2)
  - In the worst case, each character in the string may involve reversing a substring, resulting in nested iterations over portions of the string, leading to a quadratic time complexity.
  
- **Space Complexity**: O(n)
  - The function uses additional space for the stack and the result array, both proportional to the length of the input string `s`.
