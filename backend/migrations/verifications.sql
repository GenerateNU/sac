DROP TABLE IF EXISTS verifications CASCADE;

CREATE TABLE verifications(
    user_id varchar(36) NOT NULL,
    token varchar(255),
    expires_at timestamp without time zone NOT NULL,
    "type" varchar(255) NOT NULL,
    PRIMARY KEY(user_id,expires_at)
);
CREATE UNIQUE INDEX uni_verifications_token ON verifications USING btree ("token");