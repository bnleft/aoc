use regex::Regex;

fn calculate_mul_sum(input: &str) -> i32 {
    let mul_pattern = r"mul\((\d{1,3}),(\d{1,3})\)";
    let regex = Regex::new(mul_pattern).unwrap();

    regex
        .captures_iter(input)
        .map(|captures| {
            let x: i32 = captures[1].parse().unwrap();
            let y: i32 = captures[2].parse().unwrap();
            x * y
        })
        .sum()
}

fn calculate_toggled_mul_sum(input: &str) -> i32 {
    let mul_pattern = r"mul\((\d{1,3}),(\d{1,3})\)";
    let do_pattern = r"do\(\)";
    let dont_pattern = r"don't\(\)";
    let combined_pattern = format!("{}|{}|{}", mul_pattern, do_pattern, dont_pattern);
    let regex = Regex::new(&combined_pattern).unwrap();

    regex
        .captures_iter(input)
        .scan(true, |toggle, captures| {
            match captures.get(0).unwrap().as_str() {
                "do()" => {
                    *toggle = true;
                    Some(0)
                }
                "don't()" => {
                    *toggle = false;
                    Some(0)
                }
                _ if *toggle => {
                    let x: i32 = captures[1].parse().unwrap();
                    let y: i32 = captures[2].parse().unwrap();
                    Some(x * y)
                }
                _ => Some(0),
            }
        })
        .sum()
}

pub fn solve(input: &str) -> (i32, i32) {
    let mul_sum: i32 = calculate_mul_sum(input);
    let enabled_mul_sum: i32 = calculate_toggled_mul_sum(input);

    (mul_sum, enabled_mul_sum)
}
