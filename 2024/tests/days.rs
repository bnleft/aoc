use aoc::days;
use std::fs;

#[test]
fn test_day01_solve() {
    let input = fs::read_to_string("data/day01/example.txt").unwrap();
    let (p1, p2) = days::day01::solve(&input);
    assert_eq!(p1, 11);
    assert_eq!(p2, 31);
}

#[test]
fn test_day02_solve() {
    let input = fs::read_to_string("data/day02/example.txt").unwrap();
    let (p1, p2) = days::day02::solve(&input);
    assert_eq!(p1, 2);
    assert_eq!(p2, 4);
}
