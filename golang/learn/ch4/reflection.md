## What is reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming. 

- Why was reflection included in Go? :
reflection allows you to dynamically learn the type of an arbitrary object along with information about its structure. For example  `fmt.Println()` is clever enough to understand the data types of its parameters and act accordingly? Well, behind the scenes, the fmt package uses reflection to do that.
- When should I use reflection? :
reflection allows you to handle and work with data types that do not exist at the time at which you write your code but might exist in the future, which is when we use an existing package with user-defined data types. Secondly,  reflection might come in handy when you have to work with data types that do not implement a common interface and therefore have an uncommon or unknown behavior.