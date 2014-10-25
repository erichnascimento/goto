package main

import "fmt"
import "github.com/erichnascimento/goto/pkg"

const VERSION string = "0.0.1"

func main() {

  defer gotopath.Close()

  gotopath.Add("name", "path", "Description")
  gotopath.List()
  fmt.Print(gotopath.GoTo("name"))
  gotopath.Delete("name")

  /*

  // Try get the file path
  dbFile, err := getDBFilePath()
  if err != nil {
    log.Fatal(err)
  }

  // open db
  db := mapdb.OpenDB(dbFile, func () {
    mapdb.Register(PathEntry{})
  })

  defer db.Close()

  //db.Set("selecty", NewPathEntry("~/dev/selecty", "Projeto Selecty"))
  //db.Set("goto", NewPathEntry("~/dev/goto", "Projeto Goto"))
  //db.Set("selecty-mobile", NewPathEntry("~/dev/selecty-mobile", "Projeto Selecty Mobile"))
  //entry := db.Get("selecty")
  //fmt.Println(entry)
  //
  var entry *PathEntry

  db.Set("goto", NewPathEntry("~/dev/playground", "goto Project wrote in golang"))
  entry = db.Get("goto").(*PathEntry)
  fmt.Print(entry.ExpandPath())

  //fmt.Println(db.Keys())

  //db.Del("selecty")
  //entry = db.Get("selecty")
  //fmt.Println(entry)

  //db.Set("selecty", "~/dev/selecty")
  //entry = db.Get("selecty")
  //fmt.Println(entry)

  //fmt.Println(db.GetFileName())
  err = db.Save()
  if err != nil {
    fmt.Println(err)
  }
  //fmt.Println(dbFile)*/
}