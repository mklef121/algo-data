use std::collections::HashMap;

fn main() {
    let mut v: Vec<i32> = Vec::new();

    v.push(5);
    v.push(6);

    let third: Option<&i32> = v.get(6);
    match third {
        Some(third) => println!("The third element is {third}"),
        None => println!("There is no third element."),
    }

    // or creating vectors with default values
    let v = vec![1, 2, 3];

    let third: i32 = v[2];
    println!("The third element is {third}");

    // Iterating Over the Values in a Vector
    let v = vec![100, 32, 57];
    for i in &v {
        println!("the vector val: {i}");
    }

    // We can also iterate over mutable references to each element in a mutable vector
    // in order to make changes to all the elements.
    let mut v = vec![100, 32, 57];
    for i in &mut v {
        println!("the address val: {i}, and the pointer address is: {:p}", i);
        *i += 50;
    }

    //  Using an Enum to Store Multiple Types
    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];

    for cell in row {
        match cell {
            SpreadsheetCell::Float(f) => println!("Float: {}", f),
            SpreadsheetCell::Int(i) => println!("Integer: {}", i),
            SpreadsheetCell::Text(t) => println!("Text: {}", t),
        }
    }

    /*
       {
           let v = vec![1, 2, 3, 4];

           // do stuff with v
       } // <- v goes out of scope and is freed here
    */

    let data = "initial contents";
    let _s = data.to_string();

    let _hello = String::from("السلام عليكم");
    let _hello = String::from("Dobrý den");
    let _hello = String::from("Hello");
    let _hello = String::from("שלום");
    let mut hello_eng = String::from("नमस्ते");
    let _hello = String::from("Hola");
    let _hello = String::from("안녕하세요");
    let _hello = String::from("你好");
    let _hello = String::from("Olá");
    let _hello = String::from("Здравствуйте");
    let _hello = String::from("こんにちは");

    hello_eng.push_str("new ways");

    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // note s1 has been moved here and can no longer be used

    let s4 = String::from("tic");
    let s5 = String::from("tac");
    let s6 = String::from("toe");

    let s = s4 + "-" + &s5 + "-" + &s6;
    let s10 = String::from("wetin");
    let s11 = String::from("happen");
    let s12 = String::from("for-Abule");

    let s13 = format!("{s10}-{s11}-{s12}");
    let hello = "Здравствуйте";

    let _sr = &hello[0..4];

    println!(
        "\n\n Hello, world, {}!, {s3}, advanced concact: {s}. The result is {s13}",
        hello
    );

    for c in "Зд".chars() {
        println!("{c}");
    }

    let mut scores: HashMap<String, i32> = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    let team_name = String::from("Blue");
    let score = scores.get(&team_name).copied().unwrap_or(0);

    println!("the map score {score}");

    for (key, value) in &scores {
        println!("{key}: {value}");
    }

    scores.entry(String::from("Yellow")).or_insert(60);

    let text = "hello world wonderful world";

    let mut map = HashMap::new();
    // We use a hash map with the words as keys and increment the value to keep
    // track of how many times we’ve seen that word.
    for word in text.split_whitespace() {
        let count: &mut i32 = map.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{map:#?}");
}

enum SpreadsheetCell {
    Int(i32),
    Float(f64),
    Text(String),
}
