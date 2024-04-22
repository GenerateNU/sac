DROP TABLE IF EXISTS event_series;

CREATE TABLE event_series(
    event_id uuid NOT NULL,
    series_id uuid NOT NULL,
    CONSTRAINT fk_event_series_event FOREIGN key(event_id) REFERENCES events(id),
    CONSTRAINT fk_event_series_series FOREIGN key(series_id) REFERENCES series(id)
);