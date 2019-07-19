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
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var date = time.Now().Format("2006-01-02 15:04:05")

//Memo object
type Memo struct {
	ID             int
	Date           string
	Content        string
	ShortedContent string
}

//NewMemo return a *Memo with her content
func NewMemo(content string, shortedContent string) *Memo {
	var id = queryLastID() + 1
	return &Memo{ID: id, Date: date, Content: content, ShortedContent: shortedContent}
}

func queryLastID() int {
	var db, err = sql.Open("sqlite3", "./memo.db")
	if err != nil {
		log.Fatal(err)
	}
	var rows, e = db.Query("SELECT MAX(ID) FROM Things")
	if e != nil {
		log.Fatal(e)
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}
