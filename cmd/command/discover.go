package command

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
)

type Discoverer interface {
	Discover() ([]huego.Bridge, error)
}

type DefaultDiscoverer struct{}

func (dd *DefaultDiscoverer) Discover() ([]huego.Bridge, error) {
	bridges, err := huego.DiscoverAll()
	if err != nil {
		return []huego.Bridge{}, err
	}

	return bridges, nil
}

var _ Discoverer = (*DefaultDiscoverer)(nil) // Assure that when new method is added to interface, compiler will scream

func init() {
	dd := &DefaultDiscoverer{}
	rootCmd.AddCommand(NewDiscoverCmd(dd))
}

func NewDiscoverCmd(dd Discoverer) *cobra.Command {
	return &cobra.Command{
		Use:   "discover",
		Short: "Discover Philips Hue Bridge(s)",
		Long:  "Discover all available Philips Hue Bridge(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return discover(cmd, args, dd)
		},
	}
}

func discover(cmd *cobra.Command, args []string, dd Discoverer) error {
	bridges, err := dd.Discover()
	if err != nil {
		return err
	}

	for _, bridge := range bridges {
		_, err := fmt.Fprintf(cmd.OutOrStderr(), "Bridge Host/IP: %s (Id: %s)\n", bridge.Host, bridge.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
