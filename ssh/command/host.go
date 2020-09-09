package command

import (
	"fmt"

	"github.com/fasmide/remotemoe/router"
	"github.com/fasmide/remotemoe/ssh/command/host"
	"github.com/spf13/cobra"
)

func Host(r router.Routable) *cobra.Command {
	list := func(cmd *cobra.Command, _ []string) error {
		namedRoutes, err := router.Names(r)
		if err != nil {
			return fmt.Errorf("unable to lookup your custom names: %w", err)
		}

		if len(namedRoutes) == 0 {
			cmd.Printf("No active hostnames.\n")
			return nil
		}

		cmd.Printf("Active hostnames:\n")
		for _, nr := range namedRoutes {
			cmd.Printf("%s\n", nr.FQDN())
		}

		return nil
	}

	top := &cobra.Command{
		Use:   "host",
		Short: "Manage hostnames",
		RunE:  list,
	}

	top.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List hostnames",
		RunE:  list,
	})

	top.AddCommand(host.Remove(r))
	top.AddCommand(host.Add(r))

	return top
}
