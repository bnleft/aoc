fn split_line(line: &str) -> (i32, i32) {
    let parts = line.split_whitespace().collect::<Vec<&str>>();
    let first = parts[0].parse::<i32>().unwrap();
    let second = parts[1].parse::<i32>().unwrap();

    (first, second)
}

pub fn solve(input: &str) -> (i32, i32) {
    let (mut left_numbers, mut right_numbers): (Vec<i32>, Vec<i32>) = input
        .lines()
        .map(split_line)
        .fold((Vec::new(), Vec::new()), |(mut left_vec, mut right_vec), (left, right)| {
            left_vec.push(left);
            right_vec.push(right);
            (left_vec, right_vec)
        });

    left_numbers.sort();
    right_numbers.sort();

    let min_length = left_numbers.len().min(right_numbers.len());

    let distance_sum: i32 = left_numbers
        .iter()
        .take(min_length)
        .zip(right_numbers.iter())
        .map(|(left, right)| (left - right).abs())
        .sum();

    (distance_sum, 1)
}
