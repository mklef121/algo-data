use rand::Rng;
use std::{cmp::Ordering, io};

fn main() {
    println!("Guess the number game!");

    loop {
        println!("Please input your guess.");

        let secret_number = rand::rng().random_range(1..=100);

        println!("The secret number is: {secret_number}");

        // In Rust, variables are immutable by default, meaning once we give the variable a value, the value won’t change.
        // mut ensures that the values can change
        let mut guess = String::new();
        // The :: syntax in the ::new line indicates that new is an associated function of the String type

        let _apples = 5; // immutable
        io::stdin()
            .read_line(&mut guess) // The & indicates that this argument is a reference or address
            .expect("Failed to read line");

        // Rust allows us to shadow the previous value of guess with a new one. Shadowing lets us reuse
        // the guess variable name rather than forcing us to create two unique variables,
        let guess: u32 = match guess.trim().parse() {
            Ok(numb) => numb,
            Err(_) => {
                println!("kindly enter a valid number");
                continue;
            }
        };

        println!("You guessed: {}", guess);
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                return;
            }
        }
    }
}
