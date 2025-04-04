fn main() {
    let _four = IpAddrKind::V4;
    let _six = IpAddrKind::V6;

    let _home = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };

    let _loopback = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };

    let home: IpAddr2 = IpAddr2::V4(127, 0, 0, 1);

    let loopback = IpAddr2::V6(String::from("::1"));

    println!("Hello, world! {:#?} and {:#?}", home, loopback);

    let msg1 = Message::Quit;
    let msg2 = Message::Move { x: 10, y: 20 };
    let msg3 = Message::Write(String::from("Hello, Rust!"));
    let msg4 = Message::ChangeColor(255, 0, 0);

    process_message(msg1);
    process_message(msg2);
    process_message(msg3);
    process_message(msg4);

    // using the enum implementation method
    let msg = Message::Write(String::from("Rust enums are powerful!"));
    msg.display();

    let msg = Message::new_string_message(String::from("hello boys"));
    msg.display();

    let five = Some(5);
    let _six = plus_one(five);
    let _none = plus_one(None);
}

// Defining an Enum

enum IpAddrKind {
    V4,
    V6,
}

struct IpAddr {
    kind: IpAddrKind,
    address: String,
}

// we can put data directly into each enum variant.
#[derive(Debug)]
enum IpAddr2 {
    V4(u8, u8, u8, u8),
    V6(String),
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// just as we’re able to define methods on structs using `impl`, we’re also able to define methods on `enums``.
impl Message {
    fn display(&self) {
        match self {
            Message::Quit => println!("The program will quit."),
            Message::Move { x, y } => println!("Moving to ({}, {})", x, y),
            Message::Write(text) => println!("Message content: {}", text),
            Message::ChangeColor(r, g, b) => println!("New color set to RGB({}, {}, {})", r, g, b),
        }
    }

    fn new_string_message(val: String) -> Self {
        return Self::Write(val);
    }
}

// we can use a match on an enum
fn process_message(msg: Message) {
    match msg {
        Message::Quit => println!("Received Quit message"),
        Message::Move { x, y } => println!("Moving to position: ({}, {})", x, y),
        Message::Write(text) => println!("Writing message: {}", text),
        Message::ChangeColor(r, g, b) => println!("Changing color to RGB({}, {}, {})", r, g, b),
    }
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}

// using if let to reduce the match block and just choose one methc one enum value
fn value_in_cents_if(coin: Coin) -> u8 {
    let mut count: u8 = 0;
    if let Coin::Quarter = coin {
        println!("State quarter ");
    } else {
        count += 1;
    }

    count
}

// matching the option type
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

// matching the option type
// this is the samething as the one above, I just used the Enum with it's fields instead of just the Prelude
// from rust
fn plus_one_plus(x: Option<i32>) -> Option<i32> {
    match x {
        Option::None => None,
        Option::Some(i) => Some(i + 1),
    }
}

fn play_dice(dice_roll: u8) {
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        //last arm that covers every other possible value
        other => move_player(other),
    }
}

fn play_dice_2(dice_roll: u8) {
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        // _ is a special pattern that matches any value and does not bind to that value.
        _ => reroll(),
    }
}

fn reroll() {}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
fn move_player(num_spaces: u8) {}
