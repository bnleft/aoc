use std::collections::HashMap;

fn split_line(line: &str) -> (i32, i32) {
    let locations: Vec<&str> = line.split_whitespace().collect();
    let first: i32 = locations[0].parse().unwrap();
    let second: i32 = locations[1].parse().unwrap();

    (first, second)
}

pub fn solve(input: &str) -> (i32, i32) {
    let (mut left_numbers, mut right_numbers): (Vec<i32>, Vec<i32>) =
        input.lines().map(split_line).unzip();

    left_numbers.sort();
    right_numbers.sort();

    let distance_sum = left_numbers
        .iter()
        .zip(right_numbers.iter())
        .map(|(left, right)| (left - right).abs())
        .sum();

    let mut right_counts = HashMap::new();
    right_numbers.iter().for_each(|&right| {
        *right_counts.entry(right).or_insert(0) += 1;
    });

    let similarity_sum = left_numbers
        .iter()
        .map(|&left| left * *right_counts.get(&left).unwrap_or(&0))
        .sum();

    (distance_sum, similarity_sum)
}
