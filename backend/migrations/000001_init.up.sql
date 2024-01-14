-- ENTITY TABLES

CREATE TYPE user_role AS ENUM ('super', 'clubAdmin', 'student');

CREATE TYPE college AS ENUM (
    'CAMD', -- College of Arts, Media and Design
    'DMSB', -- D'Amore-McKim School of Business
    'KCCS', -- Khoury College of Computer Sciences
    'CoE', -- College of Engineering
    'BCoHS', -- Bouvé College of Health Sciences
    'SoL', -- School of Law
    'CoPS', -- College of Professional Studies
    'CoS', -- College of Science
    'CoSSH' -- College of Social Sciences and Humanities
    );

CREATE TABLE IF NOT EXISTS sac_user (
    id            UUID         NOT NULL PRIMARY KEY,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    role          user_role    NOT NULL,
    nuid          VARCHAR(9)   NOT NULL UNIQUE,
    first_name    VARCHAR(255) NOT NULL,
    last_name     VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT         NOT NULL,
    dob           DATE         NOT NULL,
    college       college      NOT NULL
);

CREATE TABLE IF NOT EXISTS interest (
    id         UUID         NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    title      VARCHAR(255) NOT NULL,
    icon       CHAR         NOT NULL
);


CREATE TYPE category_name AS ENUM ('academic', 'arts', 'business', 'cultural', 'health', 'political', 'professional', 'religious', 'social', 'sports', 'technology', 'other');

CREATE TABLE IF NOT EXISTS category
(
    id         UUID          NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ   NOT NULL DEFAULT NOW(),

    name       category_name NOT NULL
);

CREATE TYPE recruitment_cycle AS ENUM ('fall', 'spring', 'fallSpring', 'always');
CREATE TYPE recruitment_type AS ENUM ('accepting', 'tryout', 'application');

CREATE TABLE IF NOT EXISTS club (
    id                UUID              NOT NULL PRIMARY KEY,
    created_at        TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
    
    soft_deleted_at   TIMESTAMPTZ,

    parent            UUID REFERENCES club (id),

    name              VARCHAR(255)      NOT NULL,
    preview           VARCHAR(255)      NOT NULL,
    description       VARCHAR(255)      NOT NULL, -- MongoDB URI
    num_members       INT               NOT NULL, -- on sac_user join, increment this field in transaction for efficient num_members query
    is_recruiting     BOOLEAN           NOT NULL,
    recruitment_cycle recruitment_cycle NOT NULL,
    recruitment_type  recruitment_type  NOT NULL,
    application_link  VARCHAR(255),
    logo              VARCHAR(255)      NOT NULL -- S3 URI

    CONSTRAINT fk_parent
        FOREIGN KEY (parent)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS point_of_contact (
    id         UUID         NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    profile    VARCHAR(255), -- S3 URI, fallback to default logo if null
    position   VARCHAR(255) NOT NULL,

    club_id    UUID         NOT NULL REFERENCES club (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TYPE media AS ENUM ('facebook', 'instagram', 'twitter', 'linkedin', 'youtube', 'github', 'custom');

CREATE TABLE IF NOT EXISTS contact (
    id         UUID         NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    type       media        NOT NULL,
    content    VARCHAR(255) NOT NULL, -- media URI

    club_id    UUID         NOT NULL REFERENCES club (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tag (
    id         UUID         NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    name       VARCHAR(255) NOT NULL
);

CREATE TYPE event_type AS ENUM ('open', 'membersOnly');

CREATE TABLE IF NOT EXISTS event (
    id          UUID         NOT NULL PRIMARY KEY,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    name        VARCHAR(255) NOT NULL,
    preview     VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    start_time  TIMESTAMPTZ  NOT NULL,
    end_time    TIMESTAMPTZ  NOT NULL,
    location    VARCHAR(255) NOT NULL,
    type        event_type   NOT NULL
);

CREATE TABLE IF NOT EXISTS notification (
    id         UUID         NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    send_at    TIMESTAMPTZ  NOT NULL,

    title      VARCHAR(255) NOT NULL,
    content    VARCHAR(255) NOT NULL,
    deep_link  VARCHAR(255) NOT NULL,
    icon       VARCHAR(255) NOT NULL -- S3 URI
);

-- BRIDGE TABLES

-- sac_user Profile Bridges
CREATE TABLE IF NOT EXISTS sac_user_interest (
    id          UUID        NOT NULL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id     UUID        NOT NULL REFERENCES sac_user (id),
    interest_id UUID        NOT NULL REFERENCES interest (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_interest
        FOREIGN KEY (interest_id)
            REFERENCES interest (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sac_user_category (
    id          UUID        NOT NULL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id     UUID        NOT NULL REFERENCES sac_user (id),
    category_id UUID        NOT NULL REFERENCES category (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
            REFERENCES category (id)
            ON DELETE CASCADE
);

-- sac_user <-> Club Bridges

CREATE TABLE IF NOT EXISTS membership (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id    UUID        NOT NULL REFERENCES sac_user (id),
    club_id    UUID        NOT NULL REFERENCES club (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS intended_applicant (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id    UUID        NOT NULL REFERENCES sac_user (id),
    club_id    UUID        NOT NULL REFERENCES club (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS follower(
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id    UUID        NOT NULL REFERENCES sac_user (id),
    club_id    UUID        NOT NULL REFERENCES club (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS faq (
    id                UUID         NOT NULL PRIMARY KEY,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    created_by        UUID         NOT NULL REFERENCES sac_user (id),
    question          VARCHAR(255) NOT NULL,

    answered_by       UUID REFERENCES sac_user (id),
    answer            VARCHAR(255),
    num_found_helpful INT,

    club_id           UUID         NOT NULL REFERENCES club (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE
);

-- Club <-> Tag Bridge

CREATE TABLE IF NOT EXISTS club_tag (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    club_id    UUID        NOT NULL REFERENCES club (id),
    tag_id     UUID        NOT NULL REFERENCES tag (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_tag
        FOREIGN KEY (tag_id)
            REFERENCES tag (id)
            ON DELETE CASCADE
);

-- Club <-> Event Bridge

CREATE TABLE IF NOT EXISTS club_event (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    club_id    UUID        NOT NULL REFERENCES club (id),
    event_id   UUID        NOT NULL REFERENCES event (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
            REFERENCES event (id)
            ON DELETE CASCADE
);

-- Event <-> Tag Bridge

CREATE TABLE IF NOT EXISTS event_tag (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    event_id   UUID        NOT NULL REFERENCES event (id),
    tag_id     UUID        NOT NULL REFERENCES tag (id),

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
            REFERENCES event (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_tag
        FOREIGN KEY (tag_id)
            REFERENCES tag (id)
            ON DELETE CASCADE
);

-- Event <-> sac_user Bridges

CREATE TABLE IF NOT EXISTS event_waitlist (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id    UUID        NOT NULL REFERENCES sac_user (id),
    event_id   UUID        NOT NULL REFERENCES event (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
            REFERENCES event (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS event_rsvp (
    id         UUID        NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sac_user_id    UUID        NOT NULL REFERENCES sac_user (id),
    event_id   UUID        NOT NULL REFERENCES event (id),

    CONSTRAINT fk_sac_user
        FOREIGN KEY (sac_user_id)
            REFERENCES sac_user (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
            REFERENCES event (id)
            ON DELETE CASCADE
);

-- Notifcation Bridges

CREATE TABLE IF NOT EXISTS club_notification (
    id              UUID        NOT NULL PRIMARY KEY,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    club_id         UUID        NOT NULL REFERENCES club (id),
    notification_id UUID        NOT NULL REFERENCES notification (id),

    CONSTRAINT fk_club
        FOREIGN KEY (club_id)
            REFERENCES club (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_notification
        FOREIGN KEY (notification_id)
            REFERENCES notification (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS event_notification (
    id              UUID        NOT NULL PRIMARY KEY,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    event_id        UUID        NOT NULL REFERENCES event (id),
    notification_id UUID        NOT NULL REFERENCES notification (id),

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
            REFERENCES event (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_notification
        FOREIGN KEY (notification_id)
            REFERENCES notification (id)
            ON DELETE CASCADE
);