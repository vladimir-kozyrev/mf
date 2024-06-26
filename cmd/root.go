package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vladimir-kozyrev/mf/helpers"
	"github.com/vladimir-kozyrev/mf/parse"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
}

var rootCmd = &cobra.Command{
	Use:   "mf",
	Short: "Makefile helper",
	Long:  `mf shows you the contents of Makefile targets without the need to open and scan the file with your own eyes 👀`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists Makefile targets",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("Makefile") // TODO allow specifying the Makefile path via an argument
		if err != nil {
			helpers.PrintToStderrAndExit(err, 1)
		}
		defer f.Close()

		targets, err := parse.GetTargets(f)
		if err != nil {
			helpers.PrintToStderrAndExit(err, 1)
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
			helpers.PrintToStderrAndExit(errMsg, 1)
		}

		f, err := os.Open("Makefile") // TODO allow specifying the Makefile path via an argument
		if err != nil {
			helpers.PrintToStderrAndExit(err, 1)
		}
		defer f.Close()

		targets, err := parse.GetTargetsWithContent(f)
		if err != nil {
			helpers.PrintToStderrAndExit(err, 1)
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
		helpers.PrintToStderrAndExit(err, 1)
	}
}
