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

package controller

import (
	"log"
	"memo/model"
	"memo/view"
)

type Executor struct {
	dbModel *model.MemoDb
}

func NewExecutor(dbPath string) *Executor {
	executor := new(Executor)
	dbModel := model.NewMemoDb(dbPath)
	executor.dbModel = dbModel
	executor.initDb()
	return executor
}

func (e *Executor) initDb() {
	e.dbModel.Execute("CREATE TABLE IF NOT EXISTS Memo (Id integer primary key autoincrement, Content text, ShortContent text, Date text)")
}
func (e *Executor) CreateMemo(content string, shortContent string, date string) *model.Memo {
	memo := &model.Memo{
		Content:        content,
		ShortedContent: shortContent,
		Date:           date,
	}
	return memo
}

func (e *Executor) GetMemo() {
	memos, err := e.dbModel.Query("SELECT * FROM Memo")
	if err != nil {
		log.Fatalln(err)
	}
	view.ShowMemo(memos, false)
}

func (e *Executor) GetMemoById(id int) {
	memo, err := e.dbModel.Query("SELECT * FROM Memo WHERE Id=?", (id))
	if err != nil {
		log.Fatalln(err)
	}
	view.ShowMemo(memo, true)
}
func (e *Executor) DeleteAllMemo() {
	result, err := e.dbModel.Execute("DELETE FROM Memo")
	if err != nil {
		log.Fatalln(err)
	}
	view.DeleteAllMemo(result)
}

func (e *Executor) DeleteMemoById(id []int) {
	var total int
	for _, v := range id {
		result, err := e.dbModel.Execute("DELETE FROM Memo WHERE Id=?", (v))
		if err != nil {
			log.Fatalln(err)
		}
		total += result
	}
	view.DeleteMemoById(total)
}

func (e *Executor) InsertMemo(memo *model.Memo) {
	result, err := e.dbModel.Execute("INSERT INTO Memo (Content,ShortContent,Date) VALUES (?, ?, ?)", (memo.Content), (memo.ShortedContent), (memo.Date))
	if err != nil {
		log.Fatalln(err)
	}
	view.InsertMemo(result)
}

func (e *Executor) UpdateMemo(memo *model.Memo, id int) {
	var err error
	if memo.ShortedContent == "" {
		_, err = e.dbModel.Execute("UPDATE Memo SET Content=?, Date=? WHERE Id=?", (memo.Content), (memo.Date), (id))
	} else if memo.Content == "" {
		_, err = e.dbModel.Execute("UPDATE Memo SET ShortContent=?, Date=? WHERE Id=?", (memo.ShortedContent), (memo.Date), (id))
	} else {
		_, err = e.dbModel.Execute("UPDATE Memo SET Content=?, ShortContent=?, Date=? WHERE ID=?", (memo.Content), (memo.ShortedContent), (memo.Date), (id))
	}
	if err != nil {
		log.Fatalln(err)
	}
	view.UpdateMemo(id)
}
