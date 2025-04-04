#[derive(Debug)]
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}

// Using Tuple Structs Without Named Fields to Create Different Types
/*
Rust also supports structs that look similar to tuples, called tuple structs.
Tuple structs have the added meaning the struct name provides but don’t have names associated with their fields;
rather, they just have the types of the fields.
 */

struct Color(i32, i32, i32, f64);
struct Point(i32, i32, i32);

struct QuitMessage; // unit struct
struct MoveMessage {
    x: i32,
    y: i32,
}
struct WriteMessage(String); // tuple struct
struct ChangeColorMessage(i32, i32, i32); // tuple struct

fn main() {
    let mut user1 = User {
        active: true,
        username: String::from("someusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };
    user1.email = String::from("anotheremail@example.com");

    // Using struct update syntax,
    let user2 = User {
        email: String::from("another@example.com"),
        // The ..user1 must come last to specify that any remaining fields should get
        // their values from the corresponding fields in user1
        ..user1
    };

    let us2 = build_user(user2.email, user2.username);
    println!("First Print USER -> {:?}", us2);

    // build_user_short(user1.email, user1.username);
    let _black = Color(0, 0, 0, 3.56);
    let origin = Point(0, 0, 0);
    // desctructuring tuple structs
    let Point(_x, _y, _z) = origin;

    println!("Hello, world! USER -> {:#?}", us2);

    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!(
        "The area of the rectangle is {} square pixels.",
        area_struct(&rect1)
    );

    dbg!(&rect1);

    println!("using rectangle method area {}", rect1.area());

    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));

    let rect4 = Rectangle::square(3);

    let val = WriteMessage(String::from("stand tall"));

    // the value of a tuple struct can also be accessed this way
    let _he = val.0;

    println!(
        "Show Rect created from static associated functions {:#?}",
        rect4
    );
}

fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username: username,
        email: email,
        sign_in_count: 1,
    }
}

// Using the Field Init Shorthand
fn build_user_short(email: String, username: String) -> User {
    User {
        active: true,
        username,
        email,
        sign_in_count: 1,
    }
}

fn area(width: u32, height: u32) -> u32 {
    width * height
}

fn area_tuples(dimensions: (u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}

// #[derive] is an `outer attribute` used to automatically implement common traits for a struct or enum.
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn area_struct(rectangle: &Rectangle) -> u32 {
    rectangle.width * rectangle.height
}

// Method Syntax

impl Rectangle {
    fn area(self: &Self) -> u32 {
        return self.width * self.height;
    }

    // any of this methods of accessing the instance will work
    // can also be `&mut self`` and  `self``
    fn sum(&self) -> u32 {
        return self.width + self.height;
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }

    // constructors associated function
    fn square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    }
}
