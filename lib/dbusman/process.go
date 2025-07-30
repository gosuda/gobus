package dbusman

func (p Process) GetCgroup() string {
	return p.cgroup
}

func (p Process) GetPid() uint32 {
	return p.pid
}

func (p Process) GetCommand() string {
	return p.command
}
