package hange

type CommandManager struct {
	commands map[string]Command
}

type CommandExecutor interface {
	Execute(arguments ...Argument)
}

type Command struct {
	commandFunction func(arguments ...Argument)
	name string
	aliases []string
}

func (c *Command) Name() string {
	return c.name
}

func (c Command) Aliases() []string {
	return c.aliases
}

func (c *Command) Execute(arguments ...Argument) {
	c.commandFunction(arguments...)
}

func (cm *CommandManager) RegisterCommand(name string, executed func(arguments ...Argument), aliases ...string) *Command {
	result := Command{
		name: name,
		commandFunction: executed,
		aliases: aliases,
	}

	cm.commands[name] = result

	return &result
}
