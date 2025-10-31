package dbus

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

type Conn dbus.Conn

// Connect to System Bus
func ConnectSystemBus() (*dbus.Conn, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return conn, err
}

// Connect to Session Bus
func ConnectSessionBus() (*dbus.Conn, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return conn, err
}

// Connect privately to System Bus
func ConnectSystemBusPrivate() (*dbus.Conn, error) {
	conn, err := dbus.SystemBusPrivate()
	if err != nil {
		return nil, err
	}
	if err = conn.Auth(nil); err != nil {
		conn.Close()
		return nil, err
	}
	if err = conn.Hello(); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

// Connect privately to Session Bus
func ConnectSessionBusPrivate() (*dbus.Conn, error) {
	conn, err := dbus.SessionBusPrivate()
	if err != nil {
		return nil, err
	}
	if err = conn.Auth(nil); err != nil {
		conn.Close()
		return nil, err
	}
	if err = conn.Hello(); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}
