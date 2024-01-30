-- BEGIN MOCK DATA TRANSACTION
BEGIN;

INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002183108', 'oduneye.d@northeastern.edu', '$argon2id$v=19$m=65536,t=3,p=2$zYyFSnLvC5Q482mzMJrTjg$WUhpXwulvfipyWg7asQyCRUqBEnjizDOoMP2/GvWQR8', 'David', 'Oduneye', 'KCCS', 3);
INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002172052', 'ladley.g@northeastern.edu', '$argon2id$v=19$m=65536,t=3,p=2$zYyFSnLvC5Q482mzMJrTjg$WUhpXwulvfipyWg7asQyCRUqBEnjizDOoMP2/GvWQR8', 'Garrett', 'Ladley', 'KCCS', 3);

COMMIT;
-- END MOCK DATA TRANSACTION
