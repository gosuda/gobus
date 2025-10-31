## Project Tree
```bash
.
├── LICENSE
├── README.md
├── TREE.md
├── assets
│   └── icon.png
├── example
│   └── step-1
│       ├── systemd-manager-jobs-basics
│       │   └── main.go
│       └── systemd-manager-unit-basics
│           └── main.go
├── go.mod
├── go.sum
├── lib
│   ├── dbus
│   │   ├── bus.go
│   │   └── connect_smbus.go
│   ├── hostname
│   │   └── hostname.go
│   ├── locale
│   │   └── locale.go
│   ├── login
│   │   └── login.go
│   ├── machine
│   │   └── machine.go
│   ├── systemd
│   │   ├── job
│   │   │   └── job.go
│   │   ├── object
│   │   │   ├── object.go
│   │   │   ├── object_get_props.go
│   │   │   └── object_props_getter.go
│   │   ├── process
│   │   │   └── process.go
│   │   ├── systemd.go
│   │   ├── systemd_job_controller.go
│   │   ├── systemd_jobs.go
│   │   ├── systemd_test.go
│   │   ├── systemd_units.go
│   │   ├── systemd_units_manage.go
│   │   ├── systemd_units_manager.go
│   │   └── unit
│   │       └── unit.go
│   └── timedate
│       └── timedate.go
└── treegen.sh

18 directories, 29 files
```
