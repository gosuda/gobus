package systemd
import (
	"github.com/godbus/dbus"
)


type Systemd struct {
	bus dbus.BusObject
}

// Get Systemd from System Bus
func GetSystemd(conn *dbus.Conn) *Systemd {
	daemon := new(Systemd)
	daemon.bus = conn.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	return daemon
}
