# Binary Tree Construction from Description : July 15

This Go code snippet constructs a binary tree from a description provided as a 2D array `desc` containing parent-child relationships and whether each child is left or right.

## Overview

The function `createBinaryTree` takes a 2D array `desc` as input, where each inner array `vals` represents:
- `vals[0]`: Parent node value.
- `vals[1]`: Child node value.
- `vals[2]`: 1 if the child is the left child of the parent, 0 if right.

The function builds and returns the root of the binary tree based on the provided description.

## Implementation Details

### TreeNode Struct

The `TreeNode` struct represents a node in the binary tree:
- `Val`: Integer value stored in the node.
- `Left`: Pointer to the left child node.
- `Right`: Pointer to the right child node.

### createBinaryTree Function

#### Parameters

- `desc`: A 2D array where each inner array describes a parent-child relationship and the position of the child.

#### Variables

- `mp`: A map to store nodes (`TreeNode`) based on their values.
- `vis`: A map to track visited nodes to identify the root node.

#### Steps

1. **Node Initialization**: Iterates through `desc` to create and initialize parent and child nodes (`TreeNode`) in the `mp` map.
2. **Tree Construction**: Establishes parent-child relationships based on whether the child is left or right.
3. **Root Identification**: Identifies the root node by checking nodes that are not visited as children.

#### Return Value

- Returns the root node (`*TreeNode`) of the constructed binary tree.

Replace `desc` with any valid parent-child relationship array to construct and print the root of the binary tree.

## Considerations

- Ensure the input `desc` array is correctly formatted with valid parent-child relationships.
- The function assumes the description provided correctly represents a binary tree structure without cycles.
- Performance considerations include linear time complexity relative to the size of the `desc` array for tree construction.
