package systemd

import (
	"os"

	d "github.com/godbus/dbus"
	"github.com/gosuda/gobus/lib/systemd/process"
	"github.com/gosuda/gobus/lib/systemd/unit"
)

type SystemdUnitGetters interface {
	// Functions should follow FreeDesktop functions' original names.
	// These funtions are unit getters
	ListUnitsByNames(names []string) []unit.Unit
	ListUnits() []unit.Unit
	ListUnitsByPatterns([]string, []string) []unit.UnitStatus
	GetUnit(string) d.ObjectPath
	GetUnitByControlGroup(string) d.ObjectPath
	GetUnitByInvocationID([]byte) d.ObjectPath
	GetUnitByPID(uint32) d.ObjectPath
	GetUnitByPIDFD(*os.File) d.ObjectPath
	GetUnitFileLinks(string, bool) []string
	GetUnitFileState(string) string
	GetUnitProcesses(string) []process.Process
}
