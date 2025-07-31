package systemd

import (
	"os"

	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/process"
	"github.com/gosuda/gobus/lib/systemd/unit"
)

// Get All Units
func (sysd *Systemd) ListUnits() []unit.Unit {
	var units []unit.Unit
	sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnits", 0).Store(&units)
	return units
}

// List units in a range of specified names
func (sysd *Systemd) ListUnitsByNames(names []string) []unit.Unit {
	var units []unit.Unit
	sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, names).Store(&units)
	return units
}

// get a single unit by a name
func (sysd *Systemd) GetUnit(name string) d.ObjectPath {
	var unit d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnit", 0, name).Store(&unit)
	return unit
}

// get a single unit by the cgroup
func (sysd *Systemd) GetUnitByControlGroup(cgroup string) d.ObjectPath {
	var unit d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByControlGroup", 0, cgroup).Store(&unit)
	return unit
}

// get unit by invocation id. this gets []byte(not string)
func (sysd *Systemd) GetUnitByInvocationID(invocationId []byte) d.ObjectPath {
	var unit d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByInvocationID", 0, invocationId).Store(&unit)
	return unit
}

// get unit by pid.
// see process module to see pid getter's declaration
func (sysd *Systemd) GetUnitByPID(pid uint32) d.ObjectPath {
	var unit d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPID", 0, pid).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitByPIDFD(pidfd *os.File) d.ObjectPath {
	var unit d.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPIDFD", 0, pidfd).Store(&unit)
	return unit
}

// get file links
func (sysd *Systemd) GetUnitFileLinks(name string, runtime bool) []string {
	var links []string
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitFileLinks", 0, name, runtime).Store(&links)
	return links
}

// get file state while getting file path
func (sysd *Systemd) GetUnitFileState(file string) string {
	var state string
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetFileState", 0, file).Store(&state)
	return state
}

func (sysd *Systemd) GetUnitProcesses(name string) []process.Process {
	var processes []process.Process
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitProcesses", 0, name).Store(&processes)
	return processes
}
