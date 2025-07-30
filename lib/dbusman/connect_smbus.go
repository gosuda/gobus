package dbusman

import (
	"github.com/godbus/dbus"
)

type Units []struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	SubState    string
}

func ConnectDBus() (*dbus.Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return conn, err
}

func GetSystemd(conn *dbus.Conn) *Systemd {
	daemon := new(Systemd)
	daemon.bus = conn.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	return daemon
}
