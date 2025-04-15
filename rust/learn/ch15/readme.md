## Smart Pointers
In Rust, a **pointer** is a variable that holds a memory address referring to some data. The most common pointer is a **reference** (`&`), which simply borrows data without taking ownership and has no overhead.

In contrast, **smart pointers** are more advanced. They behave like pointers but also store additional metadata and capabilities. Unlike basic references, smart pointers often **own** the data they point to, offering more functionality. For example, a **reference-counted smart pointer** keeps track of how many references exist to a value and automatically cleans up when no references remain.

Rust includes several smart pointers in its standard library, such as `Box<T>`, `Rc<T>`, and `RefCell<T>`. Interestingly, types like `String` and `Vec<T>` are also considered smart pointers—they own their memory, provide extra features (like tracking capacity), and ensure safety (e.g., valid UTF-8 in `String`).

Rust’s ownership model adds a key distinction: **references borrow data**, while **smart pointers often own it**, enabling powerful memory management without a garbage collector.


Smart pointers in Rust are typically structs that implement two key traits:

* **`Deref`**: Enables the smart pointer instance to be treated like a regular reference (`&`), allowing code to work with both types seamlessly.
* **`Drop`**: Lets you define custom cleanup code that executes automatically when the smart pointer instance goes out of scope.

These traits are fundamental to the utility and behavior of smart pointers.

### Using `Box<T>` to Point to Data on the Heap

The most straightforward smart pointer is a box, whose type is written `Box<T>`. Boxes allow you to store data on the **heap** rather than the **stack**. What remains on the stack is the pointer to the heap data. 

`Box<T>` in Rust incurs minimal performance overhead (heap allocation vs. stack). It's primarily useful when:

* You need a fixed-size representation for a type whose size is dynamic at compile time.
* You want to transfer ownership of large data efficiently without copying.
* You need to own a value based on a trait it implements, rather than its concrete type.

Transferring ownership of large data on the stack can be slow due to copying. Using `Box<T>` moves the data to the heap, making ownership transfer much faster as only the small pointer on the stack is copied. The large data remains in its heap location.

### Using a `Box<T>` to Store Data on the Heap
Before discussing using `Box<T>` for heap storage, let's look at its syntax and how to work with the values it holds.

```rust
fn main() {
    let b = Box::new(5);
    println!("b = {b}");
}
```

This code creates a `Box` named `b` that allocates the integer `5` on the heap. Printing `b` directly accesses the underlying value, just as if it were on the stack. When `b` goes out of scope at the end of `main`, both the `Box` (on the stack) and the `5` it points to (on the heap) are automatically deallocated.

### Enabling Recursive Types with Boxes

Recursive types, where a value can contain another value of the same type, pose a size determination problem for Rust at compile time due to potential infinite nesting. To handle this, we can use `Box<T>` within the recursive type definition because `Box<T>` has a known, fixed size (the size of a pointer). This allows Rust to determine the memory layout.

A **cons list** is a recursive data structure from Lisp, built using the `cons` (construct) function. It creates linked lists by nesting pairs. For example, the list `1, 2, 3` would be represented as:

```
(1, (2, (3, Nil)))
```

Each item has:
- A value
- A reference to the next item (or `Nil` for the end)

This structure is built recursively by calling `cons` with a value and the rest of the list.

While cons lists are fundamental in Lisp, they’re not commonly used in Rust. Rust typically uses `Vec<T>` for lists. However, cons lists serve as a great example for learning how to define recursive types in Rust using `Box`, since they’re simple and illustrate the concept well.

Consider this Rust code that tries to define a list of values `1, 2, 3` using a recursive `List` type:

```rust
use crate::List::{Cons, Nil};

fn main() {
    let list = Cons(1, Cons(2, Cons(3, Nil)));
}

enum List {
    Cons(i32, List),
    Nil,
}
```

This defines a `List` where:
- `Cons` holds a value and another `List`
- `Nil` marks the end of the list

However, this code **won’t compile**. Rust throws an error saying the type has **infinite size**. Why?

Because the `List` type contains itself directly (not through a pointer), the compiler can’t determine how much memory to allocate—each `Cons` wraps another `Cons`, theoretically forever. Rust requires every type to have a known size at compile time, and directly recursive types like this violate that rule.

To fix this, you need to use a smart pointer like `Box` to break the infinite size chain. We'll get to that next.

