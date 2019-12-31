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
	"log"
	"os"
	"path"

	"github.com/marcktomack/memo/cmd"
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

	cmd.Execute(db)
}
