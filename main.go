package main

import "fmt"
import "github.com/tj/docopt"
import "github.com/erichnascimento/goto/pkg"

const VERSION string = "0.0.1"

const Usage = `
  Usage: 
    goto list
    goto add <name> <path> [<description>]
    goto rm <name>
    goto <name>

  The most commonly used goto commands are:
    add        Add an entry to the index
    list       List all entries indexed
    rm         Remove a entry from index

`

func add(name, path, description interface{}) {
  var desc string
  if description != nil {
    desc = description.(string)
  }
  gotopath.Add(name.(string), path.(string), desc)
}

func list() {
  fmt.Println(gotopath.List())
}

func getPath(name string) {
  path := gotopath.GetPath(name)
  if path != "" {
    fmt.Println(path)
  }
}

func del(name string) {
  gotopath.Delete(name)
}

func main() {
  args, _ := docopt.Parse(Usage, nil, true, VERSION, false)
  //fmt.Println(args)

  defer gotopath.Close()

  switch {
    case args["add"].(bool):
      add(args["<name>"], args["<path>"], args["<description>"])

    case args["list"].(bool):
      list()

    case args["rm"].(bool):
      del(args["<name>"].(string))

    default:
      getPath(args["<name>"].(string))
  }


  /*defer gotopath.Close()

  gotopath.Add("name", "path", "Description")
  gotopath.List()
  fmt.Print(gotopath.GoTo("name"))
  gotopath.Delete("name")
  */

}