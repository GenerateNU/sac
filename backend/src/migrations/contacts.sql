DROP TABLE IF EXISTS contacts;

CREATE TABLE contacts(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "type" varchar(255) NOT NULL,
    content varchar(255) NOT NULL,
    club_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_contact FOREIGN key(club_id) REFERENCES clubs(id)
);

CREATE UNIQUE INDEX idx_contact_type ON contacts USING btree ("type","club_id");