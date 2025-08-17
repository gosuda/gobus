package systemd

import (
	"github.com/godbus/dbus"
)

type Systemd struct {
	bus dbus.BusObject
}

type Object Systemd
type UnitBus Systemd

const (
    FREEDESKTOP_SYSTEMD1 = "org.freedesktop.systemd1"
)

type SystemdManager interface {
    GetSystemd(*dbus.Conn) *Systemd
    GetObject(*dbus.Conn, string) *Object 
}

// Get Systemd from System Bus
func GetSystemd(conn *dbus.Conn) *Systemd {
	daemon := new(Systemd)
	daemon.bus = conn.Object(FREEDESKTOP_SYSTEMD1, "/org/freedesktop/systemd1")
	return daemon
}

// Get Object from Systemd
func GetObject(conn *dbus.Conn, unit dbus.ObjectPath) *Object {
    daemon := new(Object)
    daemon.bus = conn.Object(FREEDESKTOP_SYSTEMD1, unit)
    return daemon
}

type Opt string

// for optional args
