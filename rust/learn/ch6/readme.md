
Rust uses the prelude `Option<T>` enum to perform null checks.
In order to have a value that can possibly be `null`, you must explicitly opt in by making the type of that value `Option<T>`. Then, when you use that value, you are required to explicitly handle the case when the value is `null`. Everywhere that a value has a type that isn’t an `Option<T>`, you can safely assume that the value isn’t null.

### The match Control Flow Construct
Rust has an extremely powerful control flow construct called `match` that allows you to compare a value against a series of patterns and then execute code based on which pattern matches. Patterns can be made up of `literal values`, `variable names`, `wildcards`, and many other things.


### Concise Control Flow with `if let` and `let else`

The if let syntax lets you combine if and let into a less verbose way to handle values that match one pattern while ignoring the rest.

consider this program below

```rust
    let config_max = Some(3u8);
    match config_max {
        Some(max) => println!("The maximum is configured to be {max}"),
        _ => (),
    }

```

This program matches on an `Option<u8>` value in the config_max variable but only wants to execute code if the value is the Some variant.

Instead, we could write this in a shorter way using `if let`.

```rust
    let config_max = Some(3u8);
    if let Some(max) = config_max {
        println!("The maximum is configured to be {max}");
    }

```

### Staying on the “happy path” with `let else`

One common pattern is to perform some computation when a value is present and return a default value otherwise.

imagine this example

```rust
impl UsState {
    fn existed_in(&self, year: u16) -> bool {
        match self {
            UsState::Alabama => year >= 1819,
            UsState::Alaska => year >= 1959,
            // -- snip --
        }
    }
}


fn describe_state_quarter(coin: Coin) -> Option<String> {
    if let Coin::Quarter(state) = coin {
        if state.existed_in(1900) {
            Some(format!("{state:?} is pretty old, for America!"))
        } else {
            Some(format!("{state:?} is relatively new."))
        }
    } else {
        None
    }
}

// This can be made easier using this format


fn describe_state_quarter(coin: Coin) -> Option<String> {
    let state = if let Coin::Quarter(state) = coin {
        state
    } else {
        return None;
    };

    if state.existed_in(1900) {
        Some(format!("{state:?} is pretty old, for America!"))
    } else {
        Some(format!("{state:?} is relatively new."))
    }
}


```

This is a bit annoying to follow in its own way, though! One branch of the if let produces a value, and the other one returns from the function entirely.

Now let's take advantage of the rust's `let-else` statment structure. The `let-else` syntax takes a pattern on the left side and an expression on the right, very similar to `if let`, but it does not have an `if` branch, only an `else` branch.

```rust
fn describe_state_quarter(coin: Coin) -> Option<String> {
    let Coin::Quarter(state) = coin else {
        return None;
    };

    if state.existed_in(1900) {
        Some(format!("{state:?} is pretty old, for America!"))
    } else {
        Some(format!("{state:?} is relatively new."))
    }
}

```



