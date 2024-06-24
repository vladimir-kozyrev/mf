package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vladimir-kozyrev/mfh/helpers"
	"github.com/vladimir-kozyrev/mfh/parse"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var rootCmd = &cobra.Command{
	Use:   "mfh",
	Short: "Makefile helper",
	Long:  `mfh shows you the contents of Makefile targets without the need to open and scan the file with your own eyes 👀`,
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
			cleanedT := parse.RemoveAllAfterFirstColon(t)
			fmt.Println(cleanedT)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		helpers.PrintToStderrAndExit(err, 1)
	}
}
