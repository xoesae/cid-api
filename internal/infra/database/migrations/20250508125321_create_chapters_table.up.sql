CREATE TABLE chapters (
    id SERIAL PRIMARY KEY,
    code_start VARCHAR(10) NOT NULL,
    code_end VARCHAR(10) NOT NULL,
    roman VARCHAR(5) NOT NULL,
    name TEXT NOT NULL
);

CREATE INDEX idx_chapters_code_range ON chapters (code_start, code_end);
