## What are trees ?
Tree represents the nodes connected by edges. They are nonlinear data structures


### Types:

- Binary tree: Binary Tree is a special datastructure used for data storage purposes. A binary tree has a special condition that each node can have a maximum of two children

There are perfect binary trees: has all child nodes complete.

To find out the number of Nodes at each level do

Level-0:  2^0 = 1
Level-1: 2^1 = 2
Level-2: 2^2 = 4
Level-3: 2^3 = 8 
Level-4: 2^4 = 16 -- This means we will need to check 16 nodes if a value is been searched
...
Level-h: 2^h-1

From the Above, we can see that the number of nodes can be calculated using 

nodes = 2^(h-1);
For perfect binary tree, all nodes above the last level + 1 =  the nodes on the last level


So where does the terminology O(logn) come from ?

In mathematics, Log 100 = 2

So replacing

nodes = Log(h-1)

Since the rule of computing time is to drop insignificant numbers

we can say **nodes = Log(h)**


### Binary Search Trees

A Binary Search Tree (BST) is a tree in which all the nodes follow the below-mentioned properties −

- The value of the key of the left sub-tree is less than the value of its parent (root) node's key.

- The value of the key of the right sub-tree is greater than or equal to the value of its parent (root) node's key.

Thus, BST divides all its sub-trees into two segments; the left sub-tree and the right sub-tree.

`left_subtree (keys) < node (key) ≤ right_subtree (keys)`

### Problem with Binary search tree.

A BST might have insertions that take a particular order and keep increasing to a particular direction.

This will in turn form a linked list which can gradually change it's complexity to O(n) instead of O(Log-n). This is a condition
called **Unbalanced Binary Search Tree**. Insertion, deletion and lookup for unbalanced binary search trees becomes `O(n)`