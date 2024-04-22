DROP TABLE IF EXISTS series;

CREATE TABLE series(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    recurring_type varchar(255) NOT NULL DEFAULT 'open'::character varying,
    max_occurrences bigint NOT NULL,
    PRIMARY KEY(id)
);