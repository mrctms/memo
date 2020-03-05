// Copyright (C) Marck Tomack <marcktomack@tutanota.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"log"
	"memo/controller"
	"memo/model"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var executor *controller.Executor
var shortFlag *bool
var deleteAllFlag *bool
var showShortedFlag *bool
var updateShortedFlag *bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a memo",
	Long:  "Add a memo or a shorted memo with --short flag",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *model.Memo
		if *shortFlag {
			newMemo = executor.CreateMemo(args[0], args[1], time.Now().Format("01-02-2006 15:04:05"))
		} else {
			newMemo = executor.CreateMemo(args[0], "", time.Now().Format("01-02-2006 15:04:05"))
		}
		executor.InsertMemo(newMemo)
	},
	Example: "memo add \"your memo\"" + "\n" + "memo add \"your long memo\" --short \"your short memo\"",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a memo",
	Long:  "Delete a memo or delete all memo with --all flag",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteAllFlag {
			executor.DeleteAllMemo()
		} else {
			var bo []int
			for _, v := range args[0:] {
				res, _ := strconv.Atoi(v)
				bo = append(bo, res)
			}
			executor.DeleteMemoById(bo)
		}
	},
	Example: "memo delete 1\n" + "memo delete --all",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all memo",
	Long:  "Show all memo or reveal the memo behind a shorted memo with --r flag",
	Run: func(cmd *cobra.Command, args []string) {
		if *showShortedFlag {
			res, _ := strconv.Atoi(args[0])
			executor.GetMemoById(res)
		} else {
			executor.GetMemo()
		}
	},
	Example: "memo show\n" + "memo show --r 1",
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a memo",
	Long:  "Update a memo use --short to edit also the shorted memo",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *model.Memo
		if *updateShortedFlag {
			if len(args) == 3 {
				newMemo = executor.CreateMemo(args[1], args[2], time.Now().Format("01-02-2006 15:04:05"))
			} else if len(args) == 2 {
				newMemo = executor.CreateMemo("", args[2], time.Now().Format("01-02-2006 15:04:05"))
			}
		} else {
			newMemo = executor.CreateMemo(args[1], "", time.Now().Format("01-02-2006 15:04:05"))
		}
		res, _ := strconv.Atoi(args[0])
		executor.UpdateMemo(newMemo, res)
	},
	Example: "memo update 1 \"new memo\"\n" + "memo update 1 \"new memo\" --short \"new shorted memo\"\n" + "memo update 1 --short \"new shorted memo\"",
}

func execute() {
	rootCmd := &cobra.Command{Use: "memo"}
	shortFlag = addCmd.Flags().Bool("short", false, "")
	deleteAllFlag = deleteCmd.Flags().Bool("all", false, "")
	showShortedFlag = showCmd.Flags().Bool("r", false, "")
	updateShortedFlag = updateCmd.Flags().Bool("short", false, "")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.Execute()
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	db := path.Join(home, ".memo", "memo.db")
	executor = controller.NewExecutor(db)
	execute()
}
