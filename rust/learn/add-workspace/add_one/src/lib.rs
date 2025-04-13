pub fn add_one(x: i32) -> i32 {
    x + 1
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn add_one_test() {
        let result = add_one(2);
        assert_eq!(result, 3);
    }
}
