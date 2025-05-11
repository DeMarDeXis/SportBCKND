CREATE TABLE nhl_teams
(
    id serial not null unique,
    id_team int references teams(id) on delete cascade not null,
    conference varchar(255) not null,
    division varchar(255) not null
);

CREATE TABLE nhl_roster(
    id serial not null unique,
    id_team int references nhl_teams(id) on delete cascade not null,
    name varchar(255) not null,
    surname varchar(255) not null,
    number int not null,
    position varchar(255) not null,
    hand varchar(255) not null,
    age int not null,
    acquired_at varchar(255) not null,
    birth_place varchar(255) not null,
    role varchar(5) not null,
    injured boolean not null
);