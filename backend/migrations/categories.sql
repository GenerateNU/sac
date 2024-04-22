DROP TABLE IF EXISTS categories;

CREATE TABLE categories(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX uni_categories_name ON categories USING btree ("name");