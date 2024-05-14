package database

import (
  "database/sql"

  _ "modernc.org/sqlite"
)

func New() *sql.DB {
  db, err := sql.Open("sqlite", "/home/eduardo/.godo/godo.db")

  if err != nil {
    panic(err)
  }

  err = db.Ping()

  if err != nil {
    panic(err)
  }

  _, err = db.Exec(`
CREATE TABLE IF NOT EXISTS ACTION (
  ID INTEGER PRIMARY KEY,
  NAME TEXT NOT NULL,
  COMMAND TEXT NOT NULL,
  OPTIONS TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS I_ACTION ON ACTION (NAME);
`)

  if err != nil {
    panic(err)
  }

  return db
}
