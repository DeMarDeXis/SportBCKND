BEGIN;

DROP INDEX IF EXISTS idx_nhl_schedule_unique;
DROP TRIGGER IF EXISTS trigg_update_nhl_schedule_updated_at ON nhl_schedule;
DROP FUNCTION IF EXISTS update_updated_at_column;
DROP TABLE IF EXISTS nhl_schedule;

COMMIT;