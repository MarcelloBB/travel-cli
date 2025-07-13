package cmd

import (
	"fmt"
	"os"
	"travel-cli/repository"

	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().BoolP("workspace", "w", false, "List all workspaces")
	listCmd.Flags().BoolP("collection", "c", false, "List all collections")
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces or collections",
	Args:  cobra.NoArgs,
	Run:   runListCommand,
}

func runListCommand(cmd *cobra.Command, args []string) {
	listWorkspaces, _ := cmd.Flags().GetBool("workspace")
	listCollections, _ := cmd.Flags().GetBool("collection")

	if listWorkspaces && listCollections {
		fmt.Println("Cannot list both workspace and collection at the same time, use --workspace (-w) or --collection (-c)")
		os.Exit(1)
	}

	if !listWorkspaces && !listCollections {
		fmt.Println("You must specify either --workspace (-w) or --collection (-c) to list")
		os.Exit(1)
	}

	if listWorkspaces {
		workspaces, err := repository.ListWorkspaces()
		if err != nil {
			fmt.Println("Error listing all workspaces:", err)
			os.Exit(1)
		}
		for _, workspace := range workspaces {
			if workspace.Current == 1 {
				fmt.Printf("-> * %s (current)\n", workspace.Name)
			} else {
				fmt.Printf("   * %s\n", workspace.Name)
			}
		}
		return
	}

	if listCollections {
		collections, err := repository.ListCollections()
		if err != nil {
			fmt.Println("Error listing all collections:", err)
			os.Exit(1)
		}
		for _, collection := range collections {
			fmt.Printf("* %s\n", collection.Title)
			for _, request := range collection.Requests {
				fmt.Printf("  - %s  %s\n", request.Method, request.Title)
			}
		}
		return
	}
}
