package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexvassiliou/gophercises/task/cmd"
	"github.com/alexvassiliou/gophercises/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
