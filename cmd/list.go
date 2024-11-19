package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/lem3s/go-for-it/model"
	"github.com/lem3s/go-for-it/utils"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("a", false, "List all the tasks.")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the uncompleted tasks",
	Long:  `Example: go-for-it list"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(0)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	showAllTasks, _ := cmd.Flags().GetBool("a")

	taskList := utils.ReadFromCsv()

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 5, ' ', 0)

	for i, task := range taskList {
		if showAllTasks {
			if i == 0 {
				fmt.Fprintln(w, "ID\tDescription\tCreated\tDone")
			}

			fmt.Fprintln(w, allTasksToString(task))

			continue
		}

		if i == 0 {
			fmt.Fprintln(w, "ID\tDescription\tCreated")
		}

		if !task.IsDone {
			fmt.Fprintln(w, taskToString(task))
		}
	}

	w.Flush()
}

func taskToString(task model.Task) string {
	converted := ""

	converted += strconv.FormatInt(int64(task.Id), 10) + "\t"
	converted += task.Description + "\t"

	_, offsetSeconds := time.Now().Zone()
	hoursDiff := offsetSeconds / 3600

	converted += timediff.TimeDiff(
		task.DateCreated,
		timediff.WithStartTime(time.Now().Add(time.Duration(hoursDiff)*time.Hour)))

	return converted
}

func allTasksToString(task model.Task) string {
	converted := ""

	converted += strconv.FormatInt(int64(task.Id), 10) + "\t"
	converted += task.Description + "\t"

	_, offsetSeconds := time.Now().Zone()
	hoursDiff := offsetSeconds / 3600

	converted += timediff.TimeDiff(
		task.DateCreated,
		timediff.WithStartTime(time.Now().Add(time.Duration(hoursDiff)*time.Hour))) + "\t"
	converted += strconv.FormatBool(task.IsDone)

	return converted
}
