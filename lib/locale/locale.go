package locale

import "github.com/godbus/dbus"

const (
	FREEDESKTOP_LOCALE1 = "org.freedesktop.locale1"
	localeObjectPath    = "/org/freedesktop/locale1"
)

type Locale struct {
	bus dbus.BusObject
}

type LocaleProperties struct {
	Locale               []string
	X11Layout            string
	X11Model             string
	X11Variant           string
	X11Options           string
	VConsoleKeymap       string
	VConsoleKeymapToggle string
}

func GetLocale(conn *dbus.Conn) *Locale {
	loc := new(Locale)
	loc.bus = conn.Object(FREEDESKTOP_LOCALE1, localeObjectPath)
	return loc
}

func (l *Locale) GetProperties() (LocaleProperties, error) {
	props, err := l.getAll()
	if err != nil {
		return LocaleProperties{}, err
	}
	return LocaleProperties{
		Locale:               getStrings(props, "Locale"),
		X11Layout:            getString(props, "X11Layout"),
		X11Model:             getString(props, "X11Model"),
		X11Variant:           getString(props, "X11Variant"),
		X11Options:           getString(props, "X11Options"),
		VConsoleKeymap:       getString(props, "VConsoleKeymap"),
		VConsoleKeymapToggle: getString(props, "VConsoleKeymapToggle"),
	}, nil
}

func (l *Locale) SetLocale(locale []string, interactive bool) error {
	call := l.bus.Call("org.freedesktop.locale1.SetLocale", 0, locale, interactive)
	return call.Err
}

func (l *Locale) SetVConsoleKeyboard(keymap string, toggle string, convert bool, interactive bool) error {
	call := l.bus.Call("org.freedesktop.locale1.SetVConsoleKeyboard", 0, keymap, toggle, convert, interactive)
	return call.Err
}

func (l *Locale) SetX11Keyboard(layout string, model string, variant string, options string, convert bool, interactive bool) error {
	call := l.bus.Call("org.freedesktop.locale1.SetX11Keyboard", 0, layout, model, variant, options, convert, interactive)
	return call.Err
}

func (l *Locale) getAll() (map[string]dbus.Variant, error) {
	call := l.bus.Call("org.freedesktop.DBus.Properties.GetAll", 0, FREEDESKTOP_LOCALE1)
	if call.Err != nil {
		return nil, call.Err
	}
	var props map[string]dbus.Variant
	if err := call.Store(&props); err != nil {
		return nil, err
	}
	return props, nil
}

func getString(props map[string]dbus.Variant, key string) string {
	if v, ok := props[key]; ok {
		if value, ok := v.Value().(string); ok {
			return value
		}
	}
	return ""
}

func getStrings(props map[string]dbus.Variant, key string) []string {
	if v, ok := props[key]; ok {
		if value, ok := v.Value().([]string); ok {
			return value
		}
	}
	return nil
}
