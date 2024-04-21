DROP TABLE IF EXISTS point_of_contacts;

CREATE TABLE point_of_contacts(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    position varchar(255) NOT NULL,
    club_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_point_of_contact FOREIGN key(club_id) REFERENCES clubs(id)
);
CREATE UNIQUE INDEX compositeindex ON point_of_contacts USING btree ("email","club_id");
CREATE INDEX idx_point_of_contacts_club_id ON point_of_contacts USING btree ("club_id");
CREATE INDEX idx_point_of_contacts_email ON point_of_contacts USING btree ("email");