mod db;

fn main() {
    let usage = "Usage: ./done <task>";
    let task = std::env::args().nth(1).unwrap_or_else(|| {
        eprintln!("{}", usage);
        std::process::exit(1);
    });
    let task = task.trim();
    // println!("Task: {}", task);
    let conn = db::init();
    db::store(&conn, task);
    db::list(&conn);
}
