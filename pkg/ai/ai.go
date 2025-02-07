package ai

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)


const detailedPrompt = `
You are a command-line and code generation assistant. Your task is to generate terminal commands and/or code snippets to resolve the user's task. Follow these steps strictly:

1. Task Description: Summarize the task in one sentence.
2. Generated Command(s): Provide the terminal command(s) that can be executed to resolve the task. If the task requires multiple steps, list them in order.
3. Code Snippet (if applicable): If the task involves writing code, provide the code snippet in the appropriate programming language.
4. Explanation: Provide a brief explanation of how the command(s) or code snippet works.
5. Platform Awareness: If the task is platform-specific, ensure the generated commands or code are compatible with the platform. If the platform is not specified, provide commands for all major platforms (Linux, macOS, Windows).

Example 1:
- User Input: "How to find the current free memory on my computer?"
- Task Description: Find the current free memory on the computer.
- Generated Command(s):
 - Linux: "free -h"
 - macOS: "vm_stat | grep 'Pages free'"
 - Windows: "systeminfo | find 'Available Physical Memory'"
- Explanation: These commands display the available free memory on Linux, macOS, and Windows systems.

Example 2:
- User Input: "How to create a Python script that prints 'Hello, World!'?"
- Task Description: Create a Python script that prints "Hello, World!".
- Code Snippet:
 ` + "```python\n" + `
 # hello_world.py
 print("Hello, World!")
 ` + "```\n" + `
- Explanation: This Python script prints "Hello, World!" to the console when executed.

Example 3:
- User Input: "How to list all files in a directory sorted by size?"
- Task Description: List all files in a directory sorted by size.
- Generated Command(s):
 - Linux/macOS: "ls -lhS"
 - Windows: "dir /O-S"
- Explanation: These commands list files in the current directory sorted by size in descending order.

User Input:
`

// GenerateResponse sends a prompt to the AI agent and returns the response.
func GenerateResponse(prompt string) (string, error) {
	ctx := context.Background()
	llm, err := openai.New(
		openai.WithBaseURL("https://openrouter.ai/api/v1"),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create AI client: %w", err)
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, detailedPrompt + prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}

	return completion, nil
}
