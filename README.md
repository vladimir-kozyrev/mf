## Makefile helper

### Installation

```
go install github.com/vladimir-kozyrev/mf@latest
```

### Example of usage

```
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
