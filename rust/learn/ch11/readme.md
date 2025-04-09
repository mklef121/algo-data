
### How to Write Tests
Tests are Rust functions that verify that the non-test code is functioning in the expected manner. 
The bodies of test functions typically perform these three actions:

- Set up any needed data or state.
- Run the code you want to test.
- Assert that the results are what you expect.

At its simplest, a test in Rust is a function that’s annotated with the test attribute. 

> Attributes are metadata about pieces of Rust code; one example is the derive attribute I used earlier (`#[derive(Debug)]`) 

To change a function into a test function, add `#[test]` on the line before fn. 

Look at the test below

```rust
pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
```

Note the #[test] annotation on the `it_works` function: this attribute indicates this is a test function, so the test runner knows to treat this function as a test. 

We might also have non-test functions in the tests module to help set up common scenarios or perform common operations, so we always need to indicate which functions are tests.

The `#[cfg(test)]` attribute in the `test mod` is used to conditionally compile code only when running tests.
It’s most often used to exclude test modules inside your code files from production builds or binaries.

Use the `cargo test` command on the package root to run tests

The `assert!` macro, provided by the standard library, is useful when you want to ensure that some condition in a test evaluates to true. 
The `assert!` macro checks if a condition is `true`. If it is, the test passes silently; if not, it triggers a `panic!` and the test fails.

see an example of where it's used below

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn larger_can_hold_smaller() {
        let larger = Rectangle {
            width: 8,
            height: 7,
        };
        let smaller = Rectangle {
            width: 5,
            height: 1,
        };

        assert!(larger.can_hold(&smaller));
    }
}
```

`assert_eq!` and `assert_ne!` Macros are used to test for equality between the result of the code under test and the value you expect the code to return. 

> The `assert_eq!` and `assert_ne!` macros in Rust internally use `==` and `!=` for comparison. Upon failure, they print the compared values using debug formatting. This requires the types being compared to implement both the `PartialEq` trait (for equality checks) and the `Debug` trait (for printing). For custom structs and enums, you can easily enable these traits by adding `#[derive(PartialEq, Debug)]` to their definition.


#### Adding Custom Failure Messages
You can also add a custom message to be printed with the failure message as optional arguments to the `assert!`, `assert_eq!`, and `assert_ne!` macros.

Look at it in practice below

```rust
pub fn greeting(name: &str) -> String {
    format!("Hello {name}!")
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn greeting_contains_name() {
        let result = greeting("Carol");
        assert!(
            result.contains("Carol"),
            "Greeting did not contain name, value was `{result}`"
        );
    }
}
```

#### Checking for Panics with should_panic
In addition to checking return values, it’s important to check that our code handles error conditions as we expect, We do this by adding the attribute `should_panic` to our test function. The test passes if the code inside the function panics; the test fails if the code inside the function doesn’t panic.



Our tests so far all panic when they fail. We can also write tests that use `Result<T, E>!`. In the body of the function, rather than calling the `assert_eq!` macro, we return `Ok(())` when the test passes and an `Err`.

```rust
    #[test]
    fn it_works() -> Result<(), String> {
        let result = add(2, 2);

        if result == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
```

### Controlling How Tests Are Run

`cargo run` compiles and executes your code. Similarly, `cargo test` compiles your code in test mode into a binary and runs the generated binary. By default, this executable runs all tests in parallel and captures output(stdout and stderr) for cleaner results, but you can use command-line options to modify this behavior.

When using `cargo test`, some command-line options control Cargo itself, while others are passed to the compiled test executable. To differentiate them, use `--` as a separator: options before `--` go to `cargo test`, and options after `--` go to the test binary. Use `cargo test --help` and `cargo test -- --help` to see the respective options.

Look at this command for example

`cargo test --package ch11 --lib -- tests::it_works --exact --show-output`

The `--package ch11 --lib` flags are passed to cargo, while `--exact --show-output` are passed to the compiled test binary.

#### Running Tests in Parallel or Consecutively

By default, Rust runs tests in parallel using threads for faster feedback. However, this means tests must not share state or resources like files or environment variables, or they might interfere with each other and fail unexpectedly. For example, if multiple tests read/write the same file, they can conflict. To avoid this, use unique files per test or run tests serially. You can control thread usage with the `--test-threads` flag.

`$ cargo test -- --test-threads=1`
We set the number of test threads to 1, telling the program not to use any parallelism. 

#### Showing Function Output
By default, Rust's test runner hides anything printed to standard output like `println!` output from passing tests, showing only the "passed" status. However, if a test fails, any `println!` output from that test will be displayed along with the failure details.

We can use the `cargo test -- --show-output` If we want to see printed values for passing tests as well.

We can ise `cargo test add`, `cargo test one_hundred` to run test functions with specifica names say `one_hundred` and also test functions that  their names contain `add`

we can use the `#[ignore]` attribute to annotate a test that should not be run


### Test Organization

Rust categorizes tests into two main types:

* **Unit Tests:** Small, focused tests that isolate and test individual modules, including private interfaces.
* **Integration Tests:** External tests that use only the library's public API, potentially testing interactions between multiple modules.


You’ll put unit tests in the `src` directory in each file with the code that they’re testing. The convention is to create a `module` named tests in each file to contain the test functions and to annotate the module with `#[cfg(test)]`.



#### Integration Tests

Rust integration tests are external to your library, using only its public API, just like any other consuming code. They verify that different parts of the library work correctly together. To write them, you create a `tests` directory at the top level of your project.

```rust
// Filename: src/lib.rs

pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

// Filename: tests/integration_test.rs
// ch11 is the crate name found in Cargo.toml under [package] name
use ch11::add_two;

#[test]
fn it_adds_two() {
    let result = add(2, 2);
    assert_eq!(result, 4);
}

```
