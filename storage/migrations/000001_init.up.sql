CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    doe_date date not null,
    created_at date not null,
    updated_at date not null
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);
