DROP TABLE IF EXISTS users;

-- Create User table
CREATE TABLE users(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role varchar(255) NOT NULL DEFAULT 'student'::character varying,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password_hash varchar(97) NOT NULL,
    major0 varchar(255),
    major1 varchar(255),
    major2 varchar(255),
    college varchar(255),
    graduation_cycle varchar(255),
    graduation_year smallint,
    is_verified boolean NOT NULL DEFAULT false,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX uni_users_email ON users USING btree ("email");

-- Insert sample data into "users" table