BEGIN;

ALTER TABLE clubs
    DROP COLUMN IF EXISTS clubsearch_index_col;

DROP INDEX IF EXISTS clubsearch_index;

COMMIT;