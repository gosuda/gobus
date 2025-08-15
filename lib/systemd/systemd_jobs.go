package systemd

import (
	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/job"
)

func (sysd *Systemd) GetJob(id uint32) (d.ObjectPath, error) { // retrieves job as uint32 and returns object path
	var path d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetJob", 0, id)
    err := call.Store(&path)
	return path, err
}

func (sysd *Systemd) EnqueueUnitJob(name string, jobType string, jobMode string) (job.Job, []job.Job, error) {
	var (
		jobID      uint32
		jobPath    d.ObjectPath
		unitName   string
		unitPath   d.ObjectPath
		jobTypeRet string
		affected   []job.Job
	)

    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.EnqueueUnitJob", 0, name, jobType, jobMode)
    err := call.Store(&jobID, &jobPath, &unitName, &unitPath, &jobTypeRet, &affected)

	// Compose current job struct from returned fields
	currentJob := job.Job{
		JobId:    jobID,
		JobPath:  jobPath,
		UnitId:   unitName,
		UnitPath: unitPath,
		JobType:  jobTypeRet,
	}

	return currentJob, affected, err
}

// INTERNAL
// Do not use this function unless you need to handle internal calls
func (sysd *Systemd) EnqueueMarkedJobs() []d.ObjectPath {
    var marked []d.ObjectPath
    sysd.bus.Call("org.freedesktop.systemd1.manager.EnqueueMarkedJobs", 0).Store(&marked)
    return marked
}

func (sysd *Systemd) CancelJob(id uint32) {
    sysd.bus.Call("org.freedesktop.systemd1.Manager.CancelJob", 0, id) // no return value to store
}


