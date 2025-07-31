package main

import (
	"log"

	"github.com/gosuda/gobus/lib/dbus"
	"github.com/gosuda/gobus/lib/systemd"
)

func main() {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalf("Error connecting DBus: (%v)", err)
	}
	sysd := systemd.GetSystemd(conn)
	jobName := "cups.service"
	jobType := "start"
	jobMode := "replace"

	currentJob, affectedJobs := sysd.EnqueueUnitJob(jobName, jobType, jobMode)
	log.Printf("Current Job: %+v", currentJob)
	for i, j := range affectedJobs {
		log.Printf("Affected Job %d: %+v", i, j)
	}

}
