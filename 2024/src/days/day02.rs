fn is_safe(report: &str) -> bool {
    let parts: Vec<i32> = report
        .split_whitespace()
        .filter_map(|part| part.parse().ok())
        .collect();

    parts.windows(2).all(|window| match window {
        [prev, curr] => {
            let diff = (curr - prev).abs();
            diff >= 1 && diff <= 3 && (*curr > *prev) == (parts[1] > parts[0])
        }
        _ => true,
    })
}

pub fn solve(input: &str) -> (i32, i32) {
    let safe_reports = input.lines().filter(|line| is_safe(line)).count() as i32;

    (safe_reports, 1)
}
