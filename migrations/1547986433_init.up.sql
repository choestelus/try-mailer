CREATE SCHEMA IF NOT EXISTS public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.history (
    id                  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email               text NOT NULL,
    access_from         text,
    created_at          timestamp with time zone DEFAULT now(),
    updated_at          timestamp with time zone DEFAULT now()
);
