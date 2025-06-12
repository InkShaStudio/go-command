package command

import "fmt"

type CommandHandle = func(cmd *SCommand)

type SCommand struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Summary     string              `json:"summary"`
	Args        []ICommandArgValue  `json:"args"`
	Flags       []ICommandFlagValue `json:"flags"`
	SubCommand  []*SCommand         `json:"subCommand"`
	Handle      CommandHandle       `json:"-"`
}

type ICommand interface {
	ChangeName(name string) *SCommand
	ChangeDescription(description string) *SCommand
	ChangeSummary(summary string) *SCommand
	AddArgs(args ...*SCommandArg[any]) *SCommand
	AddFlags(flags ...*SCommandFlag[any]) *SCommand
	AddSubCommand(subCommand ...*SCommand) *SCommand
	RegisterHandler(handle CommandHandle) *SCommand
	GetArgsCount(onlyEmpty bool) int
}

func (cmd *SCommand) ChangeName(name string) *SCommand {
	cmd.Name = name
	return cmd
}

func (cmd *SCommand) ChangeDescription(description string) *SCommand {
	cmd.Description = description
	return cmd
}

func (cmd *SCommand) ChangeSummary(summary string) *SCommand {
	cmd.Summary = summary
	return cmd
}

func (cmd *SCommand) AddArgs(args ...ICommandArgValue) *SCommand {
	cmd.Args = append(cmd.Args, args...)
	flag := false

	for i, arg := range cmd.Args {
		arg.SetIndex(i)

		if flag {
			panic(fmt.Sprintf("command '%s': empty argument '%s' (index %d) appears after non-empty argument; all optional args must follow required ones", cmd.Name, arg.GetName(), i))
		}
		if arg.GetValue() == nil {
			flag = true
		}
	}

	return cmd
}

func (cmd *SCommand) AddFlags(flags ...ICommandFlagValue) *SCommand {
	cmd.Flags = append(cmd.Flags, flags...)

	return cmd
}

func (cmd *SCommand) AddSubCommand(subCommand ...*SCommand) *SCommand {
	cmd.SubCommand = append(cmd.SubCommand, subCommand...)
	return cmd
}

func (cmd *SCommand) GetArgsCount(onlyEmpty bool) int {
	count := 0

	for _, arg := range cmd.Args {
		if !onlyEmpty && arg.GetValue() != nil {
			continue
		}

		count++
	}

	return count
}

func (cmd *SCommand) RegisterHandler(handle CommandHandle) *SCommand {
	cmd.Handle = handle
	return cmd
}

func NewCommand(name string) *SCommand {
	return &SCommand{
		Name:        name,
		Description: "",
		Summary:     "",
		Args:        []ICommandArgValue{},
		Flags:       []ICommandFlagValue{},
		SubCommand:  []*SCommand{},
		Handle:      nil,
	}
}
