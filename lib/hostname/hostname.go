package hostname

import (
	"github.com/godbus/dbus"
)

const (
	FREEDESKTOP_HOSTNAME1 = "org.freedesktop.hostname1"
	hostnameObjectPath    = "/org/freedesktop/hostname1"
)

type Hostname struct {
	bus dbus.BusObject
}

type HostnameProperties struct {
	Hostname                  string
	StaticHostname            string
	PrettyHostname            string
	IconName                  string
	Chassis                   string
	Deployment                string
	Location                  string
	KernelName                string
	KernelRelease             string
	KernelVersion             string
	OperatingSystemPrettyName string
	OperatingSystemCPEName    string
	HardwareVendor            string
	HardwareModel             string
}

func GetHostname(conn *dbus.Conn) *Hostname {
	host := new(Hostname)
	host.bus = conn.Object(FREEDESKTOP_HOSTNAME1, hostnameObjectPath)
	return host
}

func (h *Hostname) GetProperties() (HostnameProperties, error) {
	props, err := h.getAll()
	if err != nil {
		return HostnameProperties{}, err
	}
	return HostnameProperties{
		Hostname:                  getString(props, "Hostname"),
		StaticHostname:            getString(props, "StaticHostname"),
		PrettyHostname:            getString(props, "PrettyHostname"),
		IconName:                  getString(props, "IconName"),
		Chassis:                   getString(props, "Chassis"),
		Deployment:                getString(props, "Deployment"),
		Location:                  getString(props, "Location"),
		KernelName:                getString(props, "KernelName"),
		KernelRelease:             getString(props, "KernelRelease"),
		KernelVersion:             getString(props, "KernelVersion"),
		OperatingSystemPrettyName: getString(props, "OperatingSystemPrettyName"),
		OperatingSystemCPEName:    getString(props, "OperatingSystemCPEName"),
		HardwareVendor:            getString(props, "HardwareVendor"),
		HardwareModel:             getString(props, "HardwareModel"),
	}, nil
}

func (h *Hostname) SetHostname(name string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetHostname", 0, name, interactive)
	return call.Err
}

func (h *Hostname) SetStaticHostname(name string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetStaticHostname", 0, name, interactive)
	return call.Err
}

func (h *Hostname) SetPrettyHostname(name string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetPrettyHostname", 0, name, interactive)
	return call.Err
}

func (h *Hostname) SetIconName(icon string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetIconName", 0, icon, interactive)
	return call.Err
}

func (h *Hostname) SetChassis(chassis string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetChassis", 0, chassis, interactive)
	return call.Err
}

func (h *Hostname) SetDeployment(deployment string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetDeployment", 0, deployment, interactive)
	return call.Err
}

func (h *Hostname) SetLocation(location string, interactive bool) error {
	call := h.bus.Call("org.freedesktop.hostname1.SetLocation", 0, location, interactive)
	return call.Err
}

func (h *Hostname) getAll() (map[string]dbus.Variant, error) {
	call := h.bus.Call("org.freedesktop.DBus.Properties.GetAll", 0, FREEDESKTOP_HOSTNAME1)
	if call.Err != nil {
		return nil, call.Err
	}
	var props map[string]dbus.Variant
	if err := call.Store(&props); err != nil {
		return nil, err
	}
	return props, nil
}

func getString(props map[string]dbus.Variant, key string) string {
	if v, ok := props[key]; ok {
		if value, ok := v.Value().(string); ok {
			return value
		}
	}
	return ""
}
