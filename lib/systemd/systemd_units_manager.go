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
	ListUnitsByNames(names []string) ([]unit.Unit, error)
	ListUnits() ([]unit.Unit, error)
	ListUnitsByPatterns([]string, []string) ([]unit.UnitStatus, error)
	GetUnit(string) (d.ObjectPath, error)
	GetUnitByControlGroup(string) (d.ObjectPath, error)
	GetUnitByInvocationID([]byte) (d.ObjectPath, error)
	GetUnitByPID(uint32) (d.ObjectPath, error)
	GetUnitByPIDFD(*os.File) (d.ObjectPath, string, []byte, error)
	GetUnitFileLinks(string, bool) ([]string, error)
	GetUnitFileState(string) (string, error)
	GetUnitProcesses(string) ([]process.Process, error)
}

type SystemdUnitManager interface {
    KillUnit(string, string, int32)
    StartUnit(string, string) (d.ObjectPath, error)
    StopUnit(string, string) (d.ObjectPath, error)
}
