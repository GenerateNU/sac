BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
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

CREATE UNIQUE INDEX IF NOT EXISTS uni_users_email ON users USING btree ("email");

CREATE TABLE IF NOT EXISTS clubs(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    soft_deleted_at timestamp with time zone,
    name varchar(255) NOT NULL,
    preview varchar(255) NOT NULL,
    description text NOT NULL,
    num_members bigint NOT NULL,
    is_recruiting boolean NOT NULL DEFAULT false,
    recruitment_cycle varchar(255) NOT NULL DEFAULT 'always'::character varying,
    recruitment_type varchar(255) NOT NULL DEFAULT 'unrestricted'::character varying,
    weekly_time_commitment bigint,
    one_word_to_describe_us varchar(255) DEFAULT NULL::character varying,
    application_link varchar(255) DEFAULT NULL::character varying,
    logo varchar(255) DEFAULT NULL::character varying,
    parent text,
    PRIMARY KEY(id)
);

--CREATE UNIQUE INDEX IF NOT EXISTS uni_clubs_name ON clubs USING btree ("name");
--CREATE INDEX IF NOT EXISTS idx_clubs_num_members ON clubs USING btree ("num_members");
--CREATE INDEX IF NOT EXISTS idx_clubs_one_word_to_describe_us ON clubs USING btree ("one_word_to_describe_us");

CREATE TABLE IF NOT EXISTS series(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS events(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    preview varchar(255) NOT NULL,
    description text NOT NULL,
    event_type varchar(255) NOT NULL,
    location varchar(255),
    link varchar(255),
    is_public boolean NOT NULL,
    is_draft boolean NOT NULL,
    is_archived boolean NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    host uuid NOT NULL,
    series_id uuid,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_host_event FOREIGN key(host) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_series_event FOREIGN key(series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_events_series_id ON events USING btree ("series_id");

CREATE TABLE IF NOT EXISTS categories(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uni_categories_name ON categories USING btree ("name");


CREATE TABLE IF NOT EXISTS tags(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    category_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_categories_tag FOREIGN key(category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_tags_category_id ON tags USING btree ("category_id");

CREATE TABLE IF NOT EXISTS club_events(
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(club_id,event_id),
    CONSTRAINT fk_club_events_event FOREIGN key(event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_club_events_club FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS club_tags(
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(tag_id,club_id),
    CONSTRAINT fk_club_tags_tag FOREIGN key(tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_club_tags_club FOREIGN key(club_id) REFERENCES clubs(id ) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS contacts(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "type" varchar(255) NOT NULL,
    content varchar(255) NOT NULL,
    club_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_contact FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_contact_type ON contacts USING btree ("type","club_id");

CREATE TABLE IF NOT EXISTS event_tags(
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(tag_id,event_id),
    CONSTRAINT fk_event_tags_tag FOREIGN key(tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_event_tags_event FOREIGN key(event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS files(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    owner_id uuid NOT NULL,
    owner_type varchar(255) NOT NULL,
    file_name varchar(255) NOT NULL,
    file_type varchar(255) NOT NULL,
    file_size bigint NOT NULL,
    file_url varchar(255) NOT NULL,
    object_key varchar(255) NOT NULL,
    PRIMARY KEY(id)
);
CREATE INDEX IF NOT EXISTS idx_files_owner_type ON files USING btree ("owner_type");
CREATE INDEX IF NOT EXISTS idx_files_owner_id ON files USING btree ("owner_id");

-- CREATE TABLE IF NOT EXISTS notifications(
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     send_at timestamp with time zone NOT NULL,
--     title varchar(255) NOT NULL,
--     content varchar(255) NOT NULL,
--     deep_link varchar(255) NOT NULL,
--     icon varchar(255) NOT NULL,
--     reference_id uuid NOT NULL,
--     reference_type varchar(255) NOT NULL,
--     PRIMARY KEY(id)
-- );

CREATE TABLE IF NOT EXISTS point_of_contacts(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    position varchar(255) NOT NULL,
    club_id uuid NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_clubs_point_of_contact FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS compositeindex ON point_of_contacts USING btree ("email","club_id");
CREATE INDEX IF NOT EXISTS idx_point_of_contacts_club_id ON point_of_contacts USING btree ("club_id");
CREATE INDEX IF NOT EXISTS idx_point_of_contacts_email ON point_of_contacts USING btree ("email");


CREATE TABLE IF NOT EXISTS user_club_followers(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,club_id),
    CONSTRAINT fk_user_club_followers_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_club_followers_club FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_club_intended_applicants(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,club_id),
    CONSTRAINT fk_user_club_intended_applicants_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_club_intended_applicants_club FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_club_members(
    user_id uuid NOT NULL,
    club_id uuid NOT NULL,
    membership_type varchar(255) NOT NULL DEFAULT 'member'::character varying,
    PRIMARY KEY(user_id,club_id),
    CONSTRAINT fk_user_club_members_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_club_members_club FOREIGN key(club_id) REFERENCES clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_event_rsvps(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,event_id),
    CONSTRAINT fk_user_event_rsvps_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_event_rsvps_event FOREIGN key(event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_event_waitlists(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,event_id),
    CONSTRAINT fk_user_event_waitlists_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_event_waitlists_event FOREIGN key(event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_tags(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    tag_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    PRIMARY KEY(user_id,tag_id),
    CONSTRAINT fk_user_tags_user FOREIGN key(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_tags_tag FOREIGN key(tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS uni_users_email ON users USING btree ("email");

CREATE TABLE IF NOT EXISTS verifications(
    user_id uuid NOT NULL,
    token varchar(255),
    expires_at timestamp with time zone NOT NULL,
    "type" varchar(255) NOT NULL,
    PRIMARY KEY(user_id,expires_at)
);

CREATE UNIQUE INDEX IF NOT EXISTS uni_verifications_token ON verifications USING btree ("token");

CREATE TYPE OAuthResourceType AS ENUM ('google', 'outlook');

CREATE TABLE IF NOT EXISTS user_oauth_tokens(
    user_id uuid NOT NULL,
    refresh_token varchar,
    access_token varchar,
    csrf_token varchar(255),
    resource_type OAuthResourceType,
    expires_at timestamp with time zone NOT NULL,
    PRIMARY KEY(user_id,resource_type)
);

COMMIT;
