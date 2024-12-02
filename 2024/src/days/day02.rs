fn parse_levels(report: &str) -> Vec<i32> {
    report
        .split_whitespace()
        .filter_map(|level| level.parse().ok())
        .collect()
}

fn is_safe_pair(diff: i32, increasing: bool) -> bool {
    (1..=3).contains(&diff.abs()) && (diff > 0) == increasing
}

fn is_safe(levels: &Vec<i32>) -> bool {
    let increasing = levels[1] > levels[0];

    levels
        .windows(2)
        .all(|pair| is_safe_pair(pair[1] - pair[0], increasing))
}

fn is_tolerable(levels: &Vec<i32>) -> bool {
    if is_safe(levels) {
        return true;
    }

    (0..levels.len()).any(|index| {
        is_safe(
            &levels
                .iter()
                .enumerate()
                .filter(|(i, _)| *i != index)
                .map(|(_, &v)| v)
                .collect::<Vec<i32>>(),
        )
    })
}

pub fn solve(input: &str) -> (i32, i32) {
    let safe_reports = input
        .lines()
        .filter(|line| is_safe(&parse_levels(line)))
        .count() as i32;
    let tolerable_reports = input
        .lines()
        .filter(|line| is_tolerable(&parse_levels(line)))
        .count() as i32;

    (safe_reports, tolerable_reports)
}
