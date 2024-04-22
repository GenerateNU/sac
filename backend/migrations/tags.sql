DROP TABLE IF EXISTS tags;

CREATE TABLE tags(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    category_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_categories_tag FOREIGN key(category_id) REFERENCES categories(id)
);