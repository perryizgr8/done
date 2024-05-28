use chrono;
use sqlite::{Connection, State};

pub fn init() -> Connection {
    let conn: Connection = sqlite::open("tasks.db").unwrap();
    conn.execute("CREATE TABLE IF NOT EXISTS tasks (task TEXT, date TEXT)")
        .unwrap();
    conn
}

pub fn store(conn: &Connection, task: &str) {
    let query = "INSERT INTO tasks (task, date) VALUES (?1, ?2)";
    let mut stmt = conn.prepare(query).unwrap();
    stmt.bind((1, task)).unwrap();
    stmt.bind((2, chrono::Local::now().to_string().as_str())).unwrap();
    stmt.next().unwrap();
}

pub fn list(conn: &Connection) {
    let mut stmt = conn.prepare("SELECT task, date FROM tasks").unwrap();
    while let State::Row = stmt.next().unwrap() {
        let task: String = stmt.read(0).unwrap();
        let date: String = stmt.read(1).unwrap();
        println!("{}: {}", date, task);
    }
}
