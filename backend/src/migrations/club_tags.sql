DROP TABLE IF EXISTS club_tags;

CREATE TABLE club_tags(
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(tag_id,club_id),
    CONSTRAINT fk_club_tags_tag FOREIGN key(tag_id) REFERENCES tags(id),
    CONSTRAINT fk_club_tags_club FOREIGN key(club_id) REFERENCES clubs(id)
);