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

package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type MemoDb struct {
	dbConnection *sql.DB
}

func NewMemoDb(dbPath string) *MemoDb {
	memoDb := new(MemoDb)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	memoDb.dbConnection = db
	return memoDb
}

func (m *MemoDb) CloseDB() {
	m.dbConnection.Close()
}

func (m *MemoDb) Query(sqlQuery string, param ...interface{}) ([]Memo, error) {
	var memos []Memo

	var rows *sql.Rows
	var err error

	if param != nil {
		rows, err = m.dbConnection.Query(sqlQuery, param...)
	} else {
		rows, err = m.dbConnection.Query(sqlQuery)
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var memo Memo
		err := rows.Scan(&memo.Id, &memo.Content, &memo.ShortedContent, &memo.Date)
		if err != nil {
			log.Fatalln(err)
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

func (m *MemoDb) Execute(sqlQuery string, param ...interface{}) (int, error) {
	var result sql.Result
	var err error

	if param != nil {
		result, err = m.dbConnection.Exec(sqlQuery, param...)
	} else {
		result, err = m.dbConnection.Exec(sqlQuery)
	}

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}
