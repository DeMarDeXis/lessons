ALTER TABLE lessons ADD COLUMN status VARCHAR(10);

CREATE TYPE status_enum AS ENUM ('draft', 'done', 'rejected');

ALTER TABLE lessons ALTER COLUMN status TYPE status_enum USING status::status_enum;

CREATE TABLE teachers_checklist (
    id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL,
    lesson_id INT NOT NULL,
    status status_enum NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);