package cmd

import (
	"fmt"
	"travel-cli/internal/httpclient"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Send a GET request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run:   runGetCommand,
}

func runGetCommand(cmd *cobra.Command, args []string) {
	url := args[0]
	verbose, _ := cmd.Flags().GetBool("verbose")
	res, err := httpclient.Get(url, verbose)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
}
