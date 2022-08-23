create table if not exists public.refresh_tokens
(
    id bigserial PRIMARY KEY,
    user_id bigint not null,
    token text not null,
    expire_date timestamp NOT NULL,
    deleted_date timestamp DEFAULT NULL
)