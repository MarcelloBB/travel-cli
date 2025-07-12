package cmd

import (
	"fmt"
	"os"
	"travel-cli/internal/httpclient"
	"travel-cli/internal/utils"

	"github.com/spf13/cobra"
)

var outputFile string

func init() {
	getCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save the response")
	RootCmd.AddCommand(getCmd)
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
	headers, _ := cmd.Flags().GetString("headers")

	res, err := httpclient.Get(url, verbose, headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formatted, err := utils.PrettyPrintJSON(res)
	if err != nil {
		fmt.Println("Reponse body (raw):", res)
	} else {
		fmt.Println(formatted.String())
		if outputFile != "" {
			if err := os.WriteFile(outputFile, formatted.Bytes(), 0644); err != nil {
				fmt.Println("Error writing to file:", err)
			} else {
				fmt.Printf("Response saved to %s\n", outputFile)
			}
		}
	}
}
