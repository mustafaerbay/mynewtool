package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jsonparserCmd represents the jsonparser command
var jsonparserCmd = &cobra.Command{
	Use:   "jsonparser",
	Short: "this will be parse some json data and print the command line",
	Long: `This command gonna get info from binance and
	print related symbol information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonparser called")
	},
}

func init() {
	rootCmd.AddCommand(jsonparserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonparserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonparserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
