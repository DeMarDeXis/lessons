BEGIN;

create table student_courses
(
    id serial primary key,
    course_id int REFERENCES courses (id) on delete cascade NOT NULL,
--     student_id int REFERENCES students (id) on delete cascade NOT NULL,
    student_id int NOT NULL
);

COMMIT;