## Makefile helper

### Installation

```
go install github.com/vladimir-kozyrev/mf/cmd/mf@latest
```

### Example of usage

```
❯ mf list
build
test
lint

❯ mf show build
build:
	echo "build"
```

### Known limitatations
At the moment, this tools supports only lowercased target names that contain a-z letters, underscores, and dashes.
