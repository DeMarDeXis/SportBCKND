CREATE TABLE nhl_schedule
(
    id SERIAL NOT NULL PRIMARY KEY,
    date_game DATE NOT NULL,
    time_game TIME NOT NULL,
    visitor_team_id INTEGER NOT NULL REFERENCES nhl_teams(id) ON DELETE RESTRICT,
    home_team_id INTEGER NOT NULL REFERENCES nhl_teams(id) ON DELETE RESTRICT,
    visitor_score INTEGER,
    home_score INTEGER,
    attendance INTEGER,
    game_duration VARCHAR(10),
    is_overtime BOOLEAN default false,
    venue VARCHAR(255),
    notes TEXT,
    created_at timestamptz not null DEFAULT NOW(),
    updated_at timestamptz not null DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER trigg_update_nhl_schedule_updated_at
BEFORE UPDATE ON nhl_schedule
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE UNIQUE INDEX idx_nhl_schedule_unique
    ON nhl_schedule (date_game, time_game, visitor_team_id, home_team_id);
