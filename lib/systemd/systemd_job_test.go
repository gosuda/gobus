package systemd

import (
	"testing"

	"github.com/gosuda/gobus/lib/dbus"
)

func TestSystemdJobControllers(t *testing.T) {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		t.Fatalf("Error connecting DBus: (%v)", err)
	}
	sysd := GetSystemd(conn)
	j, _ := sysd.EnqueueUnitJob("cups.service", "start", "fail")
	_ = sysd.GetJob(j.JobId)
}
