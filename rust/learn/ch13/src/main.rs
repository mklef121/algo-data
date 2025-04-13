use std::{thread, time::Duration};

fn main() {
    let store = Inventory {
        shirts: vec![ShirtColor::Blue, ShirtColor::Red, ShirtColor::Blue],
    };

    let user_pref1 = Some(ShirtColor::Red);
    let giveaway1 = store.giveaway(user_pref1);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref1, giveaway1
    );

    let user_pref2 = None;
    let giveaway2 = store.giveaway(user_pref2);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref2, giveaway2
    );

    let expensive_closure = |num: u32| -> u32 {
        println!("calculating slowly...");
        thread::sleep(Duration::from_secs(0));
        println!("done sleeping ...");
        num
    };

    expensive_closure(34);

    let mut list = vec![1, 2, 3];
    println!("Before defining closure: {list:?}");

    // this closure mutably borrows the variable list
    // this is because println takes the pointer to variable alone and does not take over ownership
    let only_borrows = || {
        println!("From closure: {list:?}");
    };

    println!("Before calling closure: {list:?}");
    only_borrows();
    println!("After calling closure: {list:?}");

    let mut borrows_mutably = || list.push(7);

    borrows_mutably();
    println!("After calling borrows_mutably closure: {list:?}");

    /*
    If you want to force the closure to take ownership of the values it uses in the environment
    even though the body of the closure doesn’t strictly need ownership, you can use the `move`
    keyword before the parameter list.
    */

    let list = vec![1, 2, 3];
    println!("Before defining closure for new thread: {list:?}");

    thread::spawn(move || println!("From thread: {list:?}"))
        .join()
        .unwrap();

    let mut list = [
        Rectangle {
            width: 10,
            height: 1,
        },
        Rectangle {
            width: 3,
            height: 5,
        },
        Rectangle {
            width: 7,
            height: 12,
        },
    ];

    list.sort_by_key(|r| r.width);
    println!("{list:#?}");

    let numbers: Vec<i32> = vec![1, 2, 3, 4, 5];

    // Using .map() followed by .collect()
    let doubled: Vec<i32> = numbers.iter().map(|x| x * 2).collect();
    println!("Doubled: {:?}", doubled); // Output: Doubled: [2, 4, 6, 8, 10]

    // Using .filter() followed by .sum()
    let sum_of_even: i32 = numbers.iter().filter(|x| *x % 2 == 0).sum();
    println!("Sum of even: {}", sum_of_even); // Output: Sum of even: 6 (2 + 4)

    // Using .take() followed by .for_each()
    println!("First three numbers:");
    numbers.iter().take(3).for_each(|x| println!("{}", x));
    // Output:
    // First three numbers:
    // 1
    // 2
    // 3

    // Chaining multiple adapters before .collect()
    let skip_one_double_and_collect: Vec<i32> = numbers
        .iter()
        .skip(1) // Skip the first element (1)
        .map(|x| x * 2) // Double the remaining elements
        .collect();
    println!("Skip one and double: {:?}", skip_one_double_and_collect); // Output: Skip one and double: [4, 6, 8, 10]
}

#[derive(Debug, PartialEq, Copy, Clone)]
enum ShirtColor {
    Red,
    Blue,
}

struct Inventory {
    shirts: Vec<ShirtColor>,
}

impl Inventory {
    fn giveaway(&self, user_preference: Option<ShirtColor>) -> ShirtColor {
        user_preference.unwrap_or_else(|| self.most_stocked())
    }

    fn most_stocked(&self) -> ShirtColor {
        let mut num_red = 0;
        let mut num_blue = 0;

        for color in &self.shirts {
            match color {
                ShirtColor::Red => num_red += 1,
                ShirtColor::Blue => num_blue += 1,
            }
        }
        if num_red > num_blue {
            ShirtColor::Red
        } else {
            ShirtColor::Blue
        }
    }
}

#[derive(Debug)]
pub struct Rectangle {
    width: u32,
    pub height: u32,
}

#[test]
fn iterator_sum() {
    let v1 = vec![1, 2, 3];

    let v1_iter = v1.iter();

    let total: i32 = v1_iter.sum();

    assert_eq!(total, 6);
}
