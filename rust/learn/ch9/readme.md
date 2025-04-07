### Recoverable Errors with Result

The `Result` enum is defined as having two variants, `Ok` and `Err`, as follows:

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

take a look at the code below

```rust
    let greeting_file_result: Result<File, io::Error> = File::open("hello.txt");
```

The return type of File::open is a `Result<File, std::io::Error>`. This return type means the call to File::open might succeed and return a file handle that we can read from or write to. The function call also might fail.

We can treat the error using the `match` operator.

```rust
    let greeting_file = match greeting_file_result {
        Ok(file) => file,
        Err(error) => panic!("Problem opening the file: {error:?}"),
    };
```

### Matching on Different Errors

The code above panics on any `File::open` failure. But ideally, we want to handle errors differently: if the file doesn’t exist, create it and return the handle; if the error is something else (like permission issues), then panic.

```rust
use std::fs::File;
use std::io::ErrorKind;

fn main() {
    let greeting_file_result = File::open("hello.txt");

    let greeting_file = match greeting_file_result {
        Ok(file) => file,
        Err(error) => match error.kind() {
            ErrorKind::NotFound => match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(e) => panic!("Problem creating the file: {e:?}"),
            },
            other_error => {
                panic!("Problem opening the file: {other_error:?}");
            }
        },
    };
}
```


### Propagating Errors
When a function’s implementation calls something that might fail, instead of handling the error within the function itself you can return the error to the calling code so that it can decide what to do. This is known as **propagating the error**.


