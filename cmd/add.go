package cmd

import (
	"fmt"
	"time"

	"github.com/lem3s/go-for-it/model"
	"github.com/lem3s/go-for-it/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your list",
	Long:  `Example: go-for-it add "Clean my desk"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: add,
}

func add(cmd *cobra.Command, args []string) {
	taskToAdd := model.Task{
		Id:          utils.GetLastId() + 1,
		Description: args[0],
		DateCreated: time.Now(),
		IsDone:      false,
	}

	utils.AddTaskToPersistance(taskToAdd)

	fmt.Println("Task added successfully")
}
