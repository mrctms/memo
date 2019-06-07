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
	"os"
	"os/user"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var date = time.Now().Format("2006-01-02 15:04:05")

//CreateMemoTable create the main table in the database
func CreateMemoTable() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text, Short text, DateTime text)")
	defer db.Close()
}

//CreateShortMemo create a shorted memo
func CreateShortMemo(memo string, shortedMemo string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO Things (ToDo, Short, DateTime) VALUES (?, ?, ?)", (shortedMemo), (memo), (date))
	defer db.Close()
}

//CreateMemo create a memo
func CreateMemo(memo string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO Things (ToDo, DateTime) VALUES (?, ?)", (memo), (date))
	defer db.Close()
}

//SelectShortMemo will print shorted memo by rowid
func SelectShortMemo(rowid string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT rowid, Short, DateTime FROM Things WHERE rowid=?", (rowid))
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("\n Memo:\n")
	for rows.Next() {
		var short string
		var rowid int
		var dateTime string
		rows.Scan(&rowid, &short, &dateTime)
		fmt.Println("\n", rowid, "-", dateTime, "-", short)
	}
	fmt.Printf("\n")
	rows.Close()
	defer db.Close()
}

//SelectMemo will print all memos
func SelectMemo() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT rowid, ToDo, DateTime FROM Things")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("\n Memo:\n")
	for rows.Next() {
		var toDo string
		var rowid int
		var dateTime string
		rows.Scan(&rowid, &toDo, &dateTime)
		fmt.Println("\n", rowid, "-", dateTime, "-", toDo)
	}
	fmt.Printf("\n")
	rows.Close()
	defer db.Close()
}

//ModifyMemo edit the memo by rowid
func ModifyMemo(rowid string, newMemo string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE rowid=?", (newMemo), (date), (rowid))
	defer db.Close()
}

//ModifyMemoShort edit the shorted memo by rowid
func ModifyMemoShort(rowid string, shortedMemo string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("UPDATE Things SET Short=?, DateTime=? WHERE rowid=?", (shortedMemo), (date), (rowid))
	defer db.Close()
}

//DeleteMemos delete the memo by rowid
func DeleteMemos(rowid []string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range rowid {
		db.Exec("DELETE FROM Things WHERE rowid=?", (v))
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

//GetUserHome will set the location where create the memo database.
//Create the folder and the database
func GetUserHome() {
	var home, e = user.Current()
	if e != nil {
		log.Fatal(e)
	}
	os.Chdir(home.HomeDir)
	os.Mkdir(".memo", 0700)
	os.Chdir(".memo")
	CreateMemoTable()
}
