package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var executor *Executor
var shortFlag *bool
var deleteAllFlag *bool
var archivedFlag *bool
var ascFlag *bool
var descFlag *bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a memo",
	Long:  "Add a memo or a shorted memo with --short flag",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *Memo
		if *shortFlag {
			newMemo = executor.CreateMemo(args[0], args[1], time.Now().Format("01-02-2006 15:04:05"))
		} else {
			newMemo = executor.CreateMemo(args[0], "", time.Now().Format("01-02-2006 15:04:05"))
		}
		result, err := executor.InsertMemo(newMemo)
		if err != nil {
			fmt.Println(err)
			return
		}
		InsertMemo(result)
	},
	Example: "memo add \"your memo\"" + "\n" + "memo add \"your long memo\" --short \"your short memo\"",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a memo",
	Long:  "Delete a memo or delete all memo with --all flag",
	Run: func(cmd *cobra.Command, args []string) {
		var result int
		var err error
		if *deleteAllFlag {
			if *archivedFlag {
				result, err = executor.DeleteMemoArchive()
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				result, err = executor.DeleteAllMemo()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			var ids []int
			for _, v := range args[0:] {
				res, _ := strconv.Atoi(v)
				ids = append(ids, res)
			}
			result, err = executor.DeleteMemoById(ids)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		DeleteMemo(result)
	},
	Example: "memo delete 1\n" + "memo delete --all",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all memo",
	Long:  "Show all memo or reveal the memo behind a shorted memo with --short flag",
	Run: func(cmd *cobra.Command, args []string) {
		var orderBy OrderBy
		if *ascFlag {
			orderBy = Asc
		} else if *descFlag {
			orderBy = Desc
		}
		var memos []Memo
		var err error

		if *archivedFlag {
			if *shortFlag {
				res, _ := strconv.Atoi(args[0])
				memos, err = executor.GetArchivedMemoById(res)
			} else {
				memos, err = executor.ShowArchivedMemo(orderBy)
			}
		} else {
			if *shortFlag {
				res, _ := strconv.Atoi(args[0])
				memos, err = executor.GetMemoById(res, *shortFlag)
			} else {
				memos, err = executor.GetMemo(orderBy)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
		ShowMemo(memos, *shortFlag)
	},
	Example: "memo show\n" + "memo show --short 1",
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a memo",
	Long:  "Update a memo use --short to edit also the shorted memo",
	Run: func(cmd *cobra.Command, args []string) {
		var newMemo *Memo
		if *shortFlag {
			if len(args) == 3 {
				newMemo = executor.CreateMemo(args[1], args[2], time.Now().Format("01-02-2006 15:04:05"))
			} else if len(args) == 2 {
				newMemo = executor.CreateMemo("", args[2], time.Now().Format("01-02-2006 15:04:05"))
			}
		} else {
			newMemo = executor.CreateMemo(args[1], "", time.Now().Format("01-02-2006 15:04:05"))
		}
		res, _ := strconv.Atoi(args[0])
		err := executor.UpdateMemo(newMemo, res)
		if err != nil {
			fmt.Println(err)
			return
		}
		UpdateMemo(res)
	},
	Example: "memo update 1 \"new memo\"\n" + "memo update 1 \"new memo\" --short \"new shorted memo\"\n" + "memo update 1 --short \"new shorted memo\"",
}

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "archive a memo",
	Long:  "archive a memo",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args[0:] {
			conv, _ := strconv.Atoi(v)
			ids = append(ids, conv)
		}
		res, err := executor.ArchiveMemo(ids)
		if err != nil {
			fmt.Println(err)
			return
		}
		ArchiveMemo(res)
	},
	Example: "memo archive 1",
}

func execute() {
	rootCmd := &cobra.Command{Use: "memo"}
	shortFlag = flag.Bool("short", false, "")
	deleteAllFlag = flag.Bool("all", false, "")
	archivedFlag = flag.Bool("a", false, "")
	ascFlag = flag.Bool("asc", false, "")
	descFlag = flag.Bool("desc", false, "")

	showCmd.Flags().BoolVar(archivedFlag, "a", false, "")
	deleteCmd.Flags().BoolVar(archivedFlag, "a", false, "")
	addCmd.Flags().BoolVar(shortFlag, "short", false, "")
	showCmd.Flags().BoolVar(shortFlag, "short", false, "")
	showCmd.Flags().BoolVar(ascFlag, "asc", false, "")
	showCmd.Flags().BoolVar(descFlag, "desc", false, "")
	updateCmd.Flags().BoolVar(shortFlag, "short", false, "")
	deleteCmd.Flags().BoolVar(deleteAllFlag, "all", false, "")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.Execute()
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	db := path.Join(home, ".memo", "memo.db")
	executor = NewExecutor(db)
	defer executor.Close()
	execute()
}
