## Generic Types, Traits, and Lifetimes

### Generic Types
Every programming language has tools for effectively handling the duplication of concepts. In Rust, one such tool is `generics`.

Functions can take parameters of some generic type, instead of a concrete type like i32 or String, in the same way they take parameters with unknown values to run the same code on multiple concrete values.

We use generics to create definitions for items like function signatures or structs, which we can then use with many different concrete data types. 

```rust
fn largest_i32(list: &[i32]) -> &i32 {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn largest_char(list: &[char]) -> &char {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn main() {
    let number_list = vec![34, 50, 25, 100, 65];

    let result = largest_i32(&number_list);
    println!("The largest number is {result}");

    let char_list = vec!['y', 'm', 'a', 'q'];

    let result = largest_char(&char_list);
    println!("The largest char is {result}");
}

```

The two functions above are repeating a check we can do on different data types

let's use a generic that looks like `fn largest<T>(list: &[T]) -> &T {` to simplify this

```rust
fn largest<T: std::cmp::PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}
```

With the implementation above, we can see that one function serves all the comparison purposes

#### In Struct Definitions
We can also define structs to use a generic type parameter in one or more fields using the `<>` syntax.

```rust
struct Point<T> {
    x: T,
    y: T,
}

fn main() {
    let integer = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };
}
```

```rust
// In Enum Definitions
enum Option<T> {
    Some(T),
    None,
}


enum Result<T, E> {
    Ok(T),
    Err(E),
}

// In Method Definitions

struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}


// We can also specify constraints on generic types when defining methods on the type.
impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

// Generic type parameters in a struct definition aren’t always the same as those you
// use in that same struct’s method signatures.
struct Point<X1, Y1> {
    x: X1,
    y: Y1,
}

impl<X1, Y1> Point<X1, Y1> {
    fn mixup<X2, Y2>(self, other: Point<X2, Y2>) -> Point<X1, Y2> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

```

#### Performance of Code Using Generics

You might be wondering whether there is a runtime cost when using generic type parameters. 
The good news is that using generic types won’t make your program run any slower than it would with concrete types.

Rust accomplishes this by performing **monomorphization** of the code using generics at compile time. 

**Monomorphization** is the process of turning generic code into specific code by filling in the concrete types that are used when compiled. In this process, the compiler looks at all the places where generic code is called and generates code for the concrete types the generic code is called with.
Look at the generic code below

```rust
enum Option<T> {
    Some(T),
    None,
}

let integer = Some(5);
let float = Some(5.0);
```

The monomorphized version of the code looks similar to the following

```rust
enum Option_i32 {
    Some(i32),
    None,
}

enum Option_f64 {
    Some(f64),
    None,
}

fn main() {
    let integer = Option_i32::Some(5);
    let float = Option_f64::Some(5.0);
}
```

The generic `Option<T>` is replaced with the specific definitions created by the compiler. The process of monomorphization makes Rust’s generics extremely efficient at runtime.

### Traits: Defining Shared Behavior
A `trait` defines the functionality a particular type has and can share with other types.
We can use `traits` to define shared behavior in an abstract way. 
We can use `trait bounds` to specify that a generic type can be any type that has certain behavior.

#### Defining a Trait
A type’s behavior consists of the methods we can call on that type. Different types share the same behavior if we can call the same methods on all of those types.
Trait definitions are a way to group method signatures together to define a set of behaviors necessary to accomplish some purpose.

```rust
// trait definition
pub trait Summary {
    fn summarize(&self) -> String;
}
```

Each type implementing this trait must provide its own custom behavior for the body of the method. The compiler will enforce that any type that has the Summary trait will have the method summarize defined with this signature exactly.

> Traits are similar to a feature often called `interfaces` in other languages, although with some differences.

Now that we’ve defined the desired signatures of the Summary trait’s methods, we can implement it on the types in our media aggregator thus.

```rust
pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
    pub content: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("{}, by {} ({})", self.headline, self.author, self.location)
    }
}

pub struct Tweet {
    pub username: String,
    pub content: String,
    pub reply: bool,
    pub retweet: bool,
}

impl Summary for Tweet {
    fn summarize(&self) -> String {
        format!("{}: {}", self.username, self.content)
    }
}
```

Here are some trait rules

✅ Allowed: Trait or type is local

```rust
// In your crate
pub trait Summary {
    fn summarize(&self) -> String;
}

// This is fine! You’re implementing Summary (your trait) on Vec<T> (external type)
impl<T> Summary for Vec<T> {
    fn summarize(&self) -> String {
        format!("Vec with {} elements", self.len())
    }
}

// In your crate
pub struct Tweet {
    pub username: String,
    pub content: String,
}

// You’re implementing Display (std trait) on Tweet (your type)
impl std::fmt::Display for Tweet {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{}: {}", self.username, self.content)
    }
}

```

❌ Not allowed: Both the trait and the type are from outside

```rust
// std::fmt::Display is from the standard library
// Vec<T> is also from the standard library

impl std::fmt::Display for Vec<i32> {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "This is a vector!")
    }
}
```
#### Default Implementations in Traits

Sometimes it’s useful to have default behavior for some or all of the methods in a trait instead of requiring implementations for all methods on every type. 

```rust
pub trait Summary {
    fn summarize(&self) -> String {
        String::from("(Read more...)")
    }
}
// this empty block works since summarize method is not required
impl Summary for NewsArticle {}
```

Once a type implements this Trait, then it's not required to provide an implementation.

```rust
// Default implementations can call other methods in the same trait, 
// even if those other methods don’t have a default implementation.

pub trait Summary {
    fn summarize_author(&self) -> String;

    fn summarize(&self) -> String {
        format!("(Read more from {}...)", self.summarize_author())
    }
}
```

#### Traits as Parameters
Let us explore how to use traits to define functions that accept many different types.

```rust
// Instead of a concrete type for the item parameter, we specify the impl keyword and the trait name.
// This parameter accepts any type that implements the specified trait. 
pub fn notify(item: &impl Summary) {
    println!("Breaking news! {}", item.summarize());
}
```

#### Trait Bound Syntax

The `impl Trait` syntax used above works for straightforward cases but is actually syntax sugar for a longer form known as a trait bound; it looks like this:

```rust
pub fn notify<T: Summary>(item: &T) {
    println!("Breaking news! {}", item.summarize());
}

// Specifying Multiple Trait Bounds with the + Syntax
// We can also specify more than one trait bound. 
// Say we wanted notify to use display formatting as well as summarize on item:
pub fn notify(item: &(impl Summary + Display)){}
pub fn notify<T: Summary + Display>(item: &T) {}

```

#### Clearer Trait Bounds with where Clauses

So imagine this function definition with trait bounds below
```rust
fn some_function<T: Display + Clone, U: Clone + Debug>(t: &T, u: &U) -> i32
```

The function signature starts getting hard to read. For this reason, Rust has alternate syntax for specifying trait bounds inside a `where` clause after the function 

```rust
fn some_function<T, U>(t: &T, u: &U) -> i32
where
    T: Display + Clone
    U: Clone + Debug
{

}
```
This function’s signature is less cluttered.

```rust
// Returning Types That Implement Traits
fn returns_summarizable() -> impl Summary {
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from(
            "of course, as you probably already know, people",
        ),
        reply: false,
        retweet: false,
    }
}
```
#### Using Trait Bounds to Conditionally Implement Methods

Recall that we can use Generics to conditionally implement methods as below

```rust
use std::fmt::Display;

struct Pair<T> {
    x: T,
    y: T,
}

// this new function works for every generic type of pair
impl<T> Pair<T> {
    fn new(x: T, y: T) -> Self {
        Self { x, y }
    }
}

// this cmp_display will only worl of the Generic type of pair implements Display and PartialOrd 
impl<T: Display + PartialOrd> Pair<T> {
    fn cmp_display(&self) {
        if self.x >= self.y {
            println!("The largest member is x = {}", self.x);
        } else {
            println!("The largest member is y = {}", self.y);
        }
    }
}
```

In same manner, We can also conditionally implement a trait for any type that implements another trait.

```rust
impl<T: Display> ToString for T {
    // --snip--
}
```

Implementations of a trait on any type that satisfies the trait bounds are called **blanket implementations**

### Validating References with Lifetimes
**Lifetimes** are another kind of generic, they ensure that references are valid as long as we need them to be.
Every **reference** in Rust has a lifetime, which is the scope for which that reference is valid.

> The main aim of lifetimes is to prevent dangling references, which cause a program to reference data other than the data it’s intended to reference. 

```rust
fn main() {
    let r;                // ---------+-- 'a
                          //          |
    {                     //          |
        let x = 5;        // -+-- 'b  |
        r = &x;           //  |       |
    }                     // -+       |
                          //          |
    println!("r: {r}");   //          |
}                         // ---------+
```

Rust disallows this code above because `x` goes out of scope at the end of the inner block, while `r` still exists in the outer scope. If allowed, `r` would point to invalid/deallocated memory. Rust’s **borrow checker** catches this by ensuring references don’t outlive the data they point to.










