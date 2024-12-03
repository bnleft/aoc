use regex::Regex;

pub fn solve(input: &str) -> (i32, i32) {
    let regex = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();

    let mul_sum: i32 = regex
        .captures_iter(input)
        .map(|mul| {
            let x: i32 = mul[1].parse().unwrap();
            let y: i32 = mul[2].parse().unwrap();
            x * y
        })
        .sum();

    (mul_sum, 1)
}
