package login

import (
	"github.com/godbus/dbus"
)

const (
	FREEDESKTOP_LOGIN1 = "org.freedesktop.login1"
	loginObjectPath    = "/org/freedesktop/login1"
)

type Login struct {
	bus dbus.BusObject
}

type Session struct {
	SessionID  string
	UserID     uint32
	UserName   string
	Seat       string
	ObjectPath dbus.ObjectPath
}

type User struct {
	UserID     uint32
	UserName   string
	ObjectPath dbus.ObjectPath
}

type Seat struct {
	SeatID     string
	ObjectPath dbus.ObjectPath
}

type Inhibitor struct {
	What string
	Who  string
	Why  string
	Mode string
	PID  uint32
	UID  uint32
}

func GetLogin(conn *dbus.Conn) *Login {
	l := new(Login)
	l.bus = conn.Object(FREEDESKTOP_LOGIN1, loginObjectPath)
	return l
}

func (l *Login) ListSessions() ([]Session, error) {
	var rows [][]any
	call := l.bus.Call("org.freedesktop.login1.Manager.ListSessions", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	sessions := make([]Session, 0, len(rows))
	for _, row := range rows {
		if len(row) != 5 {
			continue
		}
		sessionID, _ := row[0].(string)
		userID := toUint32(row[1])
		userName, _ := row[2].(string)
		seat, _ := row[3].(string)
		path, _ := row[4].(dbus.ObjectPath)
		sessions = append(sessions, Session{
			SessionID:  sessionID,
			UserID:     userID,
			UserName:   userName,
			Seat:       seat,
			ObjectPath: path,
		})
	}
	return sessions, nil
}

func (l *Login) ListUsers() ([]User, error) {
	var rows [][]any
	call := l.bus.Call("org.freedesktop.login1.Manager.ListUsers", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	users := make([]User, 0, len(rows))
	for _, row := range rows {
		if len(row) != 3 {
			continue
		}
		uid := toUint32(row[0])
		name, _ := row[1].(string)
		path, _ := row[2].(dbus.ObjectPath)
		users = append(users, User{
			UserID:     uid,
			UserName:   name,
			ObjectPath: path,
		})
	}
	return users, nil
}

func (l *Login) ListSeats() ([]Seat, error) {
	var rows [][]any
	call := l.bus.Call("org.freedesktop.login1.Manager.ListSeats", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	seats := make([]Seat, 0, len(rows))
	for _, row := range rows {
		if len(row) != 2 {
			continue
		}
		id, _ := row[0].(string)
		path, _ := row[1].(dbus.ObjectPath)
		seats = append(seats, Seat{
			SeatID:     id,
			ObjectPath: path,
		})
	}
	return seats, nil
}

func (l *Login) GetSession(id string) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := l.bus.Call("org.freedesktop.login1.Manager.GetSession", 0, id)
	err := call.Store(&path)
	return path, err
}

func (l *Login) GetSessionByPID(pid uint32) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := l.bus.Call("org.freedesktop.login1.Manager.GetSessionByPID", 0, pid)
	err := call.Store(&path)
	return path, err
}

func (l *Login) GetUser(uid uint32) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := l.bus.Call("org.freedesktop.login1.Manager.GetUser", 0, uid)
	err := call.Store(&path)
	return path, err
}

func (l *Login) GetUserByPID(pid uint32) (dbus.ObjectPath, error) {
	call := l.bus.Call("org.freedesktop.login1.Manager.GetUserByPID", 0, pid)
	var path dbus.ObjectPath
	err := call.Store(&path)
	return path, err
}

func (l *Login) GetSeat(id string) (dbus.ObjectPath, error) {
	var path dbus.ObjectPath
	call := l.bus.Call("org.freedesktop.login1.Manager.GetSeat", 0, id)
	err := call.Store(&path)
	return path, err
}

func (l *Login) ListInhibitors() ([]Inhibitor, error) {
	var rows [][]any
	call := l.bus.Call("org.freedesktop.login1.Manager.ListInhibitors", 0)
	if err := call.Store(&rows); err != nil {
		return nil, err
	}
	inhibitors := make([]Inhibitor, 0, len(rows))
	for _, row := range rows {
		if len(row) != 6 {
			continue
		}
		what, _ := row[0].(string)
		who, _ := row[1].(string)
		why, _ := row[2].(string)
		mode, _ := row[3].(string)
		pid := toUint32(row[4])
		uid := toUint32(row[5])
		inhibitors = append(inhibitors, Inhibitor{
			What: what,
			Who:  who,
			Why:  why,
			Mode: mode,
			PID:  pid,
			UID:  uid,
		})
	}
	return inhibitors, nil
}

func toUint32(value any) uint32 {
	switch v := value.(type) {
	case uint32:
		return v
	case uint16:
		return uint32(v)
	case uint64:
		return uint32(v)
	case int32:
		return uint32(v)
	case int:
		return uint32(v)
	default:
		return 0
	}
}
