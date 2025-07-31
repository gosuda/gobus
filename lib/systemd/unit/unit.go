package unit

import (
	"github.com/godbus/dbus"
	_ "github.com/gosuda/gobus/lib/dbus"
)

// This is a service unit of freedesktop
type Unit struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	SubState    string
	Followed    dbus.ObjectPath
	Path        dbus.ObjectPath
	JobId       uint32
	JobType     string
	JobPath     dbus.ObjectPath
}
