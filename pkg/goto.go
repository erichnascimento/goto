package gotopath

import "log"
import "github.com/erichnascimento/mapdb"
import homeDir "github.com/mitchellh/go-homedir"

// Entry for a path
type PathEntry struct {
  Path string
  Description string
}

func (pathEntry *PathEntry) ExpandPath() string {
  path, _ := homeDir.Expand(pathEntry.Path)
  return path
}

// Create a new PathEntry
func newPathEntry(path string, description string) *PathEntry {
  return &PathEntry{path, description}
}

// Return db file path
func getDBFilePath() (string, error) {
  return homeDir.Expand("~/.goto")
}

var database *mapdb.MapDB

func getDatabase() *mapdb.MapDB {
  if database == nil {
    // Try get the file path
    dbFile, err := getDBFilePath()
    if err != nil {
      log.Fatal(err)
    }

    // open db
    database = mapdb.OpenDB(dbFile, func () {
      mapdb.Register(PathEntry{})
    })
  }

  return database
}

func Add(name, path, description string) *PathEntry {
  pathEntry := newPathEntry(path, description)
  getDatabase().Set(name, pathEntry)
  return pathEntry
}

func List() {

}

func Delete(name string) {
  getDatabase().Del(name)
}

func GoTo(name string) string {
  entry := getDatabase().Get(name).(*PathEntry)
  return entry.ExpandPath()
}

func Close() {
  if database != nil {
    database.Close()
  }
}