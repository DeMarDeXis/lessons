BEGIN;

-- ALTER TABLE lessons ADD COLUMN status VARCHAR(10);

CREATE TYPE check_status_enum AS ENUM ('On checking', 'done', 'rejected');

CREATE TABLE teachers_checklist (
    id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL,
    lesson_id INT NOT NULL,
    status check_status_enum NOT NULL DEFAULT 'On checking',
    homework bytea,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;