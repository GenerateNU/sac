BEGIN;

-- drop the tables with foreign key dependencies first

-- tables with foreign key constraints
DROP TABLE IF EXISTS user_tags;
DROP TABLE IF EXISTS user_event_waitlists;
DROP TABLE IF EXISTS user_event_rsvps;
DROP TABLE IF EXISTS user_club_members;
DROP TABLE IF EXISTS user_club_intended_applicants;
DROP TABLE IF EXISTS user_club_followers;
DROP TABLE IF EXISTS leadership;
DROP TABLE IF EXISTS event_tags;
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS club_tags;
DROP TABLE IF EXISTS club_events;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS series;
DROP TABLE IF EXISTS clubs;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS verifications;
DROP TABLE IF EXISTS user_oauth_tokens;

-- drop the indexes
DROP INDEX IF EXISTS idx_files_owner_type;
DROP INDEX IF EXISTS idx_files_owner_id;
DROP INDEX IF EXISTS uni_users_email;
DROP INDEX IF EXISTS uni_verifications_token;
DROP INDEX IF EXISTS idx_user_event_waitlists_event_id;
DROP INDEX IF EXISTS idx_user_event_rsvps_event_id;
DROP INDEX IF EXISTS idx_user_club_members_club_id;
DROP INDEX IF EXISTS idx_user_club_intended_applicants_club_id;
DROP INDEX IF EXISTS idx_user_club_followers_club_id;
DROP INDEX IF EXISTS idx_leadership_email;
DROP INDEX IF EXISTS idx_leadership_club_id;
DROP INDEX IF EXISTS compositeindex;
DROP INDEX IF EXISTS idx_contact_type;
DROP INDEX IF EXISTS idx_club_tags_club_id;
DROP INDEX IF EXISTS idx_club_tags_tag_id;
DROP INDEX IF EXISTS idx_club_events_event_id;
DROP INDEX IF EXISTS idx_club_events_club_id;
DROP INDEX IF EXISTS idx_tags_category_id;
DROP INDEX IF EXISTS uni_categories_name;
DROP INDEX IF EXISTS idx_events_series_id;
DROP INDEX IF EXISTS idx_clubs_one_word_to_describe_us;
DROP INDEX IF EXISTS idx_clubs_num_members;
DROP INDEX IF EXISTS uni_clubs_name;
DROP INDEX IF EXISTS uni_users_email;

-- drop types
DROP TYPE IF EXISTS OAuthResourceType;
DROP TYPE IF EXISTS recruitment;
DROP TYPE IF EXISTS recruitment_type;
DROP TYPE IF EXISTS recruitment_cycle;
DROP TYPE IF EXISTS application;
DROP TYPE IF EXISTS event_type;

-- drop extensions
DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;
