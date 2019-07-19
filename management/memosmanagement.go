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
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//CreateMemo create a memo
func CreateMemo(memo *Memo) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO Things (ID, ToDo, DateTime, Short) VALUES (?, ?, ?, ?)", (memo.ID), (memo.Content), (memo.Date), (memo.ShortedContent))
	defer db.Close()
}

//CreateShortedMemo create a shorted memo
func CreateShortedMemo(memo *Memo) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO Things (ID, ToDo, DateTime, Short) VALUES (?, ?, ?, ?)", (memo.ID), (memo.ShortedContent), (memo.Date), (memo.Content))
	defer db.Close()
}

//GetUserHome will set the location where create the memo database.
//Create the folder and the database
func GetUserHome() {
	var home, e = os.UserHomeDir()
	if e != nil {
		log.Fatal(e)
	}
	os.Chdir(home)
	os.Mkdir(".memo", 0700)
	os.Chdir(".memo")
	createMemoTable()
}

func createMemoTable() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS Things (ID int, ToDo text, Short text, DateTime text)")
	defer db.Close()
}

//DeleteMemos delete the memo by id
func DeleteMemos(id []string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range id {
		db.Exec("DELETE FROM Things WHERE ID=?", (v))
	}
	defer db.Close()
}

//DeleteAllMemos delete all memos
func DeleteAllMemos() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("DELETE FROM Things")
	defer db.Close()
}

//SelectMemo will print all memos
func SelectMemo() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT ID, ToDo, DateTime FROM Things")
	if e != nil {
		log.Fatal(e)
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
	defer db.Close()
}

//SelectShortMemo will print shorted memo by id
func SelectShortMemo(id string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT ID, Short, DateTime FROM Things WHERE ID=?", (id))
	if e != nil {
		log.Fatal(e)
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
	defer db.Close()
}

//ModifyMemo edit the memo by id
func ModifyMemo(id string, newMemo *Memo) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE ID=?", (newMemo.Content), (newMemo.Date), (id))
	defer db.Close()
}

//ModifyMemoShort edit the shorted memo by id
func ModifyMemoShort(id string, newShortedMemo *Memo) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("UPDATE Things SET Short=?, DateTime=? WHERE ID=?", (newShortedMemo.ShortedContent), (newShortedMemo.Date), (id))
	defer db.Close()
}
