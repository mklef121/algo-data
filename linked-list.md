### Linked List

There are two types of linked 
- Singly linked list
- Doubly linked list


#### Whats the problem with arrays
With arrays there is a certain amount of data that can be assigned to it.
Save for dynamic arrays in languages like javascript and php and *slices* in golang.

But each time a a dynamic array supasses it's capacity, it's size is doubles to accomodate more items.
This operation is O(n) and it can be expensive.

Secondly, insert, delete and search operations in arrays can be expensive O(n) too.

#### Why linked list
Linked lists are less rigid in their storage structure and elements are usually not stored in contiguous locations, hence they need to be stored with additional tags giving a reference to the next or previous (for doubly linked list) element. Linked lists are also preferable when you need to perform frequent insertion or deletion operations at any position, and do not care about random access, sorting, or searching.