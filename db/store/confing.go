package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "shopping"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if err = createTables(db); err != nil {
		return nil
	}

	return db
}
func createTables(db *sql.DB) error {

	createTables := `
	CREATE TABLE IF NOT EXISTS public.books_item (
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		bookskind character varying(100) COLLATE pg_catalog."default" NOT NULL,
		bookname  character varying(100) COLLATE pg_catalog."default" NOT NULL,
		detail text COLLATE pg_catalog."default" NOT NULL,
		created date,
		CONSTRAINT books_item_pkey PRIMARY KEY (id)
	) TABLESPACE pg_default;
	
	ALTER TABLE public.books_item OWNER to postgres;
	
	CREATE TABLE IF NOT EXISTS public.user_item (
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		username character varying(100) COLLATE pg_catalog."default" NOT NULL,
		lastname character varying(100) COLLATE pg_catalog."default" NOT NULL,
		userpassword character varying(100) COLLATE pg_catalog."default" NOT NULL,
		created date,
		CONSTRAINT user_item_pkey PRIMARY KEY (id)
	) TABLESPACE pg_default;
	
	ALTER TABLE public.user_item OWNER to postgres;
`
	_, err := db.Exec(createTables)
	return err
}
