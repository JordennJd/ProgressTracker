-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS progress_tracker;

CREATE TABLE IF NOT EXISTS progress_tracker.job (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id UUID NOT NULL,
    status int NOT NULL,
    progress real NOT NULL,
    job_type varchar(100) NOT NULL,
    message varchar(300),
    context varchar(1000),
    created_at TIMESTAMPTZ DEFAULT now(),
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    user_id UUID NOT NULL
    );