mod days;

use days::day01;
use std::fs;

fn main() {
    let day = 1;
    let input = fs::read_to_string("data/day01/input.txt").unwrap();
    let (p1, p2) = day01::solve(&input);

    println!("=== Day {:02} ===", day);
    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);
}
