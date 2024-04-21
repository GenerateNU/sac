DROP TABLE IF EXISTS files;

CREATE TABLE files(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    owner_id uuid NOT NULL,
    owner_type varchar(255) NOT NULL,
    file_name varchar(255) NOT NULL,
    file_type varchar(255) NOT NULL,
    file_size bigint NOT NULL,
    file_url varchar(255) NOT NULL,
    object_key varchar(255) NOT NULL,
    PRIMARY KEY(id)
);
CREATE INDEX idx_files_owner_type ON files USING btree ("owner_type");
CREATE INDEX idx_files_owner_id ON files USING btree ("owner_id");