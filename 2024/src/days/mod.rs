pub mod day01;
pub mod day02;

use std::collections::HashMap;

pub fn get_solver(day: u8) -> Option<fn(&str) -> (i32, i32)> {
    let mut solvers: HashMap<u8, fn(&str) -> (i32, i32)> = HashMap::new();

    solvers.insert(1, day01::solve);
    solvers.insert(2, day02::solve);

    solvers.get(&day).copied()
}
