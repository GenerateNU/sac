DROP TABLE IF EXISTS events;

CREATE TABLE events(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    preview varchar(255) NOT NULL,
    content text NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    location varchar(255) NOT NULL,
    event_type varchar(255) NOT NULL DEFAULT 'open'::character varying,
    is_recurring boolean NOT NULL DEFAULT false,
    meeting_link varchar(255),
    host uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_host_event FOREIGN key(host) REFERENCES clubs(id)
);