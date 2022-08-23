CREATE TABLE IF NOT EXISTS public.users
(
    id bigserial PRIMARY KEY,
    name character varying(40) NOT NULL,
    email character varying(40) NOT NULL UNIQUE,
    passhash character (60) NOT NULL,
    role_id smallint NOT NULL,
    created_date timestamp NOT NULL DEFAULT NOW()::timestamp,
    deleted_date timestamp DEFAULT NULL
    );