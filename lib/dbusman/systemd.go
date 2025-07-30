package dbusman

import (
	"os"

	"github.com/godbus/dbus"
)

type Systemd struct {
	bus dbus.BusObject
}
type SystemdInterface interface {
	GetUnit(string) dbus.ObjectPath
	GetUnitByControlGroup(string) dbus.ObjectPath
	GetUnitByInvocationID([]byte) dbus.ObjectPath
	GetUnitByPID(uint32) dbus.ObjectPath
	GetUnitByPIDFD(*os.File) dbus.ObjectPath
	GetUnitFileLinks(string, bool) []string
	GetUnitFileState(string) string
	GetUnitProcesses(string) []Process
}

type Process struct {
	cgroup  string
	pid     uint32
	command string
}

type ProcessInformation interface {
	GetCgroup() string
	GetPid() uint32
	GetCommand() string
}

type Unit struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	SubState    string
	Followed    dbus.ObjectPath
	Path        dbus.ObjectPath
	JobId       uint32
	JobType     string
	JobPath     dbus.ObjectPath
}

type Opt string

// for optional args
