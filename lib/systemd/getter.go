package systemd
import (
	"os"
	_ "github.com/gosuda/gobus/lib/dbusman"
	"github.com/godbus/dbus"
	process "github.com/gosuda/gobus/lib/dbusman/process"
)

type SystemdUnitGetters interface {
	// Functions should follow FreeDesktop functions' original names.
	// These funtions are unit getters
	GetUnit(string) dbus.ObjectPath
	GetUnitByControlGroup(string) dbus.ObjectPath
	GetUnitByInvocationID([]byte) dbus.ObjectPath
	GetUnitByPID(uint32) dbus.ObjectPath
	GetUnitByPIDFD(*os.File) dbus.ObjectPath
	GetUnitFileLinks(string, bool) []string
	GetUnitFileState(string) string
	GetUnitProcesses(string) []process.Process
}
