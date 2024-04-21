DROP TABLE IF EXISTS user_event_rsvps;

CREATE TABLE user_event_rsvps(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,event_id),
    CONSTRAINT fk_user_event_rsvps_user FOREIGN key(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_event_rsvps_event FOREIGN key(event_id) REFERENCES events(id)
);