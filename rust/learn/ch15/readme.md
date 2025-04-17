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

The `Drop` trait in Rust allows you to define custom cleanup logic that runs automatically when a value goes out of scope. This is crucial for smart pointers to manage resources like deallocating heap memory (as `Box<T>` does). Unlike some languages where manual resource management is required and prone to errors, Rust's `Drop` trait ensures automatic cleanup, preventing leaks by guaranteeing that your specified `drop` method (taking a mutable reference to `self`) is executed when an instance is no longer needed.

```rust
struct CustomSmartPointer {
    data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data `{}`!", self.data);
    }
}

fn main() {
    let c = CustomSmartPointer {
        data: String::from("my stuff"),
    };
    let d = CustomSmartPointer {
        data: String::from("other stuff"),
    };
    println!("CustomSmartPointers created.");
}
```

When you run the code above, you will notice the following printed output

```txt
CustomSmartPointers created.
Dropping CustomSmartPointer with data `other stuff`!
Dropping CustomSmartPointer with data `my stuff`!
```

We’re printing some text here to demonstrate visually when Rust will call drop. Variables are dropped in the reverse order of their creation, so `d` was dropped before `c`. 

### Dropping a Value Early with `std::mem::drop`

While Rust automatically runs a type’s `Drop` logic when it goes out of scope—which is the main benefit of using the `Drop` trait—there are rare cases where you might want to trigger cleanup **early**.

For example, if you're working with smart pointers that manage **locks**, you may want to **release the lock** before the end of a scope so other code can acquire it. However, you **can’t manually call** the `drop` method from the `Drop` trait—it’s reserved for Rust’s internal use.

Instead, Rust provides the `std::mem::drop` function. You can call `drop(value)` to force a value to be cleaned up immediately.

```rust
fn main() {
    let c = CustomSmartPointer {
        data: String::from("some data"),
    };
    println!("CustomSmartPointer created.");
    drop(c);
    println!("CustomSmartPointer dropped before the end of main.");
}
```

### `Rc<T>`, the Reference Counted Smart Pointer

In most cases, Rust’s ownership system is straightforward—each value has a single, clear owner. However, some situations, like graph structures, require **multiple ownership**. For example, multiple edges might point to the same node, and that node shouldn’t be dropped until **all** references to it are gone.

Rust handles this through the `Rc<T>` type (short for **Reference Counted**). It enables **shared ownership** by keeping track of how many references exist to a value. When the count drops to zero, the value is automatically cleaned up.

You can think of `Rc<T>` like a TV in a shared living room: the first person in turns it on, and it stays on as long as someone is watching. When the last person leaves, the TV is turned off.

Use `Rc<T>` when:
- You want to share **read-only** heap data across different parts of your program.
- You **don’t know** at compile time which part will stop using the data last.

> `Rc<T>` is only safe for **single-threaded** use. For multi-threading, use `Arc<T>` instead.

```rust
enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use crate::List::{Cons, Nil};
use std::rc::Rc;

fn main() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    let b = Cons(3, Rc::clone(&a));
    let c = Cons(4, Rc::clone(&a));
}
```

This can recode the implementation we have at [recursive boxes](#enabling-recursive-types-with-boxes)

The implementation of Rc::clone doesn’t make a deep copy of all the data like most types’ implementations of clone do. The call to `Rc::clone` only increments the reference count, which doesn’t take much time. Deep copies of data can take a lot of time. By using `Rc::clone` for reference counting, we can visually distinguish between the **deep-copy** kinds of clones and the kinds of clones that increase the reference count.

### `RefCell<T>` and the Interior Mutability Pattern

**Interior mutability** is a Rust design pattern enabling data mutation even with immutable references, typically forbidden by borrowing rules. It achieves this using `unsafe` code within a data structure to bypass the compiler's borrow checks. This pattern is employed only when runtime borrow rule adherence can be guaranteed, despite the compiler's inability to verify it. The `unsafe` code is encapsulated within a safe API, maintaining the immutability of the outer type.

#### Enforcing Borrowing Rules at Runtime with `RefCell<T>`

Unlike `Rc<T>`, which enables **shared ownership**, `RefCell<T>` provides **single ownership** but allows **interior mutability**—meaning you can mutate the data it holds even when it's behind an immutable reference.

So how is `RefCell<T>` different from `Box<T>`? The key difference lies in **when** Rust enforces borrowing rules:

- With `Box<T>` and references, borrowing rules are enforced **at compile time**.
- With `RefCell<T>`, borrowing rules are enforced **at runtime**.

This means that breaking the rules with a `Box<T>` causes a **compiler error**, while doing so with `RefCell<T>` results in a **runtime panic**.

#####  Why use `RefCell<T>`?
While compile-time checking is faster and catches errors early, it can be too restrictive in some valid use cases the compiler can't verify. `RefCell<T>` is useful when:
- You're confident your code follows borrowing rules.
- The compiler can't prove it.
- You still want memory safety, but with more flexibility.

Just like `Rc<T>`, `RefCell<T>` is intended for **single-threaded** contexts. If used in multithreaded code, you'll get a compile-time error.

#### A Use Case for Interior Mutability

**Mock Objects**

---

In testing, developers often use a **test double**—a placeholder type that stands in for another type to observe and verify behavior. This is similar to a “stunt double” in movies who performs tricky scenes in place of the actor. Among test doubles, **mock objects** are a special kind that **record interactions**, allowing you to assert that expected actions occurred.

Although Rust doesn’t support mock objects out of the box like some other languages, you can still build mock behavior by creating custom `struct`s that act like mock objects.

##### Example Use Case:
Imagine a library that monitors a numeric value approaching a maximum threshold (like an API rate limit). When a value gets close to the max, it should trigger messages (e.g. warnings). The library itself shouldn't send messages—it just needs a way to **tell** something else to send them.

To do this, the library defines a trait—say, `Messenger`—which the application using the library must implement. This keeps the message-sending logic flexible and decoupled.

```rust
pub trait Messenger {
    fn send(&self, msg: &str);
}

pub struct LimitTracker<'a, T: Messenger> {
    messenger: &'a T,
    value: usize,
    max: usize,
}

impl<'a, T> LimitTracker<'a, T>
where
    T: Messenger,
{
    pub fn new(messenger: &'a T, max: usize) -> LimitTracker<'a, T> {
        LimitTracker {
            messenger,
            value: 0,
            max,
        }
    }

    pub fn set_value(&mut self, value: usize) {
        self.value = value;

        let percentage_of_max = self.value as f64 / self.max as f64;

        if percentage_of_max >= 1.0 {
            self.messenger.send("Error: You are over your quota!");
        } else if percentage_of_max >= 0.9 {
            self.messenger
                .send("Urgent warning: You've used up over 90% of your quota!");
        } else if percentage_of_max >= 0.75 {
            self.messenger
                .send("Warning: You've used up over 75% of your quota!");
        }
    }
}
```

A good test that can mock the `Messenger` trait will be 

```rust

#[cfg(test)]
mod tests {
    use std::cell::RefCell;

    use super::*;

    struct MockMessenger {
        sent_messages: RefCell<Vec<String>>,
    }

    impl MockMessenger {
        fn new() -> MockMessenger {
            MockMessenger {
                sent_messages: RefCell::new(vec![]),
            }
        }
    }

    impl Messenger for MockMessenger {
        fn send(&self, message: &str) {
            self.sent_messages.borrow_mut().push(String::from(message));
        }
    }

    #[test]
    fn it_sends_an_over_75_percent_warning_message() {
        let mock_messenger = MockMessenger::new();
        let mut limit_tracker = LimitTracker::new(&mock_messenger, 100);

        limit_tracker.set_value(80);

        assert_eq!(mock_messenger.sent_messages.borrow().len(), 1);
    }
}
```

To see the reason why this is, look at https://doc.rust-lang.org/book/ch15-05-interior-mutability.html


#### Keeping Track of Borrows at Runtime with RefCell<T>
When creating references in Rust, we typically use `&` for immutable and `&mut` for mutable references. With `RefCell<T>`, we achieve similar behavior using its safe methods:  
- `.borrow()` gives us an immutable reference wrapped in `Ref<T>`  
- `.borrow_mut()` gives us a mutable reference wrapped in `RefMut<T>`  

Both `Ref<T>` and `RefMut<T>` implement the `Deref` trait, so they can be used like normal references.

`RefCell<T>` internally tracks how many active borrows exist:
- Every call to `.borrow()` increases the count of active immutable borrows.
- When the returned `Ref<T>` goes out of scope, the count is decreased.
- `.borrow_mut()` checks that no immutable or mutable borrows are active before allowing a mutable borrow.

The key difference is **timing of enforcement**:  
- With regular references, borrow rules are enforced **at compile time**.  
- With `RefCell<T>`, the rules are enforced **at runtime**.  

If you break the borrowing rules with `RefCell<T>`, the program will **panic** at runtime instead of failing to compile.

#### Having Multiple Owners of Mutable Data by Combining Rc<T> and RefCell<T>

A common and powerful pattern in Rust is combining `Rc<T>` and `RefCell<T>`. Each serves a different purpose:

- `Rc<T>` (Reference Counted): Allows **multiple ownership** of the same data, but only **immutable** access.
- `RefCell<T>`: Allows **interior mutability**, i.e., you can mutate data even through an immutable reference, but it enforces **borrowing rules at runtime** instead of compile time.

By combining them as `Rc<RefCell<T>>`, you get:
- **Multiple ownership** of the data (`Rc`)
- **Mutable access** to the data (`RefCell`), even when accessed via an `Rc`

This is useful in scenarios where shared, mutable state is needed in a single-threaded context—like in GUI frameworks, graph structures, or certain game engines.

-- Example

```rust
use std::rc::Rc;
use std::cell::RefCell;

#[derive(Debug)]
struct User {
    name: String,
    age: u32,
}

fn main() {
    // Create a shared, mutable User
    let user = Rc::new(RefCell::new(User {
        name: String::from("Alice"),
        age: 30,
    }));

    let user1 = Rc::clone(&user);
    let user2 = Rc::clone(&user);

    // Mutate via first reference
    user1.borrow_mut().age += 1;

    // Read via second reference
    println!("Updated user: {:?}", user2.borrow());
}
```

-- Output:
```
Updated user: User { name: "Alice", age: 31 }
```
