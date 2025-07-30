package dbusman

import (
	"os"

	"github.com/godbus/dbus"
)

func (sysd *Systemd) GetUnits() []Unit {
	var units []Unit
	sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnits", 0).Store(&units)
	return units
}

func (sysd *Systemd) GetUnitsByNames(names []string) []Unit {
	var units []Unit
	sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, names).Store(&units)
	return units
}
func (sysd *Systemd) GetUnit(name string) dbus.ObjectPath {
	var unit dbus.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnit", 0, name).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitByControlGroup(cgroup string) dbus.ObjectPath {
	var unit dbus.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByControlGroup", 0, cgroup).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitByInvocationID(invocationId []byte) dbus.ObjectPath {
	var unit dbus.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByInvocationID", 0, invocationId).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitByPID(pid uint32) dbus.ObjectPath {
	var unit dbus.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPID", 0, pid).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitByPIDFD(pidfd *os.File) dbus.ObjectPath {
	var unit dbus.ObjectPath
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPID", 0, pidfd).Store(&unit)
	return unit
}

func (sysd *Systemd) GetUnitFileLinks(name string, runtime bool) []string {
	var links []string
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPid", 0, name, runtime).Store(&links)
	return links
}
func (sysd *Systemd) GetUnitFileState(file string) string {
	var state string
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetFileState", 0, file).Store(&state)
	return state
}

func (sysd *Systemd) GetUnitProcesses(name string) []Process {
	var processes []Process
	sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitProcesses", 0, name).Store(&processes)
	return processes
}
