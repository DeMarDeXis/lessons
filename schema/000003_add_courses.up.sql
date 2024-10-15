create table courses
(
    id serial primary key,
    name varchar(255) not null,
    description text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    owner_id int not null
);

create table lessons_courses
(
    id serial not null unique,
    lesson_id int REFERENCES lessons (lesson_id) on delete cascade NOT NULL,
    course_id int REFERENCES courses (id) on delete cascade NOT NULL
);
