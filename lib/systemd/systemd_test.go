package systemd

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"

	"testing"

	"github.com/gosuda/gobus/lib/dbus"
)

var sysd *Systemd
var err error
const SIGTERM = 15

func TestMain(m *testing.M) {
    conn, err := dbus.ConnectSystemBus()
	if err != nil {
		os.Exit(1)
	}
	sysd = GetSystemd(conn)
	defer conn.Close()
	os.Exit(m.Run())
}

func TestListUnits(t *testing.T) {
	units, err := sysd.ListUnits()
	if err != nil {
		t.Fatalf("Error on ListUnits: %v", err)
	}
	if len(units) == 0 {
		t.Fatal("Expected at least one unit from ListUnits()")
	}
}

func TestListUnitsByNames(t *testing.T) {
	names := []string{"systemd-journald.service", "dbus.service"}
	unitsByName, err := sysd.ListUnitsByNames(names)
	if err != nil {
		t.Fatalf("Error on ListUnitsByNames: %v", err)
	}
	if len(unitsByName) != len(names) {
		t.Fatalf("Expected %d units from ListUnitsByNames(), got %d", len(names), len(unitsByName))
	}
}

func TestGetUnit(t *testing.T) {
	unitPath, err := sysd.GetUnit("systemd-journald.service")
	if err != nil {
		t.Fatalf("Error on GetUnit: %v", err)
	}
	if unitPath == "" {
		t.Fatal("GetUnit returned empty ObjectPath")
	}
}

func TestGetUnitByControlGroup(t *testing.T) {
	unitByCgroup, err := sysd.GetUnitByControlGroup("/system.slice/systemd-journald.service")
	if err != nil {
		t.Fatalf("Error on GetUnitByControlGroup: %v", err)
	}
	if unitByCgroup == "" {
		t.Fatal("GetUnitByControlGroup returned empty ObjectPath")
	}
}

func TestGetUnitByPID(t *testing.T) {
	unitByPID, err := sysd.GetUnitByPID(1)
	if err != nil {
		t.Fatalf("Error on GetUnitByPID: %v", err)
	}
	if unitByPID == "" {
		t.Log("GetUnitByPID returned empty ObjectPath")
	}
}

func TestGetUnitByInvocationID(t *testing.T) {
	unitByInvocation, err := sysd.GetUnitByInvocationID([]byte{})
	if err != nil {
		t.Fatalf("Error on GetUnitByInvocationID: %v", err)
	}
	if unitByInvocation == "" {
		t.Log("GetUnitByInvocationID returned empty ObjectPath")
	}
}

func TestGetUnitByPIDFD_CorrectScenario(t *testing.T) {
	// 1. Start a new process to get a valid PID.
	cmd := exec.Command("sleep", "100")
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start process for test: %v", err)
	}
	// 2. Ensure the process is killed after the test.
	defer cmd.Process.Kill()

	// 3. Get the PID of the new process.
	pid := cmd.Process.Pid
	if pid == 0 {
		t.Fatal("Failed to get PID from the newly created process")
	}

	// 4. Use unix.PidfdOpen to get a pidfd from the PID.
	// This is the special file descriptor type required by systemd.
	pidfd, err := unix.PidfdOpen(pid, 0)
	if err != nil {
		t.Fatalf("Failed to open pidfd for PID %d: %v", pid, err)
	}
	// 5. Ensure the pidfd is closed after the test.
	defer unix.Close(pidfd)
	
	// 6. Wrap the pidfd integer in an *os.File object.
	pidFile := os.NewFile(uintptr(pidfd), "pidfd")
	defer pidFile.Close()

	// 7. Test the function with the valid pidfd.
	unitByPIDFD, _, _, err := sysd.GetUnitByPIDFD(pidFile)
	if err != nil {
		t.Fatalf("Error on GetUnitByPIDFD with a valid pidfd: %v", err)
	}

	if unitByPIDFD == "" {
		t.Log("GetUnitByPIDFD returned empty ObjectPath, as expected for a transient service")
	}
}

func TestGetUnitFileLinks(t *testing.T) {
	links, err := sysd.GetUnitFileLinks("systemd-journald.service", false)
	if err != nil {
		t.Fatalf("Error on GetUnitFileLinks: %v", err)
	}
	if len(links) == 0 {
		t.Log("GetUnitFileLinks returned empty slice")
	}
}

func TestGetUnitProcesses(t *testing.T) {
	processes, err := sysd.GetUnitProcesses("systemd-journald.service")
	if err != nil {
		t.Fatalf("Error on GetUnitProcesses: %v", err)
	}
	if len(processes) == 0 {
		t.Log("GetUnitProcesses returned empty slice")
	}
}

func TestSystemdStartUnit(t *testing.T) {
    j, err := sysd.StartUnit("cups.service", "fail")
    if err != nil {
        t.Fatalf("Error on starting unit: (%v)", err)
    }
    t.Logf("Started unit: (%v)", j)
}

func TestSystemdStopUnit(t *testing.T) {
    for {
        units ,err := sysd.ListUnitsByNames([]string{"cups.service"})
        if err != nil {
            t.Fatalf("Failed to list units: cups.service not found. Error: (%v)", err)
        }
        if units[0].SubState == "running" {
            break
        }
    }
    j, err := sysd.StopUnit("cups.service", "fail") 
    if err != nil {
        t.Fatalf("Error on stopping unit: (%v)", err)
    }
    t.Logf("Stopped unit: (%v)", j)

}

func TestSystemdKillUnit(t *testing.T) {
    sysd.KillUnit("cups.service", "main", SIGTERM)
    
}

func TestSystemdJobControllers(t *testing.T) {
	j, _, err := sysd.EnqueueUnitJob("cups.service", "start", "replace")
	if err != nil {
		t.Fatalf("Error starting a job: %v", err)
	}
	_, err = sysd.GetJob(j.JobId)
	if err != nil {
		t.Fatalf("Error getting job after enqueue: %v", err)
	}
}


