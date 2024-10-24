BEGIN;

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash VARCHAR(255) not null
);


CREATE TABLE teams
(
    img_url varchar(255) not null,
    id serial not null unique,
    name varchar(50) not null,
    abbreviation varchar(3) not null
);

CREATE TABLE users_teams
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

-- CREATE TABLE users_admins

COMMIT;