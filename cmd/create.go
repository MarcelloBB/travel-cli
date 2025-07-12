package cmd

import (
	"fmt"
	"os"
	"travel-cli/repository"

	"github.com/spf13/cobra"
)

var workspaceName string
var collectionName string

func init() {
	createCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name to be created")
	createCmd.Flags().StringVarP(&collectionName, "collection", "c", "", "Collection name to be created")
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a workspace or collection",
	Args:  cobra.NoArgs,
	Run:   runCreateCommand,
}

func runCreateCommand(cmd *cobra.Command, args []string) {
	if workspaceName != "" && collectionName != "" {
		fmt.Println("Cannot create both workspace and collection at the same time, use --workspace (-w) or --collection (-c)")
		os.Exit(1)
	}

	if workspaceName == "" && collectionName == "" {
		fmt.Println("You must specify either a workspace or a collection to create")
		os.Exit(1)
	}

	if workspaceName != "" {
		workspace, err := repository.CreateWorkspace(workspaceName)
		if err != nil {
			fmt.Println("Error creating workspace:", err)
			os.Exit(1)
		}
		fmt.Printf("Workspace '%s' created successfully with id %d\n", workspace.Name, workspace.IdWorkspace)
		return
	}

	if collectionName != "" {
		collection, err := repository.CreateCollection(collectionName, 1) // USE CURRENT WORKSPACE ID
		// create table named execution and a column named id_workspace that stores the selected workspace id
		if err != nil {
			fmt.Println("Error creating collection:", err)
			os.Exit(1)
		}
		fmt.Printf("Collection '%s' created successfully with id %d\n", collection.Title, collection.IdCollection)
		return
	}

	// verbose, _ := cmd.Flags().GetBool("verbose")
	// headers, _ := cmd.Flags().GetString("headers")

	// res, err := httpclient.Get(url, verbose, headers)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(res)
}
