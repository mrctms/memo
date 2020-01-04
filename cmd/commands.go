/*Copyright (C) MarckTomack <marcktomack@tutanota.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.*/

package cmd

import (
	"database/sql"
	"github.com/marcktomack/memo/app"
	m "github.com/marcktomack/memo/management"
	"github.com/spf13/cobra"
)

var memoApp *app.App
var importFileFlag *bool
var shortFlag *bool
var deleteAllFlag *bool
var showShortedFlag *bool
var editShortedFlag *bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a memo",
	Long:  "Add a memo or a shorted memo with --short flag",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *m.Memo
		if *importFileFlag {
			memoApp.ImportFromFile(args[0])
		} else {
			if *shortFlag {
				newMemo = app.NewMemo(args[0], args[1])
			} else {
				newMemo = app.NewMemo(args[0], "")
			}
			memoApp.CreateMemo(newMemo)
		}

	},
	Example: "memo add \"your memo\"" + "\n" + "memo add \"your long memo\" --short \"your short memo\"",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a memo",
	Long:  "Delete a memo or delete all memo with --all flag",
	Run: func(cmd *cobra.Command, args []string) {
		if *deleteAllFlag {
			memoApp.DeleteAllMemos()
		} else {
			memoApp.DeleteMemos(args[0:])
		}
	},
	Example: "memo delete 1\n" + "memo delete -all",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all memo",
	Long:  "Show all memo or reveal the memo behind a shorted memo with --r flag",
	Run: func(cmd *cobra.Command, args []string) {
		if *showShortedFlag {
			memoApp.SelectShortMemo(args[0])
		} else {
			memoApp.SelectMemo()
		}
	},
	Example: "memo show\n" + "memo show --r 1",
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a memo",
	Long:  "Edit a memo use --short to edit also the shorted memo",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *m.Memo
		if *editShortedFlag {
			if len(args) == 3 {
				newMemo = app.NewMemo(args[1], args[2])
			} else if len(args) == 2 {
				newMemo = app.NewMemo("", args[2])
			}
		} else {
			newMemo = app.NewMemo(args[1], "")
		}
		memoApp.EditMemo(args[0], newMemo)
	},
	Example: "memo edit 1 \"new memo\"\n" + "memo edit 1 \"new memo\" --short \"new shorted memo\"\n" + "memo edit 1 --short \"new shorted memo\"",
}

func Execute(db *sql.DB) {
	memoApp = app.Initialize(db)
	rootCmd := &cobra.Command{Use: "memo"}
	importFileFlag = addCmd.Flags().Bool("file", false, "")
	shortFlag = addCmd.Flags().Bool("short", false, "")
	deleteAllFlag = deleteCmd.Flags().Bool("all", false, "")
	showShortedFlag = showCmd.Flags().Bool("r", false, "")
	editShortedFlag = editCmd.Flags().Bool("short", false, "")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.Execute()
}
