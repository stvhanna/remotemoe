package ssh

import (
	"github.com/fasmide/remotemoe/ssh/command"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CommandReset does various house keeping
func CommandReset(cmd *cobra.Command) {
	cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		flag.Value.Set(flag.DefValue)
	})

	if cmd.HasSubCommands() {
		for _, c := range cmd.Commands() {
			CommandReset(c)
		}
	}
}

// DefaultCmd is the default top level command, embedding all others
func DefaultCmd(session *Session) *cobra.Command {
	c := &cobra.Command{
		Use:          "main",
		SilenceUsage: true,
	}

	c.AddCommand(command.Firsttime())
	c.AddCommand(command.Close(session))
	c.AddCommand(command.Session(session))
	c.AddCommand(command.Host(session))
	c.AddCommand(command.Access(session))
	c.AddCommand(command.HTTP(session))

	return c
}
