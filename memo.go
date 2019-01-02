package main

import (
        "fmt"
        "database/sql"
        "os"
        "time"
        _ "github.com/mattn/go-sqlite3"
        "os/user"
        "log"
)


func CreateMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text, Short text)")
  if err != nil {
    log.Fatal(err)
  }

}

func InsertShort(ArgsString string, ShortString string){
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  _, err = db.Exec("INSERT INTO Things (ToDo, Short) VALUES (?, ?)", (ShortString + "\t\t" + "(" +date+ ")"), (ArgsString + "\t\t" + "(" +date+ ")"))
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()
}


func InsertMemo(ArgsString string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  _, err = db.Exec("INSERT INTO Things (ToDo) VALUES (?)", (ArgsString + "\t\t" + "(" +date+ ")"))
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()
}


func SelectShortMemo(ArgsRowid string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var rows, e = db.Query("SELECT rowid, Short FROM Things WHERE rowid=?", (ArgsRowid))
  if e != nil  {
    log.Fatal(e)
  }
  fmt.Printf("\n Memo:\n")
  for rows.Next() {
    var Short string
    var rowid int
    rows.Scan(&rowid, &Short)
    fmt.Println("\n", rowid, "-" + " " + Short)
  }
  fmt.Printf("\n")
  defer rows.Close()
  defer db.Close()
}



func SelectMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var rows, e = db.Query("SELECT rowid, ToDo FROM Things")
  if e != nil  {
    log.Fatal(e)
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
    log.Fatal(err)
  }
  _, err = db.Exec("DELETE FROM Things WHERE rowid=?", (ArgsInt))
  if err != nil {
    log.Fatal(err)
  defer db.Close()
  }
}


func DeleteAllMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  _, err = db.Exec("DELETE FROM Things")
  if err != nil {
    log.Fatal(err)
  defer db.Close()
  }
}

func GetUserHome() {
  var Home ,_ = user.Current()
  os.Chdir(Home.HomeDir)
  os.Mkdir(".memo", 0700)
  os.Chdir(".memo")
  CreateMemo()
}


func main() {
  GetUserHome()

  if len(os.Args) == 3 && os.Args[1] == "a" && len(os.Args[2]) >= 1{
    InsertMemo(os.Args[2])
  }else if len(os.Args) == 3 && os.Args[1] == "d" && len(os.Args[2]) >= 1{
    DeleteMemo(os.Args[2])
  }else if len(os.Args) == 2 && os.Args[1] == "da"{
    DeleteAllMemo()
  }else if len(os.Args) == 2 && os.Args[1] == "s"{
    SelectMemo()
  }else if len(os.Args) == 5 && os.Args[1] == "a" && os.Args[2] == "sh"{
    InsertShort(os.Args[3], os.Args[4])
  }else if len(os.Args) == 3 && os.Args[1] == "r"{
    SelectShortMemo(os.Args[2])
  }else if len(os.Args) == 1 || os.Args[1] == "h"{
    fmt.Printf("\nYou can use this command:\n" + "\n" +
               "a - To add a memo\n" +
               "d position number - To delete a memo\n" +
               "da  - To delete all memo\n" +
               "s - To show all memo\n" +
               "a sh - Add a shorted memo\n" +
               "r position number - Show the complete memo\n" +
               "h - This message\n" + "\n")
  }else{
    fmt.Println("Something went wrong")
  }
}
