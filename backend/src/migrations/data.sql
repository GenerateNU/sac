-- BEGIN SUPER CREATION TRANSACTION
BEGIN;

-- SUPER USER
INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '000000000', 'generatesac@gmail.com', 'rust', 'SAC', 'Super', 'KCCS', 1);

-- SAC SUPER CLUB
INSERT INTO clubs (name, preview, description, num_members, logo) VALUES ('SAC', 'SAC', 'SAC', 1, 'foo');

-- INSERT SUPER USER INTO SAC SUPER CLUB
INSERT INTO user_club_members (club_id, user_id) VALUES (1, 1);

COMMIT;
-- END SUPER CREATION TRANSACTION

-- BEGIN MOCK DATA TRANSACTION
BEGIN;

INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002183108', 'oduneye.d@northeastern.edu', 'rust', 'David', 'Oduneye', 'KCCS', 3);
INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002172052', 'ladley.g@northeastern.edu', 'rust', 'Garrett', 'Ladley', 'KCCS', 3);

COMMIT;
-- END MOCK DATA TRANSACTION