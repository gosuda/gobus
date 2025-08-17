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

	currentJob, affectedJobs, err := sysd.EnqueueUnitJob(jobName, jobType, jobMode)
    if err != nil {
        log.Printf("Error on EnqueueUnitJob: (%v)\n", err)
    }
	log.Printf("Current Job: %+v", currentJob)
	for i, j := range affectedJobs {
		log.Printf("Affected Job %d: %+v", i, j)
	}
    markedJobs := sysd.EnqueueMarkedJobs()
    log.Println("Marked Jobs: ", len(markedJobs))
    for _, m := range markedJobs {
        log.Printf("Marked Job: %v\n", m)
    }

}
