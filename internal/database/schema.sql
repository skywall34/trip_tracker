CREATE TABLE IF NOT EXISTS airports (
    iata_code TEXT PRIMARY KEY,
    name TEXT,
    country TEXT,
    latitude REAL,
    longitude REAL,
    region TEXT
);

CREATE TABLE IF NOT EXISTS trips (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    departure TEXT NOT NULL,
    arrival TEXT NOT NULL,
    departure_time INTEGER NOT NULL,
    arrival_time INTEGER NOT NULL,
    airline TEXT NOT NULL,
    flight_number TEXT NOT NULL,
    reservation TEXT,
    terminal TEXT,
    gate TEXT,
    departure_lat REAL NOT NULL,
    departure_lon REAL NOT NULL,
    arrival_lat REAL NOT NULL,
    arrival_lon REAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    `password` TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id TEXT UNIQUE NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_trips_user_id ON trips(user_id);
CREATE INDEX idx_trips_departure ON trips(departure);
CREATE INDEX idx_trips_arrival ON trips(arrival);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_session_id ON sessions(session_id);