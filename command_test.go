package crane

import (
	"fmt"
	"testing"
)

func TestRootCommand(t *testing.T) {
	rootCmd := &Command{
		Name:      "crane",
		NumOfArgs: 2,
		Handler: func(cmd *Command, args []string) {
			fmt.Println("Hello, World from the Root Command!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		t.Error(err)
	}

	if rootCmd.Name != "crane" {
		t.Error("Expected rootCmd.Name to be crane")
	}
}

func TestSubCommand(t *testing.T) {
	rootCmd := &Command{
		Name:      "crane",
		NumOfArgs: 2,
		Handler: func(cmd *Command, args []string) {
			fmt.Println("Hello, World from the Root Command!")
		},
	}

	subCmd := &Command{
		Name:      "hello",
		NumOfArgs: 1,
		Handler: func(cmd *Command, args []string) {
			fmt.Println("Hello,", args[0]+"!")
		},
	}

	rootCmd.AddCommand(subCmd)

	if err := rootCmd.Execute(); err != nil {
		t.Error(err)
	}

	if rootCmd.Name != "crane" {
		t.Error("Expected rootCmd.Name to be crane")
	}

	if subCmd.Name != "hello" {
		t.Error("Expected subCmd.Name to be hello")
	}

	if len(rootCmd.subCommands) == 0 {
		t.Error("Expected rootCmd.subCommands to have at least one sub command")
	}
}
