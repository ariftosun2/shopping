-- Table: public.books_item
-- DROP TABLE public.books_item;
CREATE TABLE IF NOT EXISTS public.books_item (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    bookskind character varying(100) COLLATE pg_catalog."default" NOT NULL,
    bookname character varying(100) COLLATE pg_catalog."default" NOT NULL,
    detail text COLLATE pg_catalog."default" NOT NULL,
    created date,
    CONSTRAINT news_item_pkey PRIMARY KEY (id)
) TABLESPACE pg_default;
ALTER TABLE public.books_item OWNER to postgres;