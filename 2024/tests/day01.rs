use aoc::days::day01::solve;
use std::fs;

#[test]
fn test_part1() {
    let input = fs::read_to_string("data/examples/day01/p1.txt").unwrap();
    let (p1, _) = solve(&input);
    assert_eq!(p1, 11);
}