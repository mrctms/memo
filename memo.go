package main

import ("fmt"
        "database/sql"
        "os"
        _ "github.com/mattn/go-sqlite3"
        "github.com/mitchellh/go-homedir"
       )


func CreateMemo() {
  db, err:= sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text)")
  if err != nil {
    fmt.Println(err)
  }

}


func InsertMemo(ArgsString string) {
  db, err:= sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  _, err = db.Exec("INSERT INTO Things (ToDo) VALUES (?)", (ArgsString))
  if err != nil {
    fmt.Println(err)
  }

  defer db.Close()
}


func SelectMemo() {
  db, err:= sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  rows, err:= db.Query("SELECT * FROM Things")
  if err != nil  {
    fmt.Println(err)
  }
  fmt.Println("\nMemo:\n")
  for rows.Next() {
    var ToDo string
    rows.Scan(&ToDo)
    fmt.Println(ToDo)
  }
  fmt.Printf("\n")
  defer rows.Close()
  defer db.Close()
}


func DeleteMemo(ArgsInt string) {
  db, err:= sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  _, err = db.Exec("DELETE FROM Things WHERE rowid=?", (ArgsInt))
  if err != nil {
    fmt.Println(err)
  defer db.Close()
  }
}


func DeleteAllMemo() {
  db, err:= sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  _, err = db.Exec("DELETE FROM Things")
  if err != nil {
    fmt.Println(err)
  defer db.Close()
  }
}



func main() {
   HomeUser ,_:= homedir.Dir()
   os.Chdir(HomeUser)
   os.Mkdir(".memo", 0700)
   ExHomeUser ,_:= homedir.Expand("/.memo")
   FullPath := HomeUser + ExHomeUser
   os.Chdir(FullPath)


  if ArgsChoice := os.Args[1]; ArgsChoice == "a"{
    ArgsString := os.Args[2]
    CreateMemo()
    InsertMemo(ArgsString)
  }else if ArgsChoice := os.Args[1]; ArgsChoice == "d"{
    ArgsInt := os.Args[2]
    CreateMemo()
    DeleteMemo(ArgsInt)
  }else if ArgsChoice := os.Args[1]; ArgsChoice == "da"{
    CreateMemo()
    DeleteAllMemo()
  }else if ArgsChoice := os.Args[1]; ArgsChoice == "s"{
    SelectMemo()
  }else if ArgsChoice := os.Args[1]; ArgsChoice == "h"{
    fmt.Printf("\nYou can use this command:\n" + "\n" +
               "a - To add a memo\n" +
               "d position number - To delete a memo\n" +
               "da  - To delete all memo\n" +
               "s - To show all memo\n" + "\n")
  }
 }
