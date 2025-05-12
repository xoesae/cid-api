CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    code VARCHAR(10) NOT NULL UNIQUE,
    name TEXT NOT NULL
);

CREATE INDEX idx_categories_code ON categories (code);