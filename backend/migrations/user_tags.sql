DROP TABLE IF EXISTS user_tags;

CREATE TABLE user_tags(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,tag_id),
    CONSTRAINT fk_user_tags_user FOREIGN key(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_tags_tag FOREIGN key(tag_id) REFERENCES tags(id)
);