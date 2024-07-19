# Max Point: July 12

This Go program calculates the maximum gain by removing specific substrings from a given string `s`. The score for removing each substring ("ab" and "ba") can be specified through parameters `x` and `y`.

### Parameters
- `s`: The input string from which substrings are removed to maximize gain.
- `x`: Score for removing the substring "ab".
- `y`: Score for removing the substring "ba".

### Returns
- An integer representing the maximum gain achievable by removing substrings "ab" and "ba" from `s`.

## How It Works

1. The function `maximumGain`:
   - Determines which substring ("ab" or "ba") has a higher score (`x` or `y`).
   - Uses a helper function `helper` to remove the chosen substring from `s` and calculates the corresponding score.
   - Returns the sum of scores obtained from removing both substrings.

2. The helper function `helper`:
   - Takes a string `s`, a substring `sub` ("ab" or "ba"), and a score `x`.
   - Removes occurrences of `sub` from `s` using a stack-based approach.
   - Returns the total score accumulated from removing `sub` and the modified string after removals.

In this example, with `s = "aabbaaxyz"`, `x = 3`, and `y = 4`, the function calculates the maximum gain by removing "ba" first and then "ab", resulting in a total gain of 7.

## Considerations

- Ensure that the input string `s` and scores `x` and `y` are provided appropriately to get the correct maximum gain value.
- The function assumes valid inputs and does not handle errors such as negative scores or invalid strings.
