package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func sqlite() *sql.DB {
	sqlite, err := sql.Open("sqlite3", "./foo.db")

	if err != nil {
		log.Fatal(err)
	}

	return sqlite
}

func postgres() *sql.DB {
	postgres, err := sql.Open("pgx", "postgres://postgres:postgres@db:5432/test")
	if err != nil {
		log.Fatal(err)
	}

	return postgres
}

func benchmarkPrepareSqlite(count int) {
	db := sqlite()

	defer db.Close()

	benchmarkInserts(count, db)
}

func benchmarkPreparePostgres(count int) {
	db := postgres()

	defer db.Close()

	benchmarkInserts(count, db)
}

func benchmarkInsertsSqlite(count int) {
	db := sqlite()

	defer db.Close()

	benchmarkInserts(count, db)
}

func benchmarkInsertsPostgres(count int) {
	db := postgres()

	defer db.Close()

	benchmarkInserts(count, db)
}

func benchmarkUpdatesPostgres() {
	db := postgres()

	defer db.Close()

	benchmarkUpdates(db)
}

func benchmarkUpdatesSqlite() {
	db := sqlite()

	defer db.Close()

	benchmarkUpdates(db)
}

func benchmarkSelectSqlite() {
	db := sqlite()

	db.Close()
	db = sqlite()

	benchmarkSelect(db)
}

func benchmarkSelectPostgres() {
	db := postgres()

	defer db.Close()

	benchmarkSelect(db)
}

func benchmarkSelect(db *sql.DB) {

	r := db.QueryRow("SELECT count(*) from test  WHERE status = $1", "test")

	if r.Err() != nil {
		log.Fatal(r.Err())
	}

	_, err := db.Query("SELECT status from test  WHERE status = $1", "wait")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT *  from test WHERE status = $1", "pending")

	if err != nil {
		log.Fatal(err)
	}
}

func benchmarkUpdates(db *sql.DB) {
	_, err := db.Exec("UPDATE test SET status = $1 WHERE status = $2", "test1", "test")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE test SET status = $1 WHERE status = $2", "wait1", "wait")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE test SET status = $1 WHERE status = $2", "pending1", "pending")

	if err != nil {
		log.Fatal(err)
	}
}

func benchmarkInserts(count int, db *sql.DB) {
	sqlStmt := `
	drop table if exists test;
	create table test (
		id integer not null primary key,
	    status text,
		createdAt timestamp default current_timestamp
	);
	`
	_, err := db.Exec(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into test(id, status) values($1, $2)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	statuses := []string{"pending", "test", "wait"}

	for i := 0; i < count; i++ {
		_, err = stmt.Exec(i, statuses[i%3])
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}
