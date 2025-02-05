package cmd
import (
    "os"
    "github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
    Use:   "devex",
    Short: "DevEx is a CLI tool to automate various devops tasks.",
    CompletionOptions: cobra.CompletionOptions{
    	HiddenDefaultCmd: true,
    },
}
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
func init() {
    // Additional global flags and configuration settings can be added here
}
