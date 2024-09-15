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



CREATE TABLE IF NOT EXISTS public.users;
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    mail character varying(255) COLLATE pg_catalog."default" NOT NULL,
    pass character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
                             CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT email_uq_cnsts UNIQUE (mail)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;