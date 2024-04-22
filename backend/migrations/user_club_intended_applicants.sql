DROP TABLE IF EXISTS user_club_intended_applicants;

CREATE TABLE user_club_intended_applicants(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,club_id),
    CONSTRAINT fk_user_club_intended_applicants_user FOREIGN key(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_club_intended_applicants_club FOREIGN key(club_id) REFERENCES clubs(id)
);