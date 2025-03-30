-- Migration: drop_users_table_constraint (UP)
-- Created at: 2025-03-30 20:30:59

BEGIN;

ALTER TABLE users
ALTER COLUMN username DROP NOT NULL,
ALTER COLUMN phone_number DROP NOT NULL;

COMMIT;
