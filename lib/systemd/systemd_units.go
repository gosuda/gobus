package systemd

import (
	"os"

	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/process"
	"github.com/gosuda/gobus/lib/systemd/unit"
)

// Get All Units
func (sysd *Systemd) ListUnits() ([]unit.Unit, error) {
	var units []unit.Unit
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnits", 0)
    err := call.Store(&units)

	return units, err
}

// List units in a range of specified names
func (sysd *Systemd) ListUnitsByNames(names []string) ([]unit.Unit, error) {
	var units []unit.Unit
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, names)
    err := call.Store(&units)

	return units, err
}

func (sysd *Systemd) ListUnitsByPatterns(states []string, patterns []string) ([]unit.UnitStatus, error) {
	var unitStatus []unit.UnitStatus
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.ListUnitsByPatterns", 0, states, patterns)
    err := call.Store(&unitStatus)

	return unitStatus, err
}

// get a single unit by a name
func (sysd *Systemd) GetUnit(name string) (d.ObjectPath, error) {
	var unit d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnit", 0, name)
    err := call.Store(&unit)

	return unit, err
}

// get a single unit by the cgroup
func (sysd *Systemd) GetUnitByControlGroup(cgroup string) (d.ObjectPath, error) {
	var unit d.ObjectPath
	call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByControlGroup", 0, cgroup)
    err := call.Store(&unit)

	return unit, err
}

// get unit by invocation id. this gets []byte(not string)
func (sysd *Systemd) GetUnitByInvocationID(invocationId []byte) (d.ObjectPath, error) {
	var unit d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByInvocationID", 0, invocationId)
    err := call.Store(&unit)

	return unit, err
}

// get unit by pid
// see process module to see pid getter's declaration
func (sysd *Systemd) GetUnitByPID(pid uint32) (d.ObjectPath, error) {
	var unit d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPID", 0, pid)
    err := call.Store(&unit)

	return unit, err
}

func (sysd *Systemd) GetUnitByPIDFD(pidfd *os.File) (d.ObjectPath, string, []byte, error) {
    var unit d.ObjectPath
    var unit_id string
    var invocationId []byte

    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitByPIDFD", 0, d.UnixFD(pidfd.Fd()))
    if call.Err != nil {
        return "", "", nil, call.Err
    }
    err := call.Store(&unit, &unit_id, &invocationId)

    return unit, unit_id, invocationId, err
}

// get file links
func (sysd *Systemd) GetUnitFileLinks(name string, runtime bool) ([]string, error) {
	var links []string
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitFileLinks", 0, name, runtime)
    err := call.Store(&links)

	return links, err
}

// get file state while getting file path
func (sysd *Systemd) GetUnitFileState(file string) (string, error) {
	var state string
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetFileState", 0, file)
    err := call.Store(&state)

	return state, err
}

func (sysd *Systemd) GetUnitProcesses(name string) ([]process.Process, error) {
    var results [][]any
	var processes []process.Process
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.GetUnitProcesses", 0, name)
    err := call.Store(&results)
   	for _, result := range results {
		if len(result) == 3 {
			p := process.NewProcess(
				    result[0].(string),
				    result[1].(uint32),
				    result[2].(string),
                )
			processes = append(processes, p)
		}
	}

	return processes, err
}
