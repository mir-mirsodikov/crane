package crane

import (
	"log"
	"os"
)

type CommandHandler func(*Command, []string)

type Command struct {
	Name        string
	NumOfArgs   int
	Args        []string
	Aliases     []string
	Short       string
	Long        string
	subCommands []*Command
	Handler     CommandHandler
	parent      *Command
}

type Commands []Command

// Execute executes the command.
// Usually called from the root command.
func (c *Command) Execute() error {
	if c == nil {
		return nil
	}
	if c.parent == nil {
		c.Args = os.Args[1:]
	}

	log.Default().Println("c.Args:", c.Args)

	subCmd, args, err := c.Traverse(c.Args[1:])

	if err != nil {
		panic(err)
	}

	if subCmd == c {
		c.execute(c.Args)
		return nil
	}

	subCmd.execute(args)

	return nil
}

// execute executes the command.
// Called from the Execute method.
// Validates the number of arguments passed to the command.
func (c *Command) execute(args []string) {
	invalidArgCount := len(args) != c.NumOfArgs

	log.Default().Println("len(args):", len(args), "c.NumOfArgs:", c.NumOfArgs)
	log.Default().Println("args:", args)

	if invalidArgCount {
		log.Default().Println("Not enough arguments passed to command:", c.Name)
		return
	}

	c.Handler(c, args)
}

func contains[T comparable](arr []T, item T) bool {
	for _, x := range arr {
		if item == x {
			return true
		}
	}
	return false
}

// AddCommand adds a command to the list of commands.
// Incrementally add commands to the application.
func (c *Command) AddCommand(cmds ...*Command) {
	for i, x := range cmds {
		if cmds[i] == c {
			panic("Command cannot be child of itself.")
		}

		x.parent = c
		c.subCommands = append(c.subCommands, x)
	}
}

// findChildCommand finds a child command by name or alias.
func (c *Command) findChildCommand(cmdName string) *Command {
	matches := make([]*Command, 0)

	for _, cmd := range c.subCommands {
		if cmd.Name == cmdName || contains(cmd.Aliases, cmdName) {
			matches = append(matches, cmd)
		}
	}

	if len(matches) == 1 {
		return matches[0]
	} else if len(matches) > 1 {
		panic("Ambiguous command name or alias: " + cmdName)
	}

	return nil
}

// Traverse traverses the command tree.
func (c *Command) Traverse(args []string) (*Command, []string, error) {
	if len(args) == 0 {
		return c, c.Args, nil
	}

	cmd := c.findChildCommand(args[0])

	if cmd == nil {
		return c, args, nil
	}

	return cmd.Traverse(args[1:])
}
