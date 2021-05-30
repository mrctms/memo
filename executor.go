package main

import (
	"fmt"
)

type OrderBy int

const (
	Asc  OrderBy = 1
	Desc         = 2
)

var orderByMap = map[OrderBy]string{Asc: "ORDER BY Id ASC", Desc: "ORDER BY Id DESC", 0: "ORDER BY Id ASC"}

type Executor struct {
	dbModel *MemoDb
}

func NewExecutor(dbPath string) *Executor {
	executor := new(Executor)
	dbModel := NewMemoDb(dbPath)
	executor.dbModel = dbModel
	executor.initDb()
	return executor
}

func (e *Executor) Close() {
	e.dbModel.CloseDB()
}

func (e *Executor) initDb() {
	_, err := e.dbModel.Execute("CREATE TABLE IF NOT EXISTS Memo (Id integer primary key autoincrement, Content text, ShortContent text, Date text);CREATE TABLE IF NOT EXISTS MemoArchive (Id integer, Content text, ShortContent text, Date text)")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (e *Executor) CreateMemo(content string, shortContent string, date string) *Memo {
	memo := &Memo{
		Content:        content,
		ShortedContent: shortContent,
		Date:           date,
	}
	return memo
}

func (e *Executor) GetMemo(by OrderBy) ([]Memo, error) {
	memos, err := e.dbModel.Query(fmt.Sprintf("SELECT * FROM Memo %s", orderByMap[by]))
	if err != nil {
		return nil, err
	}
	return memos, nil
}

func (e *Executor) GetMemoById(id int, reveal bool) ([]Memo, error) {
	memo, err := e.dbModel.Query("SELECT * FROM Memo WHERE Id=?", (id))
	if err != nil {
		return nil, err
	}
	return memo, nil
}
func (e *Executor) DeleteAllMemo() (int, error) {
	result, err := e.dbModel.Execute("DELETE FROM Memo")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (e *Executor) DeleteMemoById(id []int) (int, error) {
	var total int
	for _, v := range id {
		result, err := e.dbModel.Execute("DELETE FROM Memo WHERE Id=?", (v))
		if err != nil {
			return total, err
		}
		total += result
	}
	return total, nil
}

func (e *Executor) InsertMemo(memo *Memo) (int, error) {
	result, err := e.dbModel.Execute("INSERT INTO Memo (Content,ShortContent,Date) VALUES (?, ?, ?)", (memo.Content), (memo.ShortedContent), (memo.Date))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (e *Executor) UpdateMemo(memo *Memo, id int) error {
	var err error
	if memo.ShortedContent == "" {
		_, err = e.dbModel.Execute("UPDATE Memo SET Content=?, Date=? WHERE Id=?", (memo.Content), (memo.Date), (id))
	} else if memo.Content == "" {
		_, err = e.dbModel.Execute("UPDATE Memo SET ShortContent=?, Date=? WHERE Id=?", (memo.ShortedContent), (memo.Date), (id))
	} else {
		_, err = e.dbModel.Execute("UPDATE Memo SET Content=?, ShortContent=?, Date=? WHERE ID=?", (memo.Content), (memo.ShortedContent), (memo.Date), (id))
	}
	if err != nil {
		return err
	}
	return nil
}

func (e *Executor) ArchiveMemo(memoId []int) (int, error) {
	var total int
	var err error

	for _, v := range memoId {

		err = e.dbModel.InitTransaction()
		if err != nil {
			return total, err
		}
		insertStmt, err := e.dbModel.PrepareStatement("INSERT INTO MemoArchive (Id,Content,ShortContent,Date) SELECT * FROM Memo WHERE Id=?")
		deleteStmt, err := e.dbModel.PrepareStatement("DELETE FROM Memo WHERE Id=?")

		if err != nil {
			return total, err
		}
		_, err = e.dbModel.ExecuteStatment(insertStmt, (v))
		if err != nil {
			return total, err
		}

		_, err = e.dbModel.ExecuteStatment(deleteStmt, (v))
		if err != nil {
			return total, err
		}
		err = e.dbModel.CommitTransaction()
		if err != nil {
			return total, err
		}
		total++
	}
	return total, nil
}

func (e *Executor) RestoreMemo(memoId int) {

	// I don't know if I want implement the restore, for now, no

}

func (e *Executor) DeleteArchivedMemoById(memoId []int) {
	// same
}

func (e *Executor) DeleteMemoArchive() (int, error) {
	result, err := e.dbModel.Execute("DELETE FROM MemoArchive")
	if err != nil {
		return 0, err
	}
	return result, nil
	//DeleteAllMemo(result)
}

func (e *Executor) ShowArchivedMemo(by OrderBy) ([]Memo, error) {
	archivedMemo, err := e.dbModel.Query(fmt.Sprintf("SELECT * FROM MemoArchive %s", orderByMap[by]))
	if err != nil {
		return nil, err
	}
	return archivedMemo, nil
	//ShowMemo(archivedMemo, false)
}

func (e *Executor) GetArchivedMemoById(memoId int) ([]Memo, error) {
	archivedMemo, err := e.dbModel.Query("SELECT * FROM MemoArchive WHERE Id=?", (memoId))
	if err != nil {
		return nil, err
	}
	return archivedMemo, nil
	//ShowMemo(archivedMemo, true)
}
