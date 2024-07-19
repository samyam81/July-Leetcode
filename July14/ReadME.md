# Chemical Formula Parser and Atom Counter : July 14

This Go code snippet provides functionality to parse a chemical formula string and count the occurrences of each atom along with their quantities.

## Overview

The `countOfAtoms` function is the entry point which takes a chemical formula string as input and returns a string representation of atoms and their counts in sorted order.

## Implementation Details

### StateMachine Struct

The `StateMachine` struct manages the parsing state and processes each character in the chemical formula string.

#### Fields
- `formula`: Stores the input chemical formula string.
- `curEle`: Tracks the current element being processed.
- `state`: Represents the current state of the state machine (`Open` or `Close`).
- `stack`: A stack to manage nested element groups.

#### Methods

- **NewStateMachine**: Initializes a new `StateMachine` instance.
- **Len**: Returns the length of the `formula` string.
- **PeakGroup**: Returns the top element group from the stack.
- **Process**: Controls the flow of processing based on the current state (`Open` or `Close`).
- **processOpen**: Handles processing when encountering an opening parenthesis or element.
- **processClose**: Handles processing when encountering a closing parenthesis.
- **scanElement**: Scans and returns an element from the formula string.
- **scanNumber**: Scans and returns a number from the formula string.

### ElementGroup Struct

The `ElementGroup` struct represents a mapping of elements to their counts.

#### Methods

- **Incr**: Increments the count of a specified element.
- **MergeWithK**: Merges the counts of another `ElementGroup` multiplied by a factor `k`.
- **String**: Converts the `ElementGroup` into a formatted string representation.


Replace `"H2O"` with any valid chemical formula string to see the output of atom counts.

## Considerations

- Ensure the input formula string follows valid chemical formula syntax.
- The code assumes that the chemical formula string is correctly formatted and does not handle syntax errors.
- Performance considerations include linear time complexity relative to the length of the formula string.
