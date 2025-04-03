const _THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;

fn main() {
    let mut y = 5;
    println!("The value of y is: {y}");
    y = 6;
    println!("The value of y is: {y}");

    /*
        Shadowing

        you can declare a new variable with the same name as a previous variable.
        We can shadow a variable by using the same variable’s name and repeating the use of the let keyword as follows
    */
    let x: i32 = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");

    // the :u32 type used here gives a hint to the parse function on  what type to convert the "42" string into.
    let _guess: u32 = "42".parse().expect("Not a number!");

    // this is another expression of same code
    let _guess = "42".parse::<u32>().expect("Not a number!");

    /*
    Data Types
    Every value in Rust is of a certain data type, which tells Rust what kind of data
    is being specified so it knows how to work with that data.
    e.g: let x: i32 = 5; // here i32 is the data type

    Data types have Scalar Types and compound types

    A scalar type represents a single value. Rust has four primary scalar
    types: integers, floating-point numbers, Booleans, and characters.

    */

    let _sid: isize = 0;

    // floating point
    let _x = 2.0; // f64

    let _y: f32 = 3.0; // f32

    arithmetic();
    functions();
    control_flow();
    iteration_repetition();
}

fn arithmetic() {
    // addition
    let _sum = 5 + 10;

    // subtraction
    let _difference = 95.5 - 4.3;

    // multiplication
    let _product = 4 * 30;

    // division
    let _quotient = 56.7 / 32.2;

    // Integer division truncates toward zero to the nearest integer.
    let _truncated = -5 / 3; // Results in -1

    // remainder
    let _remainder = 43 % 5;

    // floating point numbers
    let _x = 2.0; // f64

    let _y: f32 = 3.0; // f32

    // boolean types
    let _t = true;

    let _f: bool = false; // with explicit type annotation

    // The Character Type.  we specify char literals with single quotes,
    // Rust’s char type is four bytes in size and represents a Unicode Scalar Value,
    // which means it can represent a lot more than just ASCII.
    let _c = 'z';
    let _z: char = 'ℤ'; // with explicit type annotation
    let _heart_eyed_cat = '😻';

    // Compound Types

    // Compound types can group multiple values into one type. Rust has two primitive compound types: tuples and arrays.

    // A tuple is a general way of grouping together a number of values with a variety of types into one compound type.
    // Tuples have a fixed length: once declared, they cannot grow or shrink in size.
    let _tup: (i32, f64, u8) = (500, 6.4, 1);
    let tup = (500, 6.4, 1);

    // destructuring
    let (x, y, z) = tup;
    println!("The value of y is: {y}, The value of x is: {x}, The value of z is: {z}, ");

    // We can also access a tuple element directly by using a period (.)
    // followed by the index of the value we want to access. For example:

    let x: (i32, f64, u8) = (500, 6.4, 1);

    let _five_hundred = x.0;

    let _six_point_four = x.1;

    let _one = x.2;

    // The Array Type

    // arrays in Rust have a fixed length.
    let _a: [i32; 5] = [1, 2, 3, 4, 5];

    let _months = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    ];

    // this is equivalent to let a = [3, 3, 3, 3, 3];
    let _a = [3; 5];

    let a = [1, 2, 3, 4, 5];

    let _first = a[0];
    let _second = a[1];
}

fn functions() {
    // Rust code uses snake case as the conventional style for function and variable names,
    // in which all letters are lowercase and underscores separate words.
    let y = {
        let x = 3;
        x + 1
    };

    // y will be 4 here, even though i don't understand why !!!
    println!("The value of the scoped y is: {y}");

    another_function(5, 'g');

    println!("The five value is : {}, {}", five(), five_return());
}

fn another_function(x: i32, unit_label: char) {
    println!("The value of x is: {x}, \n The measurement is: {unit_label}");
}

// You can return early from a function by using the return keyword and specifying a value,
// but most functions return the last expression implicitly.
fn five() -> i32 {
    5
}

// similar to the function above
fn five_return() -> i32 {
    return 5;
}

// Control Flow
fn control_flow() {
    let number = 6;

    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }

    // Using if in a let Statement

    let condition = true;
    let number = if condition { 5 } else { 6 };

    println!("The value of number is: {number}");
}

// Repetition with Loops
// Rust has three kinds of loops: loop, while, and for. Let’s try each one.
fn iteration_repetition() {
    /*
    The loop keyword tells Rust to execute a block of code over and over again forever or until
    you explicitly tell it to stop.
     loop {
        println!("again!");
    }
    */

    let mut counter = 0;

    // To pass a loop's result to your code, add the return value after the `break` expression.
    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The loop result is {result}");

    // Loop Labels to Disambiguate Between Multiple Loops

    // Break and continue affect the innermost loop by default, but you can use a label (starting with a single quote)
    // with them to target a specific outer loop instead.

    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");

    // Conditional Loops with while
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1;
    }
    println!("LIFTOFF!!!");

    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {element}");
    }

    // same as while loop above
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
