-- Migration: 001_create_users_table
-- Purpose: Introduces the core users table for authentication and identity management.
-- This table will serve as the foundation for role-based access and user sessions.

CREATE TABLE IF NOT EXISTS users (
    -- Unique identifier for each user (generated in application layer)
    id UUID PRIMARY KEY,

    -- User email used for authentication; must be unique across system
    email TEXT NOT NULL UNIQUE,

    -- Securely hashed password (never store plaintext passwords)
    password_hash TEXT NOT NULL,

    -- Timestamp of account creation (defaults to current time)
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
