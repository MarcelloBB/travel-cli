package cmd

import (
	"fmt"
	"os"
	"travel-cli/repository"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(useWorkspaceCmd)
}

var useWorkspaceCmd = &cobra.Command{
	Use:   "use [workspace]",
	Short: "Use a specific workspace. This sets the current workspace for subsequent commands. Set it by passing the workspace name as an argument.",
	Args:  cobra.ExactArgs(1),
	Run:   runUseWorkspaceCommand,
}

func runUseWorkspaceCommand(cmd *cobra.Command, args []string) {
	workspaceName := args[0]
	if err := repository.SetCurrentWorkspace(workspaceName); err != nil {
		fmt.Println("Error listing workspaces:", err)
		os.Exit(1)
	}
	fmt.Printf("Workspace '%s' is now set as the current workspace.\n", workspaceName)
}
