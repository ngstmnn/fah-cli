package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list [resource]",
		Short: "List information on different resources",
		Long: `The FAH-Client provides resources like Slots and Queues.
				The list command can be used to list all resources of one
				type or details of a single resource.`,
		Example:   "  fah-cli list slots",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"slot", "slots"},
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "slot", "slots":
				listSlots()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func listSlots() {
	logrus.Error("Listing slots is not yet implemented")
}
