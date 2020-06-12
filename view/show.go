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

package view

import (
	"fmt"
	"memo/model"
	"strconv"
	"strings"
)

func formatMemoView(memo model.Memo, reveal bool) string {
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

func ShowMemo(memos []model.Memo, reveal bool) {
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
