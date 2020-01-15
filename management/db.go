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
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type MemoDB struct {
	DB *sql.DB
}

func (m *MemoDB) CreateMemoTable() error {
	_, err := m.DB.Exec("CREATE TABLE IF NOT EXISTS Things (ID integer primary key autoincrement, ToDo text, Short text, DateTime text)")
	if err != nil {
		return err
	}
	return nil
}

func (m *MemoDB) CreateMemo(memo *Memo) error {
	_, err := m.DB.Exec("INSERT INTO Things (ToDo, DateTime, Short) VALUES (?, ?, ?)", (memo.Content), (memo.Date), (memo.ShortedContent))
	if err != nil {
		return err
	}
	return nil

}

func (m *MemoDB) DeleteMemos(id []string) error {
	for _, v := range id {
		_, err := m.DB.Exec("DELETE FROM Things WHERE ID=?", (v))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MemoDB) DeleteAllMemos() error {
	_, err := m.DB.Exec("DELETE FROM Things")
	if err != nil {
		return err
	}
	return nil
}

func (m *MemoDB) SelectMemo() error {
	var rows, err = m.DB.Query("SELECT ID, ToDo, DateTime FROM Things")
	if err != nil {
		return err
	}
	fmt.Printf("\n Memo:\n")
	for rows.Next() {
		var toDo string
		var id int
		var dateTime string
		rows.Scan(&id, &toDo, &dateTime)
		fmt.Println("\n", id, "-", dateTime, "-", toDo)
	}
	fmt.Printf("\n")
	rows.Close()
	return nil
}

func (m *MemoDB) SelectShortMemo(id string) error {
	var rows, err = m.DB.Query("SELECT ID, Short, DateTime FROM Things WHERE ID=?", (id))
	if err != nil {
		return err
	}
	for rows.Next() {
		var short string
		var id int
		var dateTime string
		rows.Scan(&id, &short, &dateTime)
		if short == "" {
			return errors.New(fmt.Sprintf("\nERROR: Memo with ID %d does not have a shorted memo\n\n", id))
		} else {
			fmt.Printf("\n Memo:\n")
			fmt.Println("\n", id, "-", dateTime, "-", short)
			fmt.Printf("\n")
		}
	}
	rows.Close()
	return nil
}

func (m *MemoDB) EditMemo(id string, memo *Memo) error {
	var err error
	if memo.ShortedContent == "" {
		_, err = m.DB.Exec("UPDATE Things SET ToDo=?, DateTime=? WHERE ID=?", (memo.Content), (memo.Date), (id))
	} else if memo.Content == "" {
		_, err = m.DB.Exec("UPDATE Things SET Short=?, DateTime=? WHERE ID=?", (memo.ShortedContent), (memo.Date), (id))
	} else {
		_, err = m.DB.Exec("UPDATE Things SET ToDo=?, Short=?, DateTime=? WHERE ID=?", (memo.Content), (memo.ShortedContent), (memo.Date), (id))
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *MemoDB) ImportFromFile(file string) error {
	var fileContent []string
	fileToImport, err := os.Open(file)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(fileToImport)
	for fileScanner.Scan() {
		fileContent = append(fileContent, fileScanner.Text())
	}
	for _, v := range fileContent {
		newMemo := NewMemo(v, "")
		m.CreateMemo(newMemo)
	}
	return nil
}
