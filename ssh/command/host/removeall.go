package host

import (
	"github.com/fasmide/remotemoe/router"
	"github.com/spf13/cobra"
)

// RemoveAll removes all the previously added hostnames
func RemoveAll(r router.Routable) *cobra.Command {
	c := &cobra.Command{
		Use:   "all",
		Short: "Remove all hostname(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			removed, err := router.RemoveAll(r)
			if err != nil {
				return err
			}

			// tell the user which hosts where removed
			for _, nr := range removed {
				cmd.Printf("%s removed.\r\n", nr.FQDN())
			}

			return nil
		},
	}

	return c
}
