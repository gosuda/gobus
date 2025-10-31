package timedate

import "github.com/godbus/dbus"

const (
	FREEDESKTOP_TIMEDATE1 = "org.freedesktop.timedate1"
	timedateObjectPath    = "/org/freedesktop/timedate1"
)

type Timedate struct {
	bus dbus.BusObject
}

type TimedateProperties struct {
	Timezone        string
	LocalRTC        bool
	CanNTP          bool
	NTP             bool
	NTPSynchronized bool
	TimeUSec        uint64
	RTCTimeUSec     uint64
	CanSetTime      bool
	CanSetTimezone  bool
	CanSetLocalRTC  bool
}

func GetTimedate(conn *dbus.Conn) *Timedate {
	td := new(Timedate)
	td.bus = conn.Object(FREEDESKTOP_TIMEDATE1, timedateObjectPath)
	return td
}

func (td *Timedate) GetProperties() (TimedateProperties, error) {
	props, err := td.getAll()
	if err != nil {
		return TimedateProperties{}, err
	}
	return TimedateProperties{
		Timezone:        getString(props, "Timezone"),
		LocalRTC:        getBool(props, "LocalRTC"),
		CanNTP:          getBool(props, "CanNTP"),
		NTP:             getBool(props, "NTP"),
		NTPSynchronized: getBool(props, "NTPSynchronized"),
		TimeUSec:        getUint64(props, "TimeUSec"),
		RTCTimeUSec:     getUint64(props, "RTCTimeUSec"),
		CanSetTime:      getBool(props, "CanSetTime"),
		CanSetTimezone:  getBool(props, "CanSetTimezone"),
		CanSetLocalRTC:  getBool(props, "CanSetLocalRTC"),
	}, nil
}

func (td *Timedate) SetTime(usec int64, relative bool, interactive bool) error {
	call := td.bus.Call("org.freedesktop.timedate1.SetTime", 0, usec, relative, interactive)
	return call.Err
}

func (td *Timedate) SetTimezone(timezone string, interactive bool) error {
	call := td.bus.Call("org.freedesktop.timedate1.SetTimezone", 0, timezone, interactive)
	return call.Err
}

func (td *Timedate) SetLocalRTC(local bool, fixSystem bool, interactive bool) error {
	call := td.bus.Call("org.freedesktop.timedate1.SetLocalRTC", 0, local, fixSystem, interactive)
	return call.Err
}

func (td *Timedate) SetNTP(enable bool, interactive bool) error {
	call := td.bus.Call("org.freedesktop.timedate1.SetNTP", 0, enable, interactive)
	return call.Err
}

func (td *Timedate) getAll() (map[string]dbus.Variant, error) {
	call := td.bus.Call("org.freedesktop.DBus.Properties.GetAll", 0, FREEDESKTOP_TIMEDATE1)
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

func getBool(props map[string]dbus.Variant, key string) bool {
	if v, ok := props[key]; ok {
		if value, ok := v.Value().(bool); ok {
			return value
		}
	}
	return false
}

func getUint64(props map[string]dbus.Variant, key string) uint64 {
	if v, ok := props[key]; ok {
		if value, ok := v.Value().(uint64); ok {
			return value
		}
	}
	return 0
}
