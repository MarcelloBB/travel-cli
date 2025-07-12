package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "travel",
	Short: "A minimalist curl-like HTTP client in Go",
	Long:  "travel is a CLI tool for making HTTP requests. Like curl, but Go-native.",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	RootCmd.PersistentFlags().StringP("headers", "H", "", "Custom headers in the format 'Key: Value'")
}
