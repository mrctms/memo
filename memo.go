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

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Add         = flag.Bool("a", false, "[memo] | To add a memo")
	AddShort    = flag.Bool("ash", false, "[long memo] [shorted memo] | Add a shorted memo")
	Show        = flag.Bool("s", false, "To show all memo")
	Delete      = flag.Bool("d", false, "[position number] | To delete a memo")
	DeleteAll   = flag.Bool("da", false, "To delete all memo")
	Reveal      = flag.Bool("r", false, "[position number] | Show the complete memo")
	Modify      = flag.Bool("m", false, "[position number] [memo] | To edit a memo")
	ModifyShort = flag.Bool("msh", false, "[position number] [memo] | To edit the memo behind the shorted memo")
)

func CreateMemo() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text, Short text, DateTime text)")
}

func InsertShort(ArgsString string, ShortString string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("INSERT INTO Things (ToDo, Short, DateTime) VALUES (?, ?, ?)", (ShortString), (ArgsString), (date))
	defer db.Close()
}

func InsertMemo(ArgsString string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("INSERT INTO Things (ToDo, DateTime) VALUES (?, ?)", (ArgsString), (date))
	defer db.Close()
}

func SelectShortMemo(ArgsRowid string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT rowid, Short, DateTime FROM Things WHERE rowid=?", (ArgsRowid))
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("\n Memo:\n")
	for rows.Next() {
		var Short string
		var rowid int
		var DateTime string
		rows.Scan(&rowid, &Short, &DateTime)
		fmt.Println("\n", rowid, "-", DateTime, "-", Short)
	}
	fmt.Printf("\n")
	defer rows.Close()
	defer db.Close()
}

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
		var ToDo string
		var rowid int
		var DateTime string
		rows.Scan(&rowid, &ToDo, &DateTime)
		fmt.Println("\n", rowid, "-", DateTime, "-", ToDo)
	}
	fmt.Printf("\n")
	defer rows.Close()
	defer db.Close()
}

func ModifyMemo(ArgsStringR string, ArgsStringM string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE rowid=?", (ArgsStringM), (date), (ArgsStringR))
	defer db.Close()
}

func ModifyMemoShort(ArgsStringR string, ArgsStringS string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("UPDATE Things SET Short=?, DateTime=? WHERE rowid=?", (ArgsStringS), (date), (ArgsStringR))
	defer db.Close()
}

func DeleteMemo(ArgsInt []string) {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ArgsInt {
		db.Exec("DELETE FROM Things WHERE rowid=?", (v))
	}
	defer db.Close()
}

func DeleteAllMemo() {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("DELETE FROM Things")
	defer db.Close()
}

func GetUserHome() {
	var Home, _ = user.Current()
	os.Chdir(Home.HomeDir)
	os.Mkdir(".memo", 0700)
	os.Chdir(".memo")
	CreateMemo()
}

func main() {
	GetUserHome()
	flag.Parse()
	if *Add {
		InsertMemo(os.Args[2])
	} else if *Delete {
		DeleteMemo(os.Args[1:])
	} else if *DeleteAll {
		DeleteAllMemo()
	} else if *Show {
		SelectMemo()
	} else if *AddShort {
		InsertShort(os.Args[2], os.Args[3])
	} else if *Reveal {
		SelectShortMemo(os.Args[2])
	} else if *Modify {
		ModifyMemo(os.Args[2], os.Args[3])
	} else if *ModifyShort {
		ModifyMemoShort(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Something went wrong")
	}
}
