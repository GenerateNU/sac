DROP TABLE IF EXISTS user_club_members;

CREATE TABLE user_club_members(
    user_id uuid NOT NULL,
    club_id uuid NOT NULL,
    membership_type varchar(255) NOT NULL DEFAULT 'member'::character varying,
    PRIMARY KEY(user_id,club_id),
    CONSTRAINT fk_user_club_members_user FOREIGN key(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_club_members_club FOREIGN key(club_id) REFERENCES clubs(id)
);