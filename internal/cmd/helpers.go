package cmd

import (
	"fmt"
	"os"
)

func printToStderrAndExit(x any, code int) {
	fmt.Fprintln(os.Stderr, x)
	os.Exit(code)
}
