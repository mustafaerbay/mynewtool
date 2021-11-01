package cmd

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http server",
	Long: `This command start basic http server to you.
	You can send some get request to related port 
	DEFAULT_PORT is 8090. `,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")

		if port == "" {
			port = "8090"
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Server start listening on port: " + port)
		})
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Printf("port:%s not opened", port)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	httpCmd.PersistentFlags().String("name", "n", "server name")
	httpCmd.PersistentFlags().String("file", "f", "server name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	httpCmd.Flags().StringP("port", "p", "", "Set your listening port")
}
