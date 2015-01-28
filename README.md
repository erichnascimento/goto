# goto

A simple utility which may be used to manage and navigate on your favorite folders

## Instalation

```
$ go get github.com/erichnascimento/goto
```

Append this line in your `~/.bashrc` or `~/.bash_aliases` file
```
. $GOPATH/src/github.com/erichnascimento/goto/script/goto.sh
```

## Usage

```
Usage: 
  goto list
  goto add <name> <path> [<description>]
  goto rm <name>
  goto <name>

The most commonly used goto commands are:
  add        Add an entry to the index
  list       List all entries indexed
  rm         Remove a entry from index
```

## Examples

### Add myproject path

```
$ goto add myproject ~/dev/myproject "My Project Folder"
$ goto add music ~/Music "Personal music dir"
```

### List paths
List all entries in alphabetic order
```
$ goto list
  music       Personal music dir
  myproject   My Project Folder   
```

### Go to path
Navigate to path entry
```
$ goto myproject
$ pwd
/home/erichnascimento/dev/myproject
$ goto music
$ pwd
/home/erichnascimento/Music
```

### Delete a path entry
Remove a path entry from index
```
$ goto rm myproject
$ goto myproject
  Sorry, path entry "myproject" not found...
```

# License

MIT
