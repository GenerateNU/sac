DROP TABLE IF EXISTS clubs;

CREATE TABLE clubs(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    soft_deleted_at timestamp with time zone,
    name varchar(255) NOT NULL,
    preview varchar(255) NOT NULL,
    description text NOT NULL,
    num_members bigint NOT NULL,
    is_recruiting boolean NOT NULL DEFAULT false,
    recruitment_cycle varchar(255) NOT NULL DEFAULT 'always'::character varying,
    recruitment_type varchar(255) NOT NULL DEFAULT 'unrestricted'::character varying,
    weekly_time_commitment bigint,
    one_word_to_describe_us varchar(255) DEFAULT NULL::character varying,
    application_link varchar(255) DEFAULT NULL::character varying,
    logo varchar(255) DEFAULT NULL::character varying,
    parent text,
    PRIMARY KEY(id)
);