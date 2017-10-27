package main

import (
	"os"
	"time"

	"github.com/ernoaapa/can/cmd"
	"github.com/ernoaapa/can/pkg/discovery"
	"github.com/ernoaapa/can/pkg/display"
	"github.com/ernoaapa/can/pkg/printers"
	"github.com/urfave/cli"
)

var getDevicesCommand = cli.Command{
	Name:    "devices",
	Aliases: []string{"device"},
	Usage:   "Get Device resources",
	UsageText: `canctl get devices [options]
			 
	 # Get table of known devices
	 canctl get devices`,
	Action: func(clicontext *cli.Context) error {
		display := display.NewLine()
		display.Active("Discover from network automatically...")

		devices, err := discovery.Devices(5 * time.Second)
		if err != nil {
			display.Errorf("Failed to auto-discover devices in network: %s", err)
			return err
		}
		display.Donef("Discovered %d devices from network", len(devices))

		writer := printers.GetNewTabWriter(os.Stdout)
		defer writer.Flush()

		printer := cmd.GetPrinter(clicontext)
		printer.PrintDevicesTable(devices, writer)

		return nil
	},
}
