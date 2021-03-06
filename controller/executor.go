package controller

import (
	"fmt"
	"memo/model"
	"memo/view"
)

type OrderBy int

const (
	Asc  OrderBy = 1
	Desc         = 2
)

var orderByMap = map[OrderBy]string{Asc: "ORDER BY Id ASC", Desc: "ORDER BY Id DESC", 0: "ORDER BY Id ASC"}

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
func (e *Executor) CreateMemo(content string, shortContent string, date string) *model.Memo {
	memo := &model.Memo{
		Content:        content,
		ShortedContent: shortContent,
		Date:           date,
	}
	return memo
}

func (e *Executor) GetMemo(by OrderBy) {
	memos, err := e.dbModel.Query(fmt.Sprintf("SELECT * FROM Memo %s", orderByMap[by]))
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ShowMemo(memos, false)
}

func (e *Executor) GetMemoById(id int, reveal bool) {
	memo, err := e.dbModel.Query("SELECT * FROM Memo WHERE Id=?", (id))
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ShowMemo(memo, reveal)
}
func (e *Executor) DeleteAllMemo() {
	result, err := e.dbModel.Execute("DELETE FROM Memo")
	if err != nil {
		fmt.Println(err)
		return
	}
	view.DeleteAllMemo(result)
}

func (e *Executor) DeleteMemoById(id []int) {
	var total int
	for _, v := range id {
		result, err := e.dbModel.Execute("DELETE FROM Memo WHERE Id=?", (v))
		if err != nil {
			fmt.Println(err)
			return
		}
		total += result
	}
	view.DeleteMemoById(total)
}

func (e *Executor) InsertMemo(memo *model.Memo) {
	result, err := e.dbModel.Execute("INSERT INTO Memo (Content,ShortContent,Date) VALUES (?, ?, ?)", (memo.Content), (memo.ShortedContent), (memo.Date))
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	view.UpdateMemo(id)
}

func (e *Executor) ArchiveMemo(memoId []int) {
	var total int
	var err error

	for _, v := range memoId {

		err = e.dbModel.InitTransaction()
		if err != nil {
			fmt.Println(err)
			return
		}
		insertStmt, err := e.dbModel.PrepareStatement("INSERT INTO MemoArchive (Id,Content,ShortContent,Date) SELECT * FROM Memo WHERE Id=?")
		deleteStmt, err := e.dbModel.PrepareStatement("DELETE FROM Memo WHERE Id=?")

		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = e.dbModel.ExecuteStatment(insertStmt, (v))
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = e.dbModel.ExecuteStatment(deleteStmt, (v))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = e.dbModel.CommitTransaction()
		if err != nil {
			fmt.Println(err)
			return
		}
		total++
	}
	view.ArchiveMemo(total)
}

func (e *Executor) RestoreMemo(memoId int) {

	// I don't know if I want implement the restore, for now, no

}

func (e *Executor) DeleteArchivedMemoById(memoId []int) {
	// same
}

func (e *Executor) DeleteMemoArchive() {
	result, err := e.dbModel.Execute("DELETE FROM MemoArchive")
	if err != nil {
		fmt.Println(err)
		return
	}
	view.DeleteAllMemo(result)
}

func (e *Executor) ShowArchivedMemo(by OrderBy) {
	archivedMemo, err := e.dbModel.Query(fmt.Sprintf("SELECT * FROM MemoArchive %s", orderByMap[by]))
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ShowMemo(archivedMemo, false)
}

func (e *Executor) GetArchivedMemoById(memoId int) {
	archivedMemo, err := e.dbModel.Query("SELECT * FROM MemoArchive WHERE Id=?", (memoId))
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ShowMemo(archivedMemo, true)
}
