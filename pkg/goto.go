package gotopath

import "fmt"
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

  db := getDatabase()
  db.Set(name, pathEntry)
  db.Save()
  return pathEntry
}

func getEntry(name string) *PathEntry {
  value := getDatabase().Get(name)
  if value != nil {
    entry := value.(PathEntry)
    return &entry
  }

  return nil
}

func List() (str string) {
  db := getDatabase()
  for _, name := range db.Keys() {
    if entry := getEntry(name); entry != nil {
      str += fmt.Sprintf("  \033[1;32m%-12s\033[0m%s\n", name, entry.Description)
    }
  }

  return
}

func Delete(name string) {
  db := getDatabase()
  db.Del(name)
  db.Save()
}

func GetPath(name string) string {
  entry := getEntry(name)
  if entry != nil {
    return entry.ExpandPath()
  }

  return ""
}

func Close() {
  if database != nil {
    database.Close()
  }
}