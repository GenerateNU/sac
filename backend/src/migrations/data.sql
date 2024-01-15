-- BEGIN MOCK DATA TRANSACTION
BEGIN;

INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002183108', 'oduneye.d@northeastern.edu', 'rust', 'David', 'Oduneye', 'KCCS', 3);
INSERT INTO users (role, nuid, email, password_hash, first_name, last_name, college, year) VALUES ('super', '002172052', 'ladley.g@northeastern.edu', 'rust', 'Garrett', 'Ladley', 'KCCS', 3);

COMMIT;
-- END MOCK DATA TRANSACTION
