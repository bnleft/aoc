mod days;

use std::fs;
use days::{day01};

fn main() {
    let day = 1;
    let input = fs::read_to_string("data/inputs/day01/p1.txt").unwrap();
    let (p1, p2) = day01::solve(&input);

    println!("=== Day {:02} ===", day);
    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);
}
