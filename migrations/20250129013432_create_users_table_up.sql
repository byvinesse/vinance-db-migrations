-- Migration: create_users_table (UP)
-- Created at: 2025-01-29 01:34:32

BEGIN;

CREATE TABLE users (
    id VARCHAR(255) DEFAULT concat('user-', uuid_generate_v4()) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    date_of_birth TIMESTAMP WITH TIME ZONE,
    gender CHAR,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMIT;
