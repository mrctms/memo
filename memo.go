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
	add         = flag.Bool("a", false, "[memo] | To add a memo")
	addShort    = flag.Bool("ash", false, "[long memo] [shorted memo] | Add a shorted memo")
	show        = flag.Bool("s", false, "To show all memo")
	delete      = flag.Bool("d", false, "[position number] | To delete a memo")
	deleteAll   = flag.Bool("da", false, "To delete all memo")
	reveal      = flag.Bool("r", false, "[position number] | Show the complete memo")
	modify      = flag.Bool("m", false, "[position number] [memo] | To edit a memo")
	modifyShort = flag.Bool("msh", false, "[position number] [memo] | To edit the memo behind the shorted memo")
	db, err     = sql.Open("sqlite3", "./memo.db")
)

func createMemo() {
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text, Short text, DateTime text)")
}

func insertShort(argsString string, shortString string) {
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("INSERT INTO Things (ToDo, Short, DateTime) VALUES (?, ?, ?)", (shortString), (argsString), (date))
	defer db.Close()
}

func insertMemo(argsString string) {
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("INSERT INTO Things (ToDo, DateTime) VALUES (?, ?)", (argsString), (date))
	defer db.Close()
}

func selectShortMemo(argsRowid string) {
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT rowid, Short, DateTime FROM Things WHERE rowid=?", (argsRowid))
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
	defer rows.Close()
	defer db.Close()
}

func selectMemo() {
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
	defer rows.Close()
	defer db.Close()
}

func modifyMemo(argsStringR string, argsStringM string) {
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE rowid=?", (argsStringM), (date), (argsStringR))
	defer db.Close()
}

func modifyMemoShort(argsStringR string, argsStringS string) {
	if err != nil {
		log.Fatal(err)
	}
	var t = time.Now()
	var date = t.Format("2006-01-02 15:04:05")
	db.Exec("UPDATE Things SET Short=?, DateTime=? WHERE rowid=?", (argsStringS), (date), (argsStringR))
	defer db.Close()
}

func deleteMemo(argsInt []string) {
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range argsInt {
		db.Exec("DELETE FROM Things WHERE rowid=?", (v))
	}
	defer db.Close()
}

func deleteAllMemo() {
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("DELETE FROM Things")
	defer db.Close()
}

func getUserHome() {
	var home, _ = user.Current()
	os.Chdir(home.HomeDir)
	os.Mkdir(".memo", 0700)
	os.Chdir(".memo")
	createMemo()
}

func main() {
	getUserHome()
	flag.Parse()
	if *add {
		insertMemo(os.Args[2])
	} else if *delete {
		deleteMemo(os.Args[1:])
	} else if *deleteAll {
		deleteAllMemo()
	} else if *show {
		selectMemo()
	} else if *addShort {
		insertShort(os.Args[2], os.Args[3])
	} else if *reveal {
		selectShortMemo(os.Args[2])
	} else if *modify {
		modifyMemo(os.Args[2], os.Args[3])
	} else if *modifyShort {
		modifyMemoShort(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Something went wrong")
	}
}
