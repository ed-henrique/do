package actions

import (
	"bytes"
	"database/sql"
	"os/exec"
	"strings"
)

type Action struct {
  Name string
  Command string
  Options []string
}

func New(db *sql.DB, name, command string, options []string) (Action, error) {
  _, err := db.Exec(
    "INSERT INTO ACTION (NAME, COMMAND, OPTIONS) VALUES (?, ?, ?)",
    name,
    command,
    strings.Join(options, " "),
  )

  if err != nil {
    return Action{}, err
  }

  return Action{ Name: name, Command: command, Options: options }, nil
}

func Get(db *sql.DB, name string) (Action, error) {
  row := db.QueryRow("SELECT COMMAND, OPTIONS FROM ACTION WHERE NAME = ?", name)

  a := Action{}
  if err := row.Err(); err != nil {
    return a, err
  }

  var options string
  if err := row.Scan(&a.Command, &options); err != nil {
    return a, err
  }

  a.Options = strings.Split(options, " ")
  
  return a, nil
}

func (a Action) Run() ([]byte, []byte, error) {
  cmd := exec.Command(a.Command, a.Options...)

  var stdout, stderr bytes.Buffer
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  if err := cmd.Run(); err != nil {
    return nil, nil, err
  }

  return stdout.Bytes(), stderr.Bytes(), nil
}
