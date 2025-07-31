package job

import (
	d "github.com/godbus/dbus"
)

type Job struct {
	JobId     uint32
	JobPath   d.ObjectPath
	UnitId    string
	UnitPath  d.ObjectPath
	JobType   string
	JobStatus uint32
}
