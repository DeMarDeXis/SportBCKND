CREATE TABLE nfl_teams
(
    id serial not null unique,
    id_team int references teams(id) on delete cascade not null,
    com_used_abbr varchar(255) default null,
    conference varchar(255) not null,
    division varchar(255) not null
);

CREATE TABLE nba_teams
(
    id serial not null unique,
    id_team int references teams(id) on delete cascade not null,
    com_used_abbr varchar(255) default null,
    conference varchar(255) not null,
    division varchar(255) not null
);