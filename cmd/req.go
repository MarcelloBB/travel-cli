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

var collectionToSearch string
var requestToSend string

func init() {
	requestCmd.Flags().StringVarP(&collectionToSearch, "collection", "c", "", "Collection to search for requests")
	requestCmd.Flags().StringVarP(&requestToSend, "request", "r", "", "Saved request to send")
	requestCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save the response")
	RootCmd.AddCommand(requestCmd)
}

var requestCmd = &cobra.Command{
	Use:   "req",
	Short: "Send a saved request",
	Args:  cobra.NoArgs,
	Run:   runRequestCommand,
}

func runRequestCommand(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetBool("verbose")

	if collectionToSearch == "" && requestToSend == "" {
		fmt.Println("You must specify a collection to search for requests and a saved request title to send")
		os.Exit(1)
	}

	if collectionToSearch == "" {
		fmt.Println("You must specify a collection to search for requests")
		os.Exit(1)
	}

	if requestToSend == "" {
		fmt.Println("You must specify a saved request title to send")
		os.Exit(1)
	}

	collection, err := repository.ListCollections(collectionToSearch)
	if err != nil {
		fmt.Println("Error retrieving collection:", err)
		os.Exit(1)
	}

	if len(collection) == 0 {
		fmt.Printf("No collection found with the name '%s'\n", collectionToSearch)
		os.Exit(1)
	}

	savedRequest := &model.Request{}

	for _, request := range collection[0].Requests {
		if request.Title == requestToSend {
			savedRequest = &request
			break
		}
	}

	res, err := httpclient.Get(savedRequest.Url, verbose, *savedRequest.Headers)
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
