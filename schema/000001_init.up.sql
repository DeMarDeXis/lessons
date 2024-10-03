CREATE TABLE lessons
(
    lesson_id          serial       not null unique,
    course_id          int          not null,
    lesson_name        varchar(255) not null,
    lesson_description varchar(255) not null
);