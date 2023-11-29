CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT
);

CREATE TABLE IF NOT EXISTS rentals (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name TEXT,
    type TEXT,
    description TEXT,
    sleeps INTEGER,
    price_per_day BIGINT,
    home_city TEXT,
    home_state TEXT,
    home_zip TEXT,
    home_country TEXT,
    vehicle_make TEXT,
    vehicle_model TEXT,
    vehicle_year INTEGER,
    vehicle_length NUMERIC(4,2),
    created TIMESTAMPTZ,
    updated TIMESTAMPTZ,
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION,
    primary_image_url TEXT
);
