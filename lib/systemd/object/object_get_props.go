package object

import d "github.com/godbus/dbus"

func (o Object) GetAll(iface string) (map[string]d.Variant, error) {
    call := o.Call("org.freedesktop.DBus.Properties.GetAll", 0, iface)
    var props map[string]d.Variant
    err := call.Store(&props)
    if err != nil {
        return nil, err
    }
    return props, nil
}
