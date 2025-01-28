-- Migration: create_uuid_ossp_extension (UP)
-- Created at: 2025-01-28 23:03:31

BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

COMMIT;
