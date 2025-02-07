package cmd
import (
	"log"
    "os"
    "github.com/spf13/cobra"
    "github.com/joho/godotenv"
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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
    // Additional global flags and configuration settings can be added here
}
