-- BEGIN MOCK DATA TRANSACTION
BEGIN;

INSERT INTO users (id, role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('29cac84a-362c-4ffa-9f4c-2f76057b7902', 'super', '002183108', 'oduneye.d@northeastern.edu', '$argon2id$v=19$m=65536,t=3,p=2$zYyFSnLvC5Q482mzMJrTjg$WUhpXwulvfipyWg7asQyCRUqBEnjizDOoMP2/GvWQR8', 'David', 'Oduneye', 'KCCS', 3);
INSERT INTO users (id, role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('4f4d9990-7d26-4229-911d-1aa61851c292', 'super', '002172052', 'ladley.g@northeastern.edu', '$argon2id$v=19$m=65536,t=3,p=2$zYyFSnLvC5Q482mzMJrTjg$WUhpXwulvfipyWg7asQyCRUqBEnjizDOoMP2/GvWQR8', 'Garrett', 'Ladley', 'KCCS', 3);
INSERT INTO user_club_members (user_id, club_id, type) VALUES ('29cac84a-362c-4ffa-9f4c-2f76057b7902', (SELECT id FROM clubs WHERE name = 'SAC'), 'admin');
INSERT INTO user_club_members (user_id, club_id, type) VALUES ('4f4d9990-7d26-4229-911d-1aa61851c292', (SELECT id FROM clubs WHERE name = 'SAC'), 'admin');
COMMIT;
-- END MOCK DATA TRANSACTION
