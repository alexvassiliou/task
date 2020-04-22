package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alexvassiliou/gophercises/task/db"
	"github.com/spf13/cobra"
)

var task string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if task != "" {
			db.CreateTask(task)
		} else {
			err := userInput()
			if err != nil {
				log.Fatal(err)
			}
			db.CreateTask(task)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&task, "create-task", "c", "", "enter the task directly")
}

func userInput() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your task:")
	var err error
	task, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	return nil
}
