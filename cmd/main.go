package main

import (
	"fmt"
	"github.com/mir-mirsodikov/crane"
)

func main() {
	rootCmd := &crane.Command{
		Name:  "crane",
		Short: "The crane testing application.",
		Long:  "A sample created testing application for the crane package.",
		Handler: func(cmd *crane.Command, args []string) {
			fmt.Println("Hello, World from the Root Command!")
		},
	}

	versionCmd := &crane.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Short:   "Show version of the app",
		Handler: func(cmd *crane.Command, args []string) {
			fmt.Println("Version 0.1")
		},
	}

	helloCmd := &crane.Command{
		Name:      "hello [name]",
		Aliases:   []string{"h"},
		NumOfArgs: 1,
		Short:     "Say hello to the world",
		Handler: func(cmd *crane.Command, args []string) {
			fmt.Println("Hello,", args[0]+"!")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helloCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
