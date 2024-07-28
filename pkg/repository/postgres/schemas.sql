CREATE TABLE IF NOT EXISTS public.snippets
(
    id integer NOT NULL DEFAULT nextval('snippets_id_seq'::regclass),
    title character varying(50) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    created timestamp with time zone NOT NULL,
    expires timestamp without time zone NOT NULL,
    CONSTRAINT snippets_pkey PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.snippets
    OWNER to postgres;