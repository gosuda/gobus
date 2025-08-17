package systemd

import (
    d "github.com/godbus/dbus"
)

func (sysd *Systemd) KillUnit(name string, whom string, signal int32) {
    // no return as freedesktop mentioned
    sysd.bus.Call("org.freedesktop.systemd1.Manager.KillUnit", 0, name, whom, signal)
}

func (sysd *Systemd) StartUnit(name string, mode string) (d.ObjectPath, error) {
    var job d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.StartUnit", 0, name, mode)
    err := call.Store(&job)
    return job, err
}

func (sysd *Systemd) StopUnit(name string, mode string) (d.ObjectPath, error)  {
    var job d.ObjectPath
    call := sysd.bus.Call("org.freedesktop.systemd1.Manager.StopUnit", 0, name, mode)
    err := call.Store(&job)
    return job, err
}
