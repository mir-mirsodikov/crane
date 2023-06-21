<h1>Crane - Go CLI Library</h1>

- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Pre-requisites](#pre-requisites)
  - [Setup repo locally](#setup-repo-locally)
    - [Check out the repo](#check-out-the-repo)
    - [Build the project](#build-the-project)
- [Usage](#usage)
  - [Add the library to your project](#add-the-library-to-your-project)
  - [Include the library in your project](#include-the-library-in-your-project)
  - [Create a root command](#create-a-root-command)
  - [Add a subcommand](#add-a-subcommand)
  - [Specify the number of arguments](#specify-the-number-of-arguments)
  - [All together now](#all-together-now)
- [License](#license)

# Overview

I have built some web servers with Go, but have not used Go for one of its most 
popular use cases: CLI tools. I decided to build a CLI tool to learn more about
Go and CLIs in general. Instead of just building a CLI application, I decided to 
build a library that can be used to build CLI applications. This library is 
called Crane. It is heavily inspired by the already popular [Cobra](https://github.com/spf13/cobra/tree/main) library.

> I don't know that I would recommend using this library over Cobra or using it 
in production. I built this library to learn more about Go and CLIs. It has a 
finite set of features that are useful, but it is not as robust as Cobra.

## Features

- [x] Command
- [x] Args
- [x] Nested Commands
- [ ] Flags
- [ ] Help
- [ ] Bash Completion

# Getting Started

## Pre-requisites

- Go

## Setup repo locally

### Check out the repo

```bash
git clone https://github.com/mir-mirsodikov/crane
```

### Build the project

```bash
make build
```

# Usage

## Add the library to your project

```bash
go get github.com/mir-mirsodikov/crane
```

## Include the library in your project

```go
import "github.com/mir-mirsodikov/crane"
```

## Create a root command

```go

rootCmd := &crane.Command{
		Name:  "crane",
		Short: "The crane testing application.",
		Long:  "A sample created testing application for the crane package.", 
		Handler: func(cmd *crane.Command, args []string) {
			fmt.Println("Hello from the Root Command!")
		},
	} 
	
if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
}
```

## Add a subcommand

```go
versionCmd := &crane.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Short:   "Show version of the app",
		Handler: func(cmd *crane.Command, args []string) {
			fmt.Println("Version 0.1")
		},
	}

rootCmd.AddCommand(versionCmd)
```

## Specify the number of arguments

```go
helloCmd := &crane.Command{
    Name:      "hello [name]",
    Aliases:   []string{"h"},
    NumOfArgs: 1,
    Short:     "Say hello to the world",
    Handler: func(cmd *crane.Command, args []string) {
        fmt.Println("Hello,", args[0]+"!")
    },
}
```

## All together now

```go
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
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details