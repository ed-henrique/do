package main

import (
	"fmt"
	"godo/internal/actions"
	"godo/internal/config"
	"godo/internal/database"
)

func main() {
  err := config.New()

  if err != nil {
    panic(err)
  }

  db := database.New()
  a, err := actions.New(db, "List files and directories", "ls", []string{"-la", "/home/eduardo"})

  if err != nil {
    panic(err)
  }

  stdout, stderr, err := a.Run()

  if err != nil {
    panic(err)
  }

  fmt.Printf("%s", stdout)
  fmt.Printf("%s", stderr)
}
