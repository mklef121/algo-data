
### Packages and Crates


#### Crate

A `crate` is the smallest amount of code that the Rust compiler considers at a time. Even if you run `rustc` rather than `cargo` and pass a single source code file, the compiler considers that file to be a crate.
Crates can contain `modules`, and the `modules` may be defined in other files that get compiled with the crate.

A crate can come in one of two forms: 
- **a binary crate:** Binary crates are programs you can compile to an executable that you can run, such as a command-line program or a server. Each must have a function called main that defines what happens when the executable runs
- **a library crate:** Library crates don’t have a main function, and they don’t compile to an executable. Instead, they define functionality intended to be shared with multiple projects.

The **crate root** is a source file that the Rust compiler starts from and makes up the root module of your crate. This is usually the `src/main.rs` file.


#### Package

A `package` is a bundle of one or more `crates` that provides a set of functionality. A `package` contains a `Cargo.toml` file that describes how to build those crates.

For example **Cargo** is actually a `package` that contains the `binary crate` for the command-line tool you’ve been using to build your code. The Cargo package also contains a library crate that the binary crate depends on. Other projects can depend on the Cargo library crate to use the same logic the Cargo command-line tool uses.

A package can contain as many binary crates as you like, but at most only one library crate. A package must contain at least one crate, whether that’s a library or binary crate.

Cargo knows that if the package directory contains `src/lib.rs`, the package contains a library crate with the same name as the package, and `src/lib.rs` is its crate root.
If a package contains `src/main.rs` and `src/lib.rs`, it has two crates: a **binary** and a **library**, both with the same name as the package. A **package** can have multiple binary crates by placing files in the `src/bin` directory: each file will be a separate binary crate.

#### Modules Cheat Sheet

- **Start from the crate root**: When compiling a crate, the compiler first looks in the crate root file (usually src/lib.rs for a library crate or src/main.rs for a binary crate) for code to compile.
- **Declaring modules:** In the crate root file, you can declare new modules; say you declare a “garden” module with `mod garden;`. The compiler will look for the module’s code in these places:
    - Inline, within curly brackets that replace the semicolon following `mod garden`
    - In the file src/garden.rs
    - In the file src/garden/mod.rs
- **Private vs. public:** Code within a module is private from its parent modules by default. To make a module public, declare it with `pub mod` instead of `mod`. To make items within a public module public as well, use `pub` before their declarations.
- **The `use` keyword:** Within a scope, the use keyword creates shortcuts to items to reduce repetition of long paths. In any scope that can refer to `crate::garden::vegetables::Asparagus`, you can create a shortcut with use `crate::garden::vegetables::Asparagus;` and from then on you only need to write `Asparagus` to make use of that type in the scope.

#### Grouping Related Code in Modules
`Modules` let us organize code within a crate for readability and easy reuse. Modules also allow us to control the privacy of items because code within a module is private by default. 
Private items are internal implementation details not available for outside use. We can choose to make modules and the items within them public, which exposes them to allow external code to use and depend on them.

Take a look at the code below

```rust
mod front_of_house {
    mod hosting {
        fn add_to_waitlist() {}

        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}

        fn serve_order() {}

        fn take_payment() {}
    }
}
```

The body of the module then goes inside curly brackets. Inside modules, we can place other modules, as in this case with the modules `hosting` and `serving`. Modules can also hold definitions for other items, such as `structs`, `enums`, `constants`, `traits`, and `functions`.

By using modules, we can group related definitions together and name why they’re related. 

Earlier, we mentioned that `src/main.rs` and `src/lib.rs` are called `crate roots`. The reason for their name is that the contents of either of these two files form a module named `crate` at the root of the crate’s module structure, known as the **module tree**.

#### Paths for Referring to an Item in the Module Tree
To show Rust where to find an item in a module tree, we use a **path**.
A path can take two forms:
- *Absolute paths* in Rust start from the crate's root. For external crates, the path begins with the crate's name. For the current crate, it starts with the keyword `crate`. e.g `crate::front_of_house::hosting::add_to_waitlist();`
- A *relative path* starts from the current module and uses self, super, or an identifier in the current module. e.g `front_of_house::hosting::add_to_waitlist();`

If you want to make an item like a function or struct private, you put it in a module.
