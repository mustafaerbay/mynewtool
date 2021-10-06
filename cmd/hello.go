package cmd

import (
	"fmt"
	// "strconv"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "says hello to given name",
	Long: `This command just says hello to you`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		
		if name == "" {
			name = "mustafa"
		}
		fmt.Println(name)
	},
}

func init() {
	httpCmd.AddCommand(helloCmd)
}
