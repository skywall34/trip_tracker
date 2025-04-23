-- Script to insert Mia's Trips from the csv file into the database

CREATE TEMP TABLE staging_trips (
  ownership       TEXT,
  departure_ts     TEXT,   -- e.g. '2010-07-08 7:30:00'
  arrival_ts       TEXT,   -- e.g. '2010-07-08 9:20:00'
  from_airport     TEXT,
  to_airport       TEXT,
  created_at       TEXT,
  updated_at       TEXT,
  notes            TEXT,
  flight_id        TEXT,
  flight_number    TEXT
);

-- sqlite> .mode csv
-- sqlite> .import your_trips.csv staging_trips


INSERT INTO trips (
  user_id,
  departure,
  arrival,
  departure_time,
  arrival_time,
  airline,
  flight_number,
  reservation,
  terminal,
  gate
)
SELECT
  14 AS user_id,
  from_airport                    AS departure,
  to_airport                      AS arrival,
  CAST(strftime('%s', departure_ts) AS INTEGER) AS departure_time,
  CAST(strftime('%s', arrival_ts)   AS INTEGER) AS arrival_time,
  ''                            AS airline,        -- ⚠️ missing in CSV
  flight_id || flight_number      AS flight_number,  -- e.g. 'AF' || '694' → 'AF694'
  ''                            AS reservation,    -- ⚠️ missing
  ''                            AS terminal,       -- ⚠️ missing
  ''                            AS gate            -- ⚠️ missing
FROM staging_trips;
