package systemd

import (
	"os"

	d "github.com/godbus/dbus"
	process "github.com/gosuda/gobus/lib/systemd/process"
)

type SystemdUnitGetters interface {
	// Functions should follow FreeDesktop functions' original names.
	// These funtions are unit getters
	GetUnit(string) d.ObjectPath
	GetUnitByControlGroup(string) d.ObjectPath
	GetUnitByInvocationID([]byte) d.ObjectPath
	GetUnitByPID(uint32) d.ObjectPath
	GetUnitByPIDFD(*os.File) d.ObjectPath
	GetUnitFileLinks(string, bool) []string
	GetUnitFileState(string) string
	GetUnitProcesses(string) []process.Process
}
