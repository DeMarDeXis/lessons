BEGIN;

CREATE TYPE lesson_type_enum AS ENUM ('lecture', 'practice');

CREATE TYPE status_enum AS ENUM ('Not start', 'send', 'rejected');

CREATE TABLE lessons
(
    lesson_id          serial       not null unique,
    lesson_type        lesson_type_enum not null,
    lesson_name        varchar(255) not null,
    lesson_description varchar(255) not null,
    lesson_file_name   varchar(255),
    lesson_file_content BYTEA,
    lesson_status      status_enum not null default 'Not start'
);

COMMIT;