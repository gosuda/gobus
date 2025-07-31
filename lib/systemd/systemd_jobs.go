package systemd

import (
	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/job"
)

func (sysd *Systemd) GetJob(id uint32) d.ObjectPath { // retrieves job as uint32 and returns object path
	var path d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetJob", 0, id).Store(&path)
	return path
}

func (sysd *Systemd) EnqueueUnitJob(name string, jobType string, jobMode string) (job.Job, []job.Job) {
	var (
		jobID      uint32
		jobPath    d.ObjectPath
		unitName   string
		unitPath   d.ObjectPath
		jobTypeRet string
		affected   []job.Job
	)

	sysd.bus.Call("org.freedesktop.systemd1.Manager.EnqueueUnitJob", 0, name, jobType, jobMode).
		Store(&jobID, &jobPath, &unitName, &unitPath, &jobTypeRet, &affected)

	// Compose current job struct from returned fields
	currentJob := job.Job{
		JobId:    jobID,
		JobPath:  jobPath,
		UnitId:   unitName,
		UnitPath: unitPath,
		JobType:  jobTypeRet,
	}

	return currentJob, affected
}
