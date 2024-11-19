package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lem3s/go-for-it/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Long:  `Example: "go-for-it complete <taskid>"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: complete,
}

func complete(cmd *cobra.Command, args []string) {
	targetId, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	taskList := utils.ReadFromCsv()

	for i, task := range taskList {
		if task.Id == targetId {
			taskList[i].IsDone = true

			break
		}

		if i == len(taskList)-1 {
			fmt.Println("Task not found")

			return
		}
	}

	utils.WriteToCsv(taskList)
}
