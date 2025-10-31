package dbus

import (
	"github.com/godbus/dbus"
)

const (
	FREEDESKTOP_DBUS = "org.freedesktop.DBus"
	dbusObjectPath   = "/org/freedesktop/DBus"
)

type Bus struct {
	bus dbus.BusObject
}

func GetBus(conn *dbus.Conn) *Bus {
	b := new(Bus)
	b.bus = conn.Object(FREEDESKTOP_DBUS, dbusObjectPath)
	return b
}

func (b *Bus) ListNames() ([]string, error) {
	var names []string
	call := b.bus.Call("org.freedesktop.DBus.ListNames", 0)
	err := call.Store(&names)
	return names, err
}

func (b *Bus) ListActivatableNames() ([]string, error) {
	var names []string
	call := b.bus.Call("org.freedesktop.DBus.ListActivatableNames", 0)
	err := call.Store(&names)
	return names, err
}

func (b *Bus) GetNameOwner(name string) (string, error) {
	var owner string
	call := b.bus.Call("org.freedesktop.DBus.GetNameOwner", 0, name)
	err := call.Store(&owner)
	return owner, err
}

func (b *Bus) NameHasOwner(name string) (bool, error) {
	var hasOwner bool
	call := b.bus.Call("org.freedesktop.DBus.NameHasOwner", 0, name)
	err := call.Store(&hasOwner)
	return hasOwner, err
}

func (b *Bus) RequestName(name string, flags uint32) (uint32, error) {
	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.RequestName", 0, name, flags)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) ReleaseName(name string) (uint32, error) {
	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.ReleaseName", 0, name)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) StartServiceByName(name string, flags uint32) (uint32, error) {
	var result uint32
	call := b.bus.Call("org.freedesktop.DBus.StartServiceByName", 0, name, flags)
	err := call.Store(&result)
	return result, err
}

func (b *Bus) AddMatch(rule string) error {
	call := b.bus.Call("org.freedesktop.DBus.AddMatch", 0, rule)
	return call.Err
}

func (b *Bus) RemoveMatch(rule string) error {
	call := b.bus.Call("org.freedesktop.DBus.RemoveMatch", 0, rule)
	return call.Err
}
