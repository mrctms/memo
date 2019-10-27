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
	"path"

	m "github.com/marcktomack/memo/management"
)

var (
	add         = flag.Bool("a", false, "[memo] | To add a memo")
	addShort    = flag.Bool("ash", false, "[long memo] [shorted memo] | Add a shorted memo")
	show        = flag.Bool("s", false, "To show all memo")
	delete      = flag.Bool("d", false, "[id] | To delete a memo")
	deleteAll   = flag.Bool("da", false, "To delete all memo")
	reveal      = flag.Bool("r", false, "[id] | Show the complete memo")
	modify      = flag.Bool("m", false, "[id] [new memo] | To edit a memo")
	modifyShort = flag.Bool("msh", false, "[id] [new shorted memo] | To edit the memo behind the shorted memo")
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	memoDb := path.Join(home, ".memo", "memo.db")
	db, err := sql.Open("sqlite3", memoDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m.CreateMemoTable(db)
	flag.Parse()

	if *add {
		var newMemo = m.NewMemo(os.Args[2], "")
		m.CreateMemo(db, newMemo)
	} else if *delete {
		m.DeleteMemos(db, os.Args[1:])
	} else if *deleteAll {
		m.DeleteAllMemos(db)
	} else if *show {
		m.SelectMemo(db)
	} else if *addShort {
		var newShortedMemo = m.NewMemo(os.Args[2], os.Args[3])
		m.CreateShortedMemo(db, newShortedMemo)
	} else if *reveal {
		m.SelectShortMemo(db, os.Args[2])
	} else if *modify {
		var newMemo = m.NewMemo(os.Args[3], "")
		m.ModifyMemo(db, os.Args[2], newMemo)
	} else if *modifyShort {
		var newShortedMemo = m.NewMemo("", os.Args[3])
		m.ModifyMemoShort(db, os.Args[2], newShortedMemo)
	} else {
		fmt.Println("Something went wrong")
	}
}
