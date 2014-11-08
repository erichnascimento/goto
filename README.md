# goto

A simple utility which may be used for manage your personal or project folders

## Instalation

```
$ go get github.com/erichnascimento/goto
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
```

### List paths

```
$ goto list
  myproject    "My Project Folder"
```

### Go to path

```
$ goto myproject
$ pwd
~/dev/myproject
```

### Delete a path

```
$ goto myproject
  Sorry, path entry "myproject" not found...

```

# License

MIT