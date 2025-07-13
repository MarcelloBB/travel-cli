package cmd

import (
	"fmt"
	"os"
	"travel-cli/internal/httpclient"
	"travel-cli/internal/utils"
	"travel-cli/model"
	"travel-cli/repository"

	"github.com/spf13/cobra"
)

var outputFile string
var saveToCollection string
var requestTitle string

func init() {
	getCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save the response")
	getCmd.Flags().StringVarP(&saveToCollection, "save", "s", "", "Save the request in a collection")
	getCmd.Flags().StringVarP(&requestTitle, "title", "t", "", "Request title for saving in a collection")
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

	if saveToCollection == "" && requestTitle != "" {
		fmt.Println("You must specify a collection to save the request title")
		os.Exit(1)
	}

	if saveToCollection != "" && requestTitle == "" {
		fmt.Println("You must specify a title for the request when saving to a collection")
		os.Exit(1)
	}

	if saveToCollection != "" && requestTitle != "" {
		requestToSave := &model.Request{
			Url:     url,
			Method:  "GET",
			Headers: &headers,
			Title:   requestTitle,
		}
		if err := repository.SaveRequestToCollection(requestToSave, saveToCollection); err != nil {
			fmt.Println("Error saving the request to collection:", err)
			os.Exit(1)
		}
	}

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
