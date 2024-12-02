mod days;

use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: {} <day_number>", args[0]);
        return;
    }

    let day: u8 = args[1].parse().unwrap_or_else(|_| {
        eprintln!("Invalid day number");
        std::process::exit(1);
    });

    let input = fs::read_to_string(format!("data/day{:02}/input.txt", day)).unwrap();

    if let Some(solver) = days::get_solver(day) {
        let (p1, p2) = solver(&input);
        println!("=== Day {:02} ===", day);
        println!("Part 1: {}", p1);
        println!("Part 2: {}", p2);
    } else {
        eprintln!("Day {:02} is not implemented", day);
    }
}
