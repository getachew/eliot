package printers

import (
	"fmt"
	"html/template"
	"io"
	"log"

	pb "github.com/ernoaapa/can/pkg/api/services/pods/v1"
	"github.com/ernoaapa/can/pkg/config"
	"github.com/ernoaapa/can/pkg/model"
	"github.com/ernoaapa/can/pkg/printers/humanreadable"
)

// HumanReadablePrinter is an implementation of ResourcePrinter which prints
// resources in human readable format (tables etc.).
type HumanReadablePrinter struct {
}

// NewHumanReadablePrinter creates new HumanReadablePrinter
func NewHumanReadablePrinter() *HumanReadablePrinter {
	return &HumanReadablePrinter{}
}

// PrintPodsTable writes list of Pods in human readable table format to the writer
func (p *HumanReadablePrinter) PrintPodsTable(pods []*pb.Pod, writer io.Writer) error {
	fmt.Fprintln(writer, humanreadable.PodsTableHeader)
	t := template.New("pods-table")
	t, err := t.Parse(humanreadable.PodsTableRowTemplate)
	if err != nil {
		log.Fatalf("Invalid pod template: %s", err)
	}

	for _, pod := range pods {
		if err := t.Execute(writer, pod); err != nil {
			return err
		}
	}
	// TODO: For some reason, don't output without printing something to the writer
	// Find out how to flush the writer
	fmt.Fprintln(writer, "\n")
	return nil
}

// PrintDevicesTable writes stream of Devices in human readable table format to the writer
func (p *HumanReadablePrinter) PrintDevicesTable(devices <-chan model.DeviceInfo, writer io.Writer) error {
	fmt.Fprintln(writer, "HOSTNAME\tENDPOINT")

	go func(devices <-chan model.DeviceInfo) {
		for device := range devices {
			fmt.Fprintf(writer, "%s\t%s", device.Hostname, device.GetPrimaryEndpoint())
			fmt.Fprintln(writer, "\n")
		}
	}(devices)
	return nil
}

// PrintPodDetails writes list of pods in human readable detailed format to the writer
func (p *HumanReadablePrinter) PrintPodDetails(pod *pb.Pod, writer io.Writer) error {
	t := template.New("pod-details")
	t, err := t.Parse(humanreadable.PodDetailsTemplate)
	if err != nil {
		log.Fatalf("Invalid pod template: %s", err)
	}

	if err := t.Execute(writer, pod); err != nil {
		return err
	}
	return nil
}

// PrintConfig writes list of pods in human readable detailed format to the writer
func (p *HumanReadablePrinter) PrintConfig(config *config.Config, writer io.Writer) error {
	t := template.New("config")
	t, err := t.Parse(humanreadable.ConfigTemplate)
	if err != nil {
		log.Fatalf("Invalid config template: %s", err)
	}

	if err := t.Execute(writer, config); err != nil {
		return err
	}

	fmt.Fprintln(writer, "\n")
	return nil
}
