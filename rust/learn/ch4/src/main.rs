fn main() {
    let s1 = String::from("hello");

    let (s1, len) = calculate_length(s1);

    println!("The length of '{s1}' is {len}.");

    // using references
    let s1 = String::from("Returns the length of this");

    let len = calculate_length2(&s1);

    println!("The length of '{s1}' is {len}.");

    let mut s = String::from("original string");
    change(&mut s);
    change(&mut s);
    change(&mut s);

    println!("The value of mutable s is '{s}'.");

    let mut s = String::from("hello");

    let r1 = &s; // no problem
    let r2 = &s; // no problem
    println!("{r1} and {r2}");
    // variables r1 and r2 will not be used after this point

    println!("another {r1} and {r2}");

    let r3 = &mut s; // no problem
    println!("{r3}");

    let a = [1, 2, 3, 4, 5];
    let slice = &a[1..3];
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() returns the length of a String

    (s, length)
}

fn calculate_length2(s: &String) -> usize {
    s.len()
}

fn change(s: &mut String) {
    s.push_str(", add some more string to the original address");
}

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
