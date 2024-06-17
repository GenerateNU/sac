BEGIN;

ALTER TABLE clubs
    DROP COLUMN IF EXISTS clubsearch_index_col;

DROP INDEX IF EXISTS clubsearch_index;

ALTER TABLE "events"
    DROP COLUMN IF EXISTS eventsearch_index_col;

DROP INDEX IF EXISTS eventsearch_index;

COMMIT;