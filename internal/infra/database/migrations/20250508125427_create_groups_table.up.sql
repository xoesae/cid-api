CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    chapter_id INTEGER NOT NULL REFERENCES chapters(id) ON DELETE CASCADE,
    code_start VARCHAR(10) NOT NULL,
    code_end VARCHAR(10) NOT NULL,
    name TEXT NOT NULL
);

CREATE INDEX idx_groups_code_range ON groups (code_start, code_end);