CREATE TABLE news_item(
    id uuid DEFAULT uuid_generate_v4 (),
    bookskind character varying(100) NOT NULL,
    bookname character varying(100) NOT NULL,
    detail text NOT NULL,
    created date,
    CONSTRAINT news_item_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE);