use aggregator::Tweet;

fn main() {
    let number_list = vec![102, 34, 6000, 89, 54, 2, 43, 8];
    let integer_point = Point { x: 5, y: 10 };
    let double = PointTwo { x: 4, y: 7.90 };
    println!(
        "Hello, world!, First largest: {}, the new largest is {}, the generic based largest {}, => integer_point{:#?},  => double{:#?}",
        find_largest_number(),
        largest_new(&number_list),
        largest(&number_list),
        integer_point,
        double
    );

    let p1 = PointTwo { x: 5, y: 10.4 };
    let p2 = PointTwo { x: "Hello", y: 'c' };

    let p3 = p1.mixup(p2);

    println!("p3.x = {}, p3.y = {}", p3.x, p3.y);

    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };

    println!("1 new tweet: {}", tweet.summarize());

    notify(&tweet);
    notify_second(&tweet);
}

// Generics

fn largest<T: std::cmp::PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn find_largest_number() -> i32 {
    let number_list = vec![34, 50, 25, 100, 65];
    let mut largest = number_list[0];

    for n in number_list {
        if n > largest {
            largest = n
        }
    }

    largest
}

fn largest_new(list: &[i32]) -> &i32 {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

#[derive(Debug)]
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

#[derive(Debug)]
struct PointTwo<T, U> {
    x: T,
    y: U,
}

impl<X1, Y1> PointTwo<X1, Y1> {
    fn mixup<X2, Y2>(self, other: PointTwo<X2, Y2>) -> PointTwo<X1, Y2> {
        PointTwo {
            x: self.x,
            y: other.y,
        }
    }
}

// traits

pub trait Summary {
    fn summarize(&self) -> String;
}

mod aggregator {
    use crate::Summary;

    pub struct NewsArticle {
        pub headline: String,
        pub location: String,
        pub author: String,
        pub content: String,
    }

    impl Summary for NewsArticle {
        fn summarize(&self) -> String {
            format!("{}, by {} ({})", self.headline, self.author, self.location)
        }
    }

    pub struct Tweet {
        pub username: String,
        pub content: String,
        pub reply: bool,
        pub retweet: bool,
    }

    impl Summary for Tweet {
        fn summarize(&self) -> String {
            format!("{}: {}", self.username, self.content)
        }
    }
}

pub fn notify(param: &impl Summary) {
    println!("Breaking news! {}", param.summarize());
}

// using trait bounds on generics
pub fn notify_second<T: Summary>(param: &T) {
    println!("Breaking news! {}", param.summarize());
}
