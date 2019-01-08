package main

import (
        "fmt"
        "database/sql"
        "os"
        "time"
        _ "github.com/mattn/go-sqlite3"
        "os/user"
        "log"
        "flag"
)

var (
  Add = flag.Bool("a", false, "[memo] | To add a memo")
  AddShort = flag.Bool("ash", false, "[long memo] [shorted memo] | Add a shorted memo")
  Show = flag.Bool("s", false, "To show all memo")
  Delete = flag.Bool("d", false, "[position number] | To delete a memo")
  DeleteAll = flag.Bool("da", false, "To delete all memo")
  Reveal = flag.Bool("r", false, "[position number] | Show the complete memo")
  Modify = flag.Bool("m", false, "[position number] | To edit a memo")
  ModifyShort = flag.Bool("msh", false, "[position number] | To edit the memo behind the shorted memo")
)

func CreateMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  db.Exec("CREATE TABLE IF NOT EXISTS Things (ToDo text, Short text)")
}

func InsertShort(ArgsString string, ShortString string){
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  db.Exec("INSERT INTO Things (ToDo, Short) VALUES (?, ?)", (ShortString + "\t\t" + "(" +date+ ")"), (ArgsString + "\t\t" + "(" +date+ ")"))
  defer db.Close()
}


func InsertMemo(ArgsString string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  db.Exec("INSERT INTO Things (ToDo) VALUES (?)", (ArgsString + "\t\t" + "(" +date+ ")"))
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

func ModifyMemo(ArgsStringR string, ArgsStringM string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  db.Exec("UPDATE Things SET ToDo=? WHERE rowid=?", (ArgsStringM + "\t\t" + "("+date+")"), (ArgsStringR))
  defer db.Close()
}

func ModifyMemoShort(ArgsStringR string, ArgsStringS string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  var t = time.Now()
  var date = t.Format("2006-01-02 15:04:05")
  db.Exec("UPDATE Things SET Short=? WHERE rowid=?", (ArgsStringS + "\t\t" + "("+date+")"), (ArgsStringR))
  defer db.Close()
}

func DeleteMemo(ArgsInt string) {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  db.Exec("DELETE FROM Things WHERE rowid=?", (ArgsInt))
  defer db.Close()
}


func DeleteAllMemo() {
  var db, err = sql.Open("sqlite3", "./memo.db")
  if err != nil {
    log.Fatal(err)
  }
  db.Exec("DELETE FROM Things")
  defer db.Close()
}

func GetUserHome() {
  var Home, _ = user.Current()
  os.Chdir(Home.HomeDir)
  os.Mkdir(".memo", 0700)
  os.Chdir(".memo")
  CreateMemo()
}


func main() {
  GetUserHome()
  flag.Parse()
  if *Add {
    InsertMemo(os.Args[2])
  }else if *Delete {
    DeleteMemo(os.Args[2])
  }else if *DeleteAll {
    DeleteAllMemo()
  }else if *Show {
    SelectMemo()
  }else if *AddShort {
    InsertShort(os.Args[2], os.Args[3])
  }else if *Reveal {
    SelectShortMemo(os.Args[2])
  }else if *Modify {
    ModifyMemo(os.Args[2], os.Args[3])
  }else if *ModifyShort {
    ModifyMemoShort(os.Args[2], os.Args[3])
  }else {
    fmt.Println("Something went wrong")
  }
}
