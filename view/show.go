package view

import (
	"fmt"
	"memo/model"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	// config?
	maxMemoLength   = 50
	splitMemoLength = 30
)

func renderMemos(memos []model.Memo, reveal bool) {
	t := table.NewWriter()
	t.Style().Format.Header = text.FormatDefault
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Memo"})
	for _, v := range memos {
		var content string
		if reveal {
			content = v.Content
		} else {
			if v.ShortedContent != "" {
				content = v.ShortedContent
			} else {
				content = v.Content
			}
		}
		if len(content) > maxMemoLength {
			contentSlice := strings.Split(content, "")
			first := contentSlice[:splitMemoLength]
			second := contentSlice[splitMemoLength:]
			content = fmt.Sprintf("%s\n%s", strings.Join(first, ""), strings.Join(second, ""))
		}
		t.AppendRow(table.Row{v.Id, content})
		t.AppendSeparator()
	}
	t.Render()
}

func ShowMemo(memos []model.Memo, reveal bool) {
	if len(memos) == 0 {
		fmt.Printf("\n No Memo \n\n")
	} else {
		renderMemos(memos, reveal)
	}
}

func InsertMemo(row int) {
	fmt.Printf("\n Added %v memo \n\n", row)
}

func DeleteMemoById(count int) {
	fmt.Printf("\n Deleted %v memo \n\n", count)
}

func DeleteAllMemo(count int) {
	fmt.Printf("\n Deleted %v memo \n\n", count)
}
func UpdateMemo(memoId int) {
	fmt.Printf("\n Memo with id %v updated \n\n", memoId)
}

func ArchiveMemo(count int) {
	fmt.Printf("\n Archived %d memo \n\n", count)
}
