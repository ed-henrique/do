package config

import (
	"fmt"
	"os"
	"os/user"
)

type Config struct {
  Path string
}

func New() error {
  currentUser, err := user.Current()

  if err != nil {
    return err
  }

  defaultPath := fmt.Sprintf("/home/%s/.godo", currentUser.Username)
  if _, err := os.Stat(defaultPath); err != nil {
    return os.Mkdir(defaultPath, 0777)
  }

  return nil
}

