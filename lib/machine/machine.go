package machine

import "github.com/godbus/dbus"

const (
	FREEDESKTOP_MACHINE1 = "org.freedesktop.machine1"
	machineObjectPath    = "/org/freedesktop/machine1"
)

type Machine struct {
	bus dbus.BusObject
}

type MachineEntry struct {
	Name       string
	Class      string
	Service    string
	ObjectPath dbus.ObjectPath
}

type ImageEntry struct {
	Name         string
	Path         string
	ReadOnly     bool
	CreationUSec uint64
	ModifyUSec   uint64
	DiskUsage    uint64
	ObjectPath   dbus.ObjectPath
}

type MachineAddress struct {
	Family  int32
	Address []byte
}

func GetMachine(conn *dbus.Conn) *Machine {
	m := new(Machine)
	m.bus = conn.Object(FREEDESKTOP_MACHINE1, machineObjectPath)
	return m
}

func (m *Machine) ListMachines() ([]MachineEntry, error) {
	var rows [][]any
	call := m.bus.Call("org.freedesktop.machine1.Manager.ListMachines", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	machines := make([]MachineEntry, 0, len(rows))
	for _, row := range rows {
		if len(row) != 4 {
			continue
		}
		name, _ := row[0].(string)
		class, _ := row[1].(string)
		service, _ := row[2].(string)
		path, _ := row[3].(dbus.ObjectPath)
		machines = append(machines, MachineEntry{
			Name:       name,
			Class:      class,
			Service:    service,
			ObjectPath: path,
		})
	}
	return machines, nil
}

func (m *Machine) ListImages() ([]ImageEntry, error) {
	var rows [][]any
	call := m.bus.Call("org.freedesktop.machine1.Manager.ListImages", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	images := make([]ImageEntry, 0, len(rows))
	for _, row := range rows {
		if len(row) != 7 {
			continue
		}
		name, _ := row[0].(string)
		path, _ := row[1].(string)
		readOnly := toBool(row[2])
		creation := toUint64(row[3])
		modified := toUint64(row[4])
		disk := toUint64(row[5])
		obj, _ := row[6].(dbus.ObjectPath)
		images = append(images, ImageEntry{
			Name:         name,
			Path:         path,
			ReadOnly:     readOnly,
			CreationUSec: creation,
			ModifyUSec:   modified,
			DiskUsage:    disk,
			ObjectPath:   obj,
		})
	}
	return images, nil
}

func (m *Machine) GetMachine(name string) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetMachine", 0, name)
	err := call.Store(&path)
	return path, err
}

func (m *Machine) GetImage(name string) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetImage", 0, name)
	err := call.Store(&path)
	return path, err
}

func (m *Machine) GetMachineByPID(pid uint32) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetMachineByPID", 0, pid)
	err := call.Store(&path)
	return path, err
}

func (m *Machine) GetMachineAddresses(name string) ([]MachineAddress, error) {
	var rows [][]any
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetMachineAddresses", 0, name)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	addresses := make([]MachineAddress, 0, len(rows))
	for _, row := range rows {
		if len(row) != 2 {
			continue
		}
		family := toInt32(row[0])
		if bytes, ok := row[1].([]byte); ok {
			addrCopy := make([]byte, len(bytes))
			copy(addrCopy, bytes)
			addresses = append(addresses, MachineAddress{
				Family:  family,
				Address: addrCopy,
			})
		}
	}
	return addresses, nil
}

func (m *Machine) GetMachineSSHInfo(name string) (string, string, error) {
	var (
		address string
		keyPath string
	)
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetMachineSSHInfo", 0, name)
	if err := call.Store(&address, &keyPath); err != nil {
		return "", "", err
	}
	return address, keyPath, nil
}

func (m *Machine) GetMachineOSRelease(name string) (map[string]string, error) {
	var fields map[string]string
	call := m.bus.Call("org.freedesktop.machine1.Manager.GetMachineOSRelease", 0, name)
	if err := call.Store(&fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func toBool(value any) bool {
	switch v := value.(type) {
	case bool:
		return v
	default:
		return false
	}
}

func toUint64(value any) uint64 {
	switch v := value.(type) {
	case uint64:
		return v
	case uint32:
		return uint64(v)
	case int64:
		return uint64(v)
	case int:
		return uint64(v)
	default:
		return 0
	}
}

func toInt32(value any) int32 {
	switch v := value.(type) {
	case int32:
		return v
	case uint32:
		return int32(v)
	case int:
		return int32(v)
	default:
		return 0
	}
}
