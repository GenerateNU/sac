BEGIN;

ALTER TABLE clubs
    ADD COLUMN clubsearch_index_col tsvector
        GENERATED ALWAYS AS (to_tsvector('english', coalesce(name, '') || ' ' || coalesce(preview, '') || ' ' || coalesce(description, ''))) STORED;

CREATE INDEX
    clubsearch_index ON clubs using GIN(clubsearch_index_col);
     
COMMIT;