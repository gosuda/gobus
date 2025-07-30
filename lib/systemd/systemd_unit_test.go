package systemd

import (
	"os"
	"testing"

	"github.com/gosuda/gobus/lib/dbusman"
)

func TestDBusSystemdFunctions(t *testing.T) {
	conn, err := dbusman.ConnectDBus()
	if err != nil {
		t.Fatalf("Failed to connect to DBus: %v", err)
	}
	sysd := GetSystemd(conn)

	// 1. Test ListUnits - expect at least one unit returned
	units := sysd.ListUnits()
	if len(units) == 0 {
		t.Fatal("Expected at least one unit from GetUnits()")

		t.Logf("GetUnits returned %d units", len(units))
	}

	// 2. Test ListUnitsByNames with existing service names
	names := []string{"systemd-journald.service", "dbus.service"}
	unitsByName := sysd.ListUnitsByNames(names)
	if len(unitsByName) != len(names) {
		t.Fatalf("Expected %d units from GetUnitsByNames(), got %d", len(names), len(unitsByName))
	}
	for i, u := range unitsByName {
		t.Logf("Unit %d: %s (ActiveState: %s)", i, u.Name, u.ActiveState)
	}

	// 3. Test GetUnit by service name
	unitPath := sysd.GetUnit("systemd-journald.service")
	if unitPath == "" {
		t.Fatal("GetUnit returned empty ObjectPath")
	}
	t.Logf("GetUnit returned: %s", unitPath)

	// 4. Test GetUnitByControlGroup with typical cgroup path
	unitByCgroup := sysd.GetUnitByControlGroup("/system.slice/systemd-journald.service")
	if unitByCgroup == "" {
		t.Fatal("GetUnitByControlGroup returned empty ObjectPath")
	}
	t.Logf("GetUnitByControlGroup returned: %s", unitByCgroup)

	// 5. Test GetUnitByPID - using PID 1 as example
	unitByPID := sysd.GetUnitByPID(1)
	if unitByPID == "" {
		t.Log("GetUnitByPID returned empty ObjectPath - this may be expected if PID 1 is not a unit")
	} else {
		t.Logf("GetUnitByPID returned: %s", unitByPID)
	}

	// 6. Test GetUnitByInvocationID with empty invocation ID (likely no match)
	unitByInvocation := sysd.GetUnitByInvocationID([]byte{})
	if unitByInvocation == "" {
		t.Log("GetUnitByInvocationID returned empty ObjectPath - likely no invocation with empty ID")
	} else {
		t.Logf("GetUnitByInvocationID returned: %s", unitByInvocation)
	}

	// 7. Test GetUnitByPIDFD using /proc/self/fd/0 as PIDFD example
	fd, err := os.Open("/proc/self/fd/0")
	if err != nil {
		t.Logf("Could not open fd for GetUnitByPIDFD test: %v", err)
	} else {
		defer fd.Close()
		unitByPIDFD := sysd.GetUnitByPIDFD(fd)
		if unitByPIDFD == "" {
			t.Log("GetUnitByPIDFD returned empty ObjectPath")
		} else {
			t.Logf("GetUnitByPIDFD returned: %s", unitByPIDFD)
		}
	}

	// 8. Test GetUnitFileLinks for a service (runtime = false)
	links := sysd.GetUnitFileLinks("systemd-journald.service", false)
	if len(links) == 0 {
		t.Log("GetUnitFileLinks returned empty slice")
	} else {
		t.Logf("GetUnitFileLinks returned %d links", len(links))
	}

	// 9. Test GetUnitProcesses for a service
	processes := sysd.GetUnitProcesses("systemd-journald.service")
	if len(processes) == 0 {
		t.Log("GetUnitProcesses returned empty slice")
	} else {
		t.Logf("GetUnitProcesses returned %d processes", len(processes))
		for _, p := range processes {
			t.Logf("Process PID: %d, Command: %s", p.GetPid(), p.GetCommand())
		}
	}
}
