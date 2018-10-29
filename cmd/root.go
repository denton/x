package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "x",
	Short: "Tool for quering csv and xlsx files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute root comand
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
