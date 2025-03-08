-- Migration: create_accounts_table (UP)
-- Created at: 2025-01-29 22:24:03

BEGIN;

CREATE TABLE accounts (
    id VARCHAR(255) DEFAULT concat('account-', uuid_generate_v4()) NOT NULL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    balance NUMERIC NOT NULL,
    currency VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    is_archived BOOLEAN NOT NULL,
    is_excluded BOOLEAN NOT NULL,
    mark_for_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMIT;
