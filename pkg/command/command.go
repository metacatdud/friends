package command

import "github.com/spf13/cobra"

var (
	CmdRoot  *cobra.Command
	commands []Command
)

func init() {
	CmdRoot = &cobra.Command{
		Use: "friends",
	}
}

type Command interface {
	Cmd() *cobra.Command
	Execute(cmd *cobra.Command, args []string) error
}

type CmdOutput struct {
	Result string
	Err    error
	Cmd    *cobra.Command
}

func Load(c []Command) {
	commands = append(commands, c...)
}

func Run(args []string) CmdOutput {

	buildCommands(commands)

	CmdRoot.SetArgs(args)

	c, err := CmdRoot.ExecuteC()
	return CmdOutput{
		Err: err,
		Cmd: c,
	}
}

func buildCommands(cmds []Command) {
	for _, c := range cmds {
		cmd := c.Cmd()
		if cmd == nil {
			continue
		}
		CmdRoot.AddCommand(cmd)
	}
}
