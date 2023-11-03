# Getting Started

# Info

## PostgreSQL Setup

- Download the latest version of [PostgreSQL](https://www.postgresql.org/download/) - Make sure you install pgAdmin within the PostgreSQL config options. Otherwise you can manually install it [here](https://www.pgadmin.org/)
- Configure your local testing database however you want. Just make sure you keep track of the info. `host`, `user`, `password`, etc...

- Copy down all of [our schemas](https://github.com/Minuteman-PWD-2025/Biotech-CSI-Server/edit/psql-implmentation/README.md#database-schemas) and add them to your local test database.

## Database Schemas

### `public.people`
```sql
-- Table: public.people

-- DROP TABLE IF EXISTS public.people;

CREATE TABLE IF NOT EXISTS public.people
(
    id integer NOT NULL DEFAULT nextval('people_id_seq'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT people_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.people
    OWNER to postgres;
```
### `public.people_id_seq`
```sql
-- SEQUENCE: public.people_id_seq

-- DROP SEQUENCE IF EXISTS public.people_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.people_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY people.id;

ALTER SEQUENCE public.people_id_seq
    OWNER TO postgres;
```
