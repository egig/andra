package assets

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "assets",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		//..
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func RootCmd() *cobra.Command {
	return rootCmd
}
