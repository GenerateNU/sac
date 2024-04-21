DROP TABLE IF EXISTS user_event_waitlists;

CREATE TABLE user_event_waitlists(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,event_id),
    CONSTRAINT fk_user_event_waitlists_user FOREIGN key(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_event_waitlists_event FOREIGN key(event_id) REFERENCES events(id)
);