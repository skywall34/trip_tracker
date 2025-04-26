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
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE,                     -- Nullable for Google users
    `password` TEXT,                            -- Nullable for Google users
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL,
    google_id TEXT UNIQUE,                   -- New: stores Google account ID
    auth_provider TEXT DEFAULT 'local',      -- 'local' or 'google'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id TEXT UNIQUE NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS password_reset_tokens (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  token_hash TEXT NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  used BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_password_reset_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_trips_user_id ON trips(user_id);
CREATE INDEX idx_trips_departure ON trips(departure);
CREATE INDEX idx_trips_arrival ON trips(arrival);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_session_id ON sessions(session_id);