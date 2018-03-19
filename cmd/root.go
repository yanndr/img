package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgresize",
	Short: "Imgresize is a bulk image resizer written in go.",
	Run: func(cmd *cobra.Command, args []string) {

	},
	SilenceUsage:  false,
	SilenceErrors: false,
	Args:          cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	rootCmd.AddCommand(versionCmd)
}

//Execute is the root entry of the Cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