Let's see how Rust calculates enum size & handles recursion. Consider this enum:

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}
```

Rust determines how much memory to allocate for a `Message` by checking **all its variants** and choosing the size of the **largest** one—because only one variant is active at any time.

- `Quit` needs **0 bytes** (no data).
- `Move` needs **8 bytes** (two `i32`s).
- `Write(String)` needs **the size of a `String`**, which includes a pointer, length, and capacity (usually 24 bytes on 64-bit systems).
- `ChangeColor(i32, i32, i32)` needs **12 bytes** (three `i32`s).

Among these, `Write(String)` usually takes the most space because a `String` is a smart pointer that manages heap data. So, Rust allocates enough space to fit the **largest possible variant**—in this case, likely the `Write` variant.

#### Using `Box<T>` to Get a Recursive Type with a Known Size

```rust
enum List {
    Cons(i32, Box<List>),
    Nil,
}

use crate::List::{Cons, Nil};

fn main() {
    let list = Cons(1, Box::new(Cons(2, Box::new(Cons(3, Box::new(Nil))))));
}
```

A `Box<T>` is a smart pointer with a **fixed size**, regardless of the size of the data it points to. This means that the Rust compiler always knows how much memory a `Box<T>` needs.
The `Cons` variant needs the size of an `i32` plus the space to store the `box’s` pointer data. The `Nil` variant stores no values, so it needs less space than the `Cons` variant. We now know that any List value will take up the size of an `i32` plus the size of a box’s pointer data. 

This puts each list node on the **heap**, and the `Box` acts like a pointer linking them together. So while it still logically forms a recursive structure, it's implemented as a sequence of heap-allocated items **linked side-by-side**, not nested within each other.


### Treating Smart Pointers Like Regular References with the `Deref` Trait
Implementing the `Deref` trait allows you to customize the behavior of the dereference operator `*`. By implementing Deref in such a way that a smart pointer can be treated like a regular reference, you can write code that operates on references and use that code with smart pointers too.

```rust
// approach 1
let x = 5;
let y = &x;

assert_eq!(5, x);
assert_eq!(5, *y);

// approach 2
let x = 5;
let y = Box::new(x);

assert_eq!(5, x);
assert_eq!(5, *y);
```

In  approach 2, the key difference from approach 1 is that `y` is a `Box<T>`—it owns a **copied value of `x`** stored on the heap, rather than just referencing `x` like before.

Despite this, we can still use the **dereference operator (`*`)** on `y` just like we did with a reference. This works because `Box<T>` implements the `Deref` trait, which allows us to access the value it points to.

```rust
fn main() {

}

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}


let x = 5;
let y = MyBox::new(x);

assert_eq!(5, x);
assert_eq!(5, *y);
```

Look at the code above, Without the `Deref` trait, the compiler can only dereference `&` references. The `deref` method gives the compiler the ability to take a value of any type that implements `Deref` and call the `deref` method to get a `&` reference that it knows how to dereference.
The `deref` method returns a **reference** to the inner value, not the value itself, to avoid **moving** the value out of `self`.

If `deref` returned the actual value, Rust would **transfer ownership**, which isn’t usually what we want when using the `*` operator. Most of the time, we just want to **access** the inner value, not take it.
When you write `*y`, Rust automatically translates it to: `*(y.deref())`

This happens **once**—Rust doesn’t keep applying `.deref()` recursively. After one level of indirection, the `*` operator gets applied to the reference returned by `deref`.

#### Implicit Deref Coercions with Functions and Methods

**Deref coercion** is a Rust feature that converts a reference to a type that implements the `Deref` trait into a reference to another type by calling the `deref` method automatically behind the scenes. A sequence of calls to the `deref` method converts the type we provided into the type the parameter needs (the calls can be recursive until till the expected type can be gotten or not).

For example:

```rust
fn greet(name: &str) {
    println!("Hello, {name}!");
}

let name = String::from("Alice");
greet(&name); // Rust automatically converts &String to &str

// using the MyBox deref implemented earlier
let m = MyBox::new(String::from("Rust")); // deref is called twice, first on MyBox and secondly on String
hello(&m);
```

Here, Rust sees `&String` but expects `&str`, so it calls `.deref()` for us to make the conversion.

### Running Code on Cleanup with the `Drop` Trait

