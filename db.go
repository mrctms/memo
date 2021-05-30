package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type MemoDb struct {
	dbConnection *sql.DB
	tx           *sql.Tx
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
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

func (m *MemoDb) InitTransaction() error {
	var err error

	m.tx, err = m.dbConnection.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (m *MemoDb) PrepareStatement(sqlQuery string) (*sql.Stmt, error) {
	stmt, err := m.dbConnection.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func (m *MemoDb) ExecuteStatment(stmt *sql.Stmt, params ...interface{}) (int, error) {
	result, err := stmt.Exec(params...)
	if err != nil {
		if err := m.tx.Rollback(); err != nil {
			return 0, fmt.Errorf("error on rollback transaction: %w", err)
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (m *MemoDb) CommitTransaction() error {
	err := m.tx.Commit()
	if err != nil {
		return fmt.Errorf("error on commit transaction: %w", err)
	}
	return nil
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
