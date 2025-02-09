package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vladimir-kozyrev/mf/internal/parse"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.PersistentFlags().StringVarP(&makefilePath, "file", "f", "Makefile", "path to Makefile")
}

var makefilePath string

var rootCmd = &cobra.Command{
	Use:   "mf",
	Short: "Makefile helper",
	Long:  `mf shows you the contents of Makefile targets without the need to open and scan the file with your own eyes 👀`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists Makefile targets",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(makefilePath)
		if err != nil {
			printToStderrAndExit(err, 1)
		}
		defer f.Close()

		targets, err := parse.GetTargets(f)
		if err != nil {
			printToStderrAndExit(err, 1)
		}

		for _, t := range targets {
			fmt.Println(t)
		}
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows Makefile target declaration and its contents",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errMsg := fmt.Sprintf("show subcommand expects 1 argument, got: %v", args)
			printToStderrAndExit(errMsg, 1)
		}

		f, err := os.Open(makefilePath)
		if err != nil {
			printToStderrAndExit(err, 1)
		}
		defer f.Close()

		targets, err := parse.GetTargetsWithContent(f)
		if err != nil {
			printToStderrAndExit(err, 1)
		}

		for _, t := range targets {
			if t.Name == args[0] {
				fmt.Println(t.Declaration)
				fmt.Println(t.Content)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printToStderrAndExit(err, 1)
	}
}
