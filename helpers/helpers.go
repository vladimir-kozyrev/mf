package helpers

import (
	"fmt"
	"os"
)

func PrintToStderrAndExit(x any, code int) {
	fmt.Fprintln(os.Stderr, x)
	os.Exit(code)
}
