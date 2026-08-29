CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    category TEXT,
    price DOUBLE PRECISION,
    stock INTEGER
);

CREATE TABLE IF NOT EXISTS sales (
    id SERIAL PRIMARY KEY,
    product_name TEXT,
    quantity INTEGER,
    buyer TEXT,
    date TEXT,
    total DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    farmer_name TEXT,
    crop_name TEXT,
    land_area TEXT,
    task TEXT,
    date TEXT,
    status TEXT
);
