CREATE TABLE subcategories (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    code VARCHAR(10) NOT NULL UNIQUE,
    name TEXT NOT NULL
);

CREATE INDEX idx_subcategories_code ON subcategories (code);