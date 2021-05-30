package main

import (
	"fmt"
	"strconv"
	"strings"
)

func formatMemoView(memo Memo, reveal bool) string {
	var result strings.Builder
	result.WriteString("\n")
	result.WriteString(" " + strconv.Itoa(memo.Id))
	result.WriteString(" " + "-" + " ")
	result.WriteString(memo.Date)
	result.WriteString(" " + "-" + " ")
	if reveal {
		result.WriteString(memo.Content)
	} else {
		if memo.ShortedContent != "" {
			result.WriteString(memo.ShortedContent)
		} else {
			result.WriteString(memo.Content)
		}
	}
	result.WriteString("\n")

	return result.String()
}

func ShowMemo(memos []Memo, reveal bool) {
	if len(memos) == 0 {
		fmt.Printf("\n No Memo \n\n")
	} else {
		for _, v := range memos {
			result := formatMemoView(v, reveal)
			fmt.Println(result)
		}
	}
}

func InsertMemo(row int) {
	fmt.Printf("\n Added %v memo \n\n", row)
}

func DeleteMemo(count int) {
	fmt.Printf("\n Deleted %v memo \n\n", count)
}
func UpdateMemo(memoId int) {
	fmt.Printf("\n Memo with id %v updated \n\n", memoId)
}

func ArchiveMemo(count int) {
	fmt.Printf("\n Archived %d memo \n\n", count)
}
