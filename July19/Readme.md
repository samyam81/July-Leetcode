# Finding Lucky Numbers in Matrix: July 19

This Go code snippet finds "lucky" numbers in a matrix. A "lucky" number in this context is a number that is both the smallest in its row and the largest in its column.

## Overview

The function `luckyNumbers` takes a 2D matrix `matrix` of integers as input. It identifies and returns a list of "lucky" numbers from the matrix.

## Implementation Details

### Variables

- `minArr`: A slice to store the minimum numbers from each row of the matrix.
- `maxArr`: A slice to store the maximum numbers from each column of the matrix.
- `res`: A slice to store the final list of "lucky" numbers found in both `minArr` and `maxArr`.

### Steps

1. **Find Minimum in Rows**: Iterates through each row of the matrix to find the minimum number and stores it in `minArr`.
2. **Find Maximum in Columns**: Iterates through each column of the matrix to find the maximum number and stores it in `maxArr`.
3. **Identify Lucky Numbers**: Compares elements from `minArr` and `maxArr`. If an element is found in both lists, it is added to `res`.

### Edge Cases

- Handles matrices of varying sizes and dimensions.
- Ensures correct identification of "lucky" numbers even when there is only one number meeting the criteria.

## Considerations

- Ensure the input matrix `matrix` contains valid integer values.
- The function assumes that there is at least one "lucky" number in the matrix based on the problem definition.
- Performance considerations include time complexity relative to the size of the matrix due to nested iterations for finding minimums and maximums.
