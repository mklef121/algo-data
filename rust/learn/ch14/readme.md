### Customizing Builds with Release Profiles

In Rust, **release profiles** are configurable settings in Cargo that control how your code is compiled. The two main profiles are:

- **`dev`** – used with `cargo build`, optimized for faster compile times during development.
- **`release`** – used with `cargo build --release`, optimized for performance in production.

Each profile can be customized independently in `Cargo.toml` to suit specific needs.

Cargo uses default settings for each build profile unless you override them in your `Cargo.toml` file using `[profile.*]` sections. This lets you customize specific options, like `opt-level`, without redefining the entire profile. For example, the `dev` profile uses lower optimization for faster builds, while `release` enables higher optimization for better performance.

```toml
[profile.dev]
opt-level = 0

[profile.release]
opt-level = 3
```

### Making Useful Documentation Comments
Rust also has a particular kind of comment for documentation, known conveniently as a **documentation comment**, that will generate HTML documentation. The HTML displays the contents of documentation comments for public API items intended for programmers interested in knowing how to use your crate as opposed to how your crate is implemented.
Documentation comments use three slashes, `///`, instead of two and support Markdown notation for formatting the text. Place documentation comments just before the item they’re documenting. 

```rust
/// Adds one to the number given.
///
/// # Examples
///
/// ```
/// let arg = 5;
/// let answer = my_crate::add_one(arg);
///
/// assert_eq!(6, answer);
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}
```

Running `cargo doc --open` generates HTML documentation for your crate and its dependencies, then automatically opens it in your web browser for easy viewing.


### Commenting Contained Items
Doc comments starting with `//!` document the *enclosing* item (crate or module), not the code that follows. They are typically used at the beginning of `src/lib.rs` to document the crate or within a module file to document that module. For example, `//! My awesome crate` at the top of `src/lib.rs` describes the entire crate.

```rust
//! # My Crate
//!
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.

/// Adds one to the number given.
pub fn add_one(x: i32) -> i32
```

to read more about publishing a crate, read on https://doc.rust-lang.org/book/ch14-02-publishing-to-crates-io.html


### Cargo Workspaces
Cargo offers a feature called workspaces that can help manage multiple related packages that are developed in tandem. As your project develops, you might find that the library crate continues to get bigger and you want to split your package further into multiple library crates. 

A workspace in Rust is a collection of packages that share the same `Cargo.lock` and output directory. It helps organize related crates within a single project. For example, you can create a workspace with one binary crate (the main app) and two library crates (e.g., `add_one` and `add_two`), where the binary depends on both libraries. This setup keeps the project modular and maintainable.

For entire doucmentation on workspaces look at https://doc.rust-lang.org/book/ch14-03-cargo-workspaces.html

also look up the  [workspace folder](../add-workspace/)