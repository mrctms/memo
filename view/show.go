package view

import (
	"fmt"
	"memo/model"
	"strconv"
	"strings"
)

// func main() {
// 	var b strings.Builder

// 	w := "hello world"

// 	b.WriteString("+")
// 	for range w {
// 		b.WriteString("-")
// 	}
// 	b.WriteString("+")
// 	b.WriteString("\n")
// 	b.WriteString("|")
// 	b.WriteString(w)
// 	b.WriteString("|\n")
// 	b.WriteString("+")
// 	for range w {
// 		b.WriteString("-")
// 	}
// 	b.WriteString("+")
// 	fmt.Println(b.String())
// }

func formatMemoView(memos []model.Memo, reveal bool) string {

	var longestMemo string
	var longestId string

	idLabel := "ID"
	memoLabel := "Memo"

	for _, v := range memos {
		id := strconv.Itoa(v.Id)
		if len(v.Content) > len(longestMemo) {
			longestMemo = v.Content
		}
		if len(id) > len(longestId) {
			longestId = id
		}
	}

	var result strings.Builder

	result.WriteString(fmt.Sprintf("+%s+%s+\n|%s", strings.Repeat("-", len(longestId)), strings.Repeat("-", len(longestMemo)), idLabel))

	for i := 0; i < (len(longestId) - len(idLabel)); i++ {
		result.WriteString(" ")
	}
	result.WriteString("|")

	result.WriteString(memoLabel)
	for i := 0; i < (len(longestMemo) - len(memoLabel)); i++ {
		result.WriteString(" ")
	}

	result.WriteString(fmt.Sprintf("|\n+%s+%s+\n", strings.Repeat("-", len(longestId)), strings.Repeat("-", len(longestMemo))))

	// memos

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
		result.WriteString("|")
		id := strconv.Itoa(v.Id)
		result.WriteString(id)
		for i := 0; i < (len(longestId) - len(id)); i++ {
			result.WriteString(" ")
		}
		result.WriteString("|")
		result.WriteString(content)
		for i := 0; i < (len(longestMemo) - len(content)); i++ {
			result.WriteString(" ")
		}
		result.WriteString(fmt.Sprintf("|\n+%s+%s+\n", strings.Repeat("-", len(longestId)), strings.Repeat("-", len(longestMemo))))
	}

	return result.String()
}

func ShowMemo(memos []model.Memo, reveal bool) {
	if len(memos) == 0 {
		fmt.Printf("\n No Memo \n\n")
	} else {
		result := formatMemoView(memos, reveal)
		fmt.Println(result)
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
