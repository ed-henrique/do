CREATE TABLE IF NOT EXISTS ACTION (
  ID INTEGER PRIMARY KEY,
  NAME TEXT NOT NULL,
  COMMAND TEXT NOT NULL,
  OPTIONS TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS I_ACTION ON ACTION (NAME);
