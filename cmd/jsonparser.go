package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/mustafaerbay/mynewtool/common"
)

// jsonparserCmd represents the jsonparser command
var jsonparserCmd = &cobra.Command{
	Use:   "jsonparser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()
		router.GET("/albums", common.GetAlbums)
		router.GET("/albums/:id", common.GetAlbumByID)
		router.Run(":8080")
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
