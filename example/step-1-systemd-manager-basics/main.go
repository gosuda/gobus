package main

import (
	"fmt"
	"log"

	"github.com/gosuda/gobus/lib/dbusman"
)

func main() {
	// Connect to the system bus
	conn, err := dbusman.ConnectDBus()
	if err != nil {
		log.Fatalf("Failed to connect to system bus: %v", err)
	}

	// Create a systemd manager instance
	sysd := dbusman.GetSystemd(conn)

	// Fetch all units
	units := sysd.ListUnits()

	// Print the basic info of each unit
	fmt.Println("List of units:")
	for _, unit := range units {
		fmt.Printf("Name: %s, Description: %s, LoadState: %s, ActiveState: %s, SubState: %s\n",
			unit.Name, unit.Description, unit.LoadState, unit.ActiveState, unit.SubState)
	}
	fmt.Println("Filtered Units:")
	serviceNames := []string{"systemd-journald.service", "network.service"}
	filteredUnits := sysd.ListUnitsByNames(serviceNames)
	for _, unit := range filteredUnits {
		fmt.Printf("Name: %s, Description: %s, LoadState: %s, ActiveState: %s, SubState: %s\n",
			unit.Name, unit.Description, unit.LoadState, unit.ActiveState, unit.SubState)
	}

	// Example: Get a specific unit by name
	unitPath := sysd.GetUnit("systemd-journald.service")
	fmt.Printf("\nObject path of systemd-journald.service: %s\n", unitPath)

	// Example: Get processes of the unit
	processes := sysd.GetUnitProcesses("systemd-journald.service")
	fmt.Println("\nProcesses running under systemd-journald.service:")
	for _, p := range processes {
		fmt.Printf("PID: %d, Command: %s, Cgroup: %s\n", p.GetPid(), p.GetCommand(), p.GetCgroup())
	}
}

