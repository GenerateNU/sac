BEGIN;

ALTER TABLE clubs
    DROP COLUMN IF EXISTS clubsearch_index_col;

ALTER TABLE clubs
    DROP COLUMN IF EXISTS embedding;

DROP INDEX IF EXISTS clubsearch_index;

ALTER TABLE "events"
    DROP COLUMN IF EXISTS eventsearch_index_col;

ALTER TABLE "events"
    DROP COLUMN IF EXISTS embedding;

DROP INDEX IF EXISTS eventsearch_index;

DROP EXTENSION IF EXISTS vector;

COMMIT;