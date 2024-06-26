BEGIN;

CREATE EXTENSION vector;

ALTER TABLE clubs
    ADD COLUMN clubsearch_index_col tsvector
        GENERATED ALWAYS AS (to_tsvector('english', coalesce(name, '') || ' ' || coalesce(preview, '') || ' ' || coalesce(description, ''))) STORED;

ALTER TABLE clubs
    ADD COLUMN embedding vector(512);

CREATE INDEX
    clubsearch_index ON clubs using GIN(clubsearch_index_col);

ALTER TABLE "events"
    ADD COLUMN eventsearch_index_col tsvector
        GENERATED ALWAYS AS (to_tsvector('english', coalesce(name, '') || ' ' || coalesce(preview, '') || ' ' || coalesce(description, ''))) STORED;

ALTER TABLE "events"
    ADD COLUMN embedding vector(512);

CREATE INDEX
    eventsearch_index on "events" using GIN(eventsearch_index_col);
     
COMMIT;