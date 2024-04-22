DROP TABLE IF EXISTS event_tags;

CREATE TABLE event_tags(
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(tag_id,event_id),
    CONSTRAINT fk_event_tags_tag FOREIGN key(tag_id) REFERENCES tags(id),
    CONSTRAINT fk_event_tags_event FOREIGN key(event_id) REFERENCES events(id)
);