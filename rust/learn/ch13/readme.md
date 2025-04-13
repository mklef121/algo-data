Rust closures automatically implement `FnOnce`, `FnMut`, and `Fn` traits based on how they use captured environment variables.

* **`FnOnce`**: All closures implement this (can be called at least once). Closures that *move* captured values out of their body only implement `FnOnce` because they can only be called once (ownership is transferred).
* **`FnMut`**: Closures that don't move captured values but *do mutate* them implement `FnMut` and can be called multiple times.
* **`Fn`**: Closures that neither move nor mutate captured values (or capture nothing) implement `Fn`. They can be called multiple times concurrently without affecting their environment.

These traits allow functions and structs to specify the types of closures they can accept as arguments.


## Processing a Series of Items with Iterators
The iterator pattern lets you process a sequence of items one by one without manually handling the iteration logic. In Rust, iterators are `lazy`: they do nothing until a method that consumes them is called.

```rust
let v1 = vec![1, 2, 3];

let v1_iter = v1.iter();

for val in v1_iter {
    println!("Got: {val}");
}
```

In the example above, we separate the creation of the iterator from the use of the iterator in the for loop. When the for loop is called using the iterator in `v1_iter`, each element in the iterator is used in one iteration of the loop, which prints out each value.

In languages without built-in iterators, you'd manually track indices in a loop to access elements—an error-prone and repetitive process. Rust’s iterators simplify this by handling the iteration logic for you, offering reusable, flexible ways to work with various data structures, not just indexable ones like vectors.

### The Iterator Trait and the next Method
All iterators implement a trait named `Iterator` that is defined in the standard library. The definition of the trait looks like this:

```rust
pub trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self::Item>;

    // methods with default implementations elided
}
```

Notice this definition uses some new syntax: `type Item` and `Self::Item`, which are defining an associated type with this trait. 

We can call the next method on iterators directly as shown below

```rust
    #[test]
    fn iterator_demonstration() {
        let v1 = vec![1, 2, 3];

        let mut v1_iter = v1.iter();

        assert_eq!(v1_iter.next(), Some(&1));
        assert_eq!(v1_iter.next(), Some(&2));
        assert_eq!(v1_iter.next(), Some(&3));
        assert_eq!(v1_iter.next(), None);
    }
```

Each call to next eats up an item from the iterator. We didn’t need to make v1_iter mutable when we used a for loop because the loop took ownership of v1_iter and made it mutable behind the scenes.

Rust offers three ways to create iterators from a collection:

* `.iter()`: Creates an iterator yielding **immutable** references (`&T`) to the items.
* `.into_iter()`: Creates an iterator that takes ownership of the collection and yields owned values (`T`).
* `.iter_mut()`: Creates an iterator yielding mutable references (`&mut T`) to the items, allowing in-place modification.

### Methods that Consume the Iterator
The `Iterator` trait in Rust provides many useful methods with default implementations, often relying on the fundamental `next` method that you must implement. Methods that use `next` are called **consuming adapters** because they exhaust the iterator by calling `next` repeatedly until it returns `None`. An example is the `sum` method, which takes ownership of the iterator, consumes it by calling `next` on each item, and returns the final sum.

```rust
    #[test]
    fn iterator_sum() {
        let v1 = vec![1, 2, 3];

        let v1_iter = v1.iter();

        let total: i32 = v1_iter.sum();

        assert_eq!(total, 6);
    }
```


### Methods that Produce Other Iterators

Iterator adapters are methods defined on the `Iterator` trait that transform an iterator into a new iterator without consuming the original. They allow you to chain operations and modify how elements are produced. Common examples include:

* **`.map(f)`:** Applies a function `f` to each item yielded by the iterator, producing a new iterator with the results.

    ```rust
    let numbers = vec![1, 2, 3];
    let doubled = numbers.iter().map(|x| x * 2); // Creates an iterator yielding 2, 4, 6
    ```

* **`.filter(predicate)`:** Creates a new iterator that yields only the items from the original iterator for which the `predicate` function returns `true`.

    ```rust
    let numbers = vec![1, 2, 3, 4];
    let even = numbers.iter().filter(|x| x % 2 == 0); // Creates an iterator yielding 2, 4
    ```

* **`.take(n)`:** Creates a new iterator that yields at most `n` items from the original iterator.

    ```rust
    let numbers = vec![1, 2, 3, 4, 5];
    let first_three = numbers.iter().take(3); // Creates an iterator yielding 1, 2, 3
    ```

* **`.skip(n)`:** Creates a new iterator that skips the first `n` items of the original iterator and then yields the remaining items.

    ```rust
    let numbers = vec![1, 2, 3, 4, 5];
    let after_first_two = numbers.iter().skip(2); // Creates an iterator yielding 3, 4, 5
    ```

Iterator adapters are lazy, meaning they don't do any work until a consuming adapter (like `.collect()`, `.sum()`, `.for_each()`) is called on the resulting iterator. This allows for efficient, chained operations on sequences.

