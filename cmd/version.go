package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of imgresize",
	Long:  `All software has versions. This is imgresize's`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Img CLI tool v%s %s/%s BuildDate: %s CommitHash:%s\n", Version, runtime.GOOS, runtime.GOARCH, BuildDate, CommitHash)
}
