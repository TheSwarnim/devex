package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theswarnim/devex/pkg/ai"

	"github.com/charmbracelet/glamour"
)



// aiCmd represents the ai command
var aiCmd = &cobra.Command{
	Use:   "ai <prompt>",
	Short: "Interact with the AI agent to generate commands or code snippets.",
	Long: `The ai command allows you to interact with an AI agent to generate commands or code snippets
to automate various tasks. For example:

devex ai "generate a new go module"
devex ai "generate a new aws cli command using prod profile to copy file from local to s3 bucket"
devex ai "generate code to read a file in go"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := strings.Join(args, " ")
		response, err := ai.GenerateResponse(prompt)
		if err != nil {
			fmt.Println("Error generating response:", err)
			return
		}

  		out, err := glamour.Render(response, "dark") // Use "light" for light themes
	    if err != nil {
	        panic(err)
	    }

	    // Print the rendered output
	    fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(aiCmd)
}
