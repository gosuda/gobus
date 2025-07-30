package process

func (p Process) GetCgroup() string {
	return p.cgroup
}

func (p Process) GetPid() uint32 {
	return p.pid
}

func (p Process) GetCommand() string {
	return p.command
}
// A *nix Process.
type Process struct {
	cgroup  string
	pid     uint32
	command string
}

// Do not directly access an information.
type ProcessInformation interface {
	GetCgroup() string
	GetPid() uint32
	GetCommand() string
}

