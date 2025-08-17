package main

import (
	"fmt"
	"log"

	"github.com/gosuda/gobus/lib/dbus"
	"github.com/gosuda/gobus/lib/systemd"
)

func main() {
	// Connect to the system bus
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalf("Failed to connect to system bus: %v", err)
	}

	// Create a systemd manager instance
	sysd := systemd.GetSystemd(conn)

	// Fetch all units
	units, err := sysd.ListUnits()
    if err != nil {
        log.Fatalf("[CRITICAL] Error on loading all units: (%v)\n", err)
    }

	// Print the basic info of each unit
	fmt.Println("List of units:")
	for _, unit := range units {
		fmt.Printf("Name: %s, Description: %s, LoadState: %s, ActiveState: %s, SubState: %s\n",
			unit.Name, unit.Description, unit.LoadState, unit.ActiveState, unit.SubState)
	}
	fmt.Println("Filtered Units:")
	serviceNames := []string{"systemd-journald.service", "network.service"}
	filteredUnits, err := sysd.ListUnitsByNames(serviceNames)
    if err != nil {
        log.Fatalf("Failed to list units by name: (%v)", err)
        for _, service := range serviceNames {
            log.Printf("%s service not found.", service)
        }
    }
	for _, unit := range filteredUnits {
		fmt.Printf("Name: %s, Description: %s, LoadState: %s, ActiveState: %s, SubState: %s\n",
			unit.Name, unit.Description, unit.LoadState, unit.ActiveState, unit.SubState)
	}

    cur := "systemd-journald.service"
	// Example: Get a specific unit by name
	unitPath, err := sysd.GetUnit(cur)
    if err != nil {
        log.Fatalf("Failed to get unit %s: error(%v)", cur, err)
    }
	fmt.Printf("\nObject path of systemd-journald.service: %s\n", unitPath)

	// Example: Get processes of the unit
	processes, err := sysd.GetUnitProcesses(cur)
    if err != nil {
        log.Fatalf("Failed to get process from %s: error(%v)", cur, err)
    }

	fmt.Println("\nProcesses running under ", cur)
	for _, p := range processes {
		fmt.Printf("PID: %d, Command: %s, Cgroup: %s\n", p.GetPid(), p.GetCommand(), p.GetCgroup())
	}
}
