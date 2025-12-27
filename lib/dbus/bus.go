package dbus

import (
	"github.com/godbus/dbus"
	"errors"
)

const (
	FREEDESKTOP_DBUS             = "org.freedesktop.DBus"
	FREEDESKTOP_NETWORKMANAGER   = "org.freedesktop.NetworkManager"
	dbusObjectPath               = "/org/freedesktop/DBus"
	networkManagerObjectPath     = "/org/freedesktop/NetworkManager"
	KIND_FREEDESKTOP_DBUS        = "__KIND_FREEDESKTOP_DBUS__"
	KIND_FREEDESKTOP_NM          = "__KIND_FREEDESKTOP_NM__"
)

type Bus struct {
	bus  dbus.BusObject
	kind string
}

type NmSettings struct {
	variant dbus.Variant
}

/* start of networkmanager functions */

func GetNm(conn *dbus.Conn) *Bus {
	b := new(Bus)
	b.bus = conn.Object(FREEDESKTOP_NETWORKMANAGER, networkManagerObjectPath)
	b.kind = KIND_FREEDESKTOP_NM
	return b
}

func (b *Bus) GetNmSettings(conn *dbus.Conn, settings *NmSettings) error {
	if b.kind != KIND_FREEDESKTOP_NM {
		return errors.New("This bus is not NetworkManager bus.")
	}
	call := b.bus.Call("org.freedesktop.NetworkManager.Settings.Get", 0)
	err := call.Store(&settings.variant)
	return err
}
/* start of dbus functions */

func GetBus(conn *dbus.Conn) *Bus {
	b := new(Bus)
	b.bus  = conn.Object(FREEDESKTOP_DBUS, dbusObjectPath)
	b.kind = KIND_FREEDESKTOP_DBUS
	return b
}

func (b *Bus) ListNames() ([]string, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return []string{}, errors.New("This bus is not DBus")
	}

	var names []string
	call := b.bus.Call("org.freedesktop.DBus.ListNames", 0)
	err := call.Store(&names)
	return names, err
}

func (b *Bus) ListActivatableNames() ([]string, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return []string{}, errors.New("This bus is not DBus")
	}

	var names []string
	call := b.bus.Call("org.freedesktop.DBus.ListActivatableNames", 0)
	err := call.Store(&names)
	return names, err
}

func (b *Bus) GetNameOwner(name string) (string, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return "", errors.New("This bus is not DBus")
	}

	var owner string
	call := b.bus.Call("org.freedesktop.DBus.GetNameOwner", 0, name)
	err := call.Store(&owner)
	return owner, err
}

func (b *Bus) NameHasOwner(name string) (bool, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return false, errors.New("This bus is not DBus")
	}

	var hasOwner bool
	call := b.bus.Call("org.freedesktop.DBus.NameHasOwner", 0, name)
	err := call.Store(&hasOwner)
	return hasOwner, err
}

func (b *Bus) RequestName(name string, flags uint32) (uint32, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return 0, errors.New("This bus is not DBus")
	}

	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.RequestName", 0, name, flags)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) ReleaseName(name string) (uint32, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return 0, errors.New("This bus is not DBus")
	}

	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.ReleaseName", 0, name)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) StartServiceByName(name string, flags uint32) (uint32, error) {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return 0, errors.New("This bus is not DBus")
	}

	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.StartServiceByName", 0, name, flags)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) AddMatch(rule string) error {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return errors.New("This bus is not DBus")
	}

	call := b.bus.Call("org.freedesktop.DBus.AddMatch", 0, rule)
	return call.Err
}

func (b *Bus) RemoveMatch(rule string) error {
	if b.kind != KIND_FREEDESKTOP_DBUS {
		return errors.New("This bus is not DBus")
	}

	call := b.bus.Call("org.freedesktop.DBus.RemoveMatch", 0, rule)
	return call.Err
}
