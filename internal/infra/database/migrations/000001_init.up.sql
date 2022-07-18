CREATE TABLE IF NOT EXISTS public.coordinate
(
    id      serial
        constraint coordinate_pk
            primary key,
    mt      integer          not null,
    axis    varchar          not null,
    horizon varchar          not null,
    x       double precision not null,
    y       double precision not null
);

alter table coordinate
    owner to postgres;

create unique index coordinate_id_uindex
    on coordinate (id);