package object
import d "github.com/godbus/dbus"

type ObjectPropertiesGetter interface {
    GetAll(string) map[string]d.Variant
}
