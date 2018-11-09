package main

import ("fmt"
        "database/sql"
        "os"
        "time"
        _ "github.com/mattn/go-sqlite3"
        "github.com/mitchellh/go-homedir"
       )


func CreateMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text)")
  if err != nil {
    fmt.Println(err)
  }

}


func InsertMemo(ArgsString string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  _, err = db.Exec("INSERT INTO Things (ToDo) VALUES (?)", (ArgsString + "\t\t" + "(" +date+ ")"))
  if err != nil {
    fmt.Println(err)
  }

  defer db.Close()
}


func SelectMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    fmt.Println(err)
  }
  var rows, error = db.Query("SELECT rowid, ToDo FROM Things")
  if error != nil  {
    fmt.Println(err)
  }
  fmt.Println("\n Memo:\n")
  for rows.Next() {
    var ToDo string
    var rowid int
    rows.Scan(&rowid, &ToDo)
    fmt.Println("\n", rowid, "-" + " " + ToDo)
  }
  fmt.Printf("\n")
  defer rows.Close()
  defer db.Close()
}


func DeleteMemo(ArgsInt string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
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
  var db, err = sql.Open("sqlite3", "./memo.db")
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
   var HomeUser ,_ = homedir.Dir()
   os.Chdir(HomeUser)
   os.Mkdir(".memo", 0700)
   var ExHomeUser ,_ = homedir.Expand("/.memo")
   var FullPath = HomeUser + ExHomeUser
   os.Chdir(FullPath)


  if len(os.Args) == 3 && os.Args[1] == "a" && len(os.Args[2]) >= 1{
    ArgsString := os.Args[2]
    CreateMemo()
    InsertMemo(ArgsString)
  }else if len(os.Args) == 3 && os.Args[1] == "d" && len(os.Args[2]) >= 1{
    ArgsInt := os.Args[2]
    CreateMemo()
    DeleteMemo(ArgsInt)
  }else if len(os.Args) == 2 && os.Args[1] == "da"{
    CreateMemo()
    DeleteAllMemo()
  }else if len(os.Args) == 2 && os.Args[1] == "s"{
    SelectMemo()
  }else if len(os.Args) == 1 || os.Args[1] == "h"{
    fmt.Printf("\nYou can use this command:\n" + "\n" +
               "a - To add a memo\n" +
               "d position number - To delete a memo\n" +
               "da  - To delete all memo\n" +
               "s - To show all memo\n" +
               "h - This message\n" + "\n")
  }else{
    fmt.Println("Something went wrong")
  }
 }
