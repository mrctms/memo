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

package management

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//CreateMemo create a memo
func CreateMemo(db *sql.DB, memo *Memo) {
	db.Exec("INSERT INTO Things (ToDo, DateTime, Short) VALUES (?, ?, ?)", (memo.Content), (memo.Date), (memo.ShortedContent))

}

//CreateShortedMemo create a shorted memo
func CreateShortedMemo(db *sql.DB, memo *Memo) {
	db.Exec("INSERT INTO Things (ToDo, DateTime, Short) VALUES (?, ?, ?)", (memo.ShortedContent), (memo.Date), (memo.Content))
}

func CreateMemoTable(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS Things (ID integer primary key autoincrement, ToDo text, Short text, DateTime text)")
}

//DeleteMemos delete the memo by id
func DeleteMemos(db *sql.DB, id []string) {
	for _, v := range id {
		db.Exec("DELETE FROM Things WHERE ID=?", (v))
	}
}

//DeleteAllMemos delete all memos
func DeleteAllMemos(db *sql.DB) {
	db.Exec("DELETE FROM Things")
}

//SelectMemo will print all memos
func SelectMemo(db *sql.DB) {
	var rows, err = db.Query("SELECT ID, ToDo, DateTime FROM Things")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n Memo:\n")
	for rows.Next() {
		var toDo string
		var id int
		var dateTime string
		rows.Scan(&id, &toDo, &dateTime)
		fmt.Println("\n", id, "-", dateTime, "-", toDo)
	}
	fmt.Printf("\n")
	rows.Close()
}

//SelectShortMemo will print shorted memo by id
func SelectShortMemo(db *sql.DB, id string) {
	var rows, err = db.Query("SELECT ID, Short, DateTime FROM Things WHERE ID=?", (id))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var short string
		var id int
		var dateTime string
		rows.Scan(&id, &short, &dateTime)
		if short == "" {
			fmt.Printf("\nERROR: Memo with ID %d does not have a shorted memo\n\n", id)
		} else {
			fmt.Printf("\n Memo:\n")
			fmt.Println("\n", id, "-", dateTime, "-", short)
			fmt.Printf("\n")
		}
	}
	rows.Close()
}

//ModifyMemo edit the memo by id
func ModifyMemo(db *sql.DB, id string, newMemo *Memo) {
	db.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE ID=?", (newMemo.Content), (newMemo.Date), (id))
}

//ModifyMemoShort edit the shorted memo by id
func ModifyMemoShort(db *sql.DB, id string, newShortedMemo *Memo) {
	db.Exec("UPDATE Things SET Short=?, DateTime=? WHERE ID=?", (newShortedMemo.ShortedContent), (newShortedMemo.Date), (id))
}
