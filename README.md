## Makefile helper

### Installation

```
go install github.com/vladimir-kozyrev/mf@latest
```

### Example of usage

```
❯ mf
mf shows you the contents of Makefile targets without the need to open and scan the file with your own eyes 👀

Usage:
  mf [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        lists Makefile targets
  show        shows Makefile target declaration and its contents

Flags:
  -f, --file string   path to Makefile (default "Makefile")
  -h, --help          help for mf

Use "mf [command] --help" for more information about a command.

❯ mf list -f tests/Makefile
build
test
lint

❯ mf show build -f tests/Makefile
build:
	echo "build"
```

### Known limitatations
At the moment, this tools supports only lowercased target names that contain a-z letters, underscores, and dashes.
