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

// Connect to System Bus
func ConnectDBus() (*dbus.Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return conn, err
}


