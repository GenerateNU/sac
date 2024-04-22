DROP TABLE IF EXISTS club_events CASCADE;

CREATE TABLE club_events(
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(club_id,event_id),
    CONSTRAINT fk_club_events_event FOREIGN key(event_id) REFERENCES events(id),
    CONSTRAINT fk_club_events_club FOREIGN key(club_id) REFERENCES clubs(id)
);