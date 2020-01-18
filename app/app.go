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

package app

import (
	"database/sql"
	"log"

	m "github.com/marcktomack/memo/management"
)

type App struct {
	memoDb m.MemoDB
}

func NewMemo(content string, shortedContent string) *m.Memo {
	return m.NewMemo(content, shortedContent)
}

func Initialize(db *sql.DB) *App {
	app := new(App)
	app.memoDb.DB = db
	err := app.memoDb.CreateMemoTable()
	if err != nil {
		log.Fatalln(err)
	}
	return app
}

func (a *App) CreateMemo(memo *m.Memo) {
	err := a.memoDb.CreateMemo(memo)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *App) DeleteMemos(id []string) {
	err := a.memoDb.DeleteMemos(id)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *App) DeleteAllMemos() {
	err := a.memoDb.DeleteAllMemos()
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *App) SelectMemo() {
	err := a.memoDb.SelectMemo()
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *App) SelectShortMemo(id string) {
	err := a.memoDb.SelectShortMemo(id)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *App) EditMemo(id string, memo *m.Memo) {
	err := a.memoDb.EditMemo(id, memo)
	if err != nil {
		log.Fatalln(err)
	}
}
