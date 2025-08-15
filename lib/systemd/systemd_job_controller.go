package systemd

import (
	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/job"
)

type SystemdJobController interface {
	// this should contain systemd job related methods
	CancelJob(uint32)
	ClearJobs()
	EnqueueUnitJob(string, string, string) (job.Job, []job.Job, error)
    EnqueueMarkedJobs() ([]d.ObjectPath, error)
	GetJob(uint32) (d.ObjectPath, error)
}
