package trace

import "github.com/rogercoll/flogs/internal/feature"

// Connect always traces connect feature.
func Connect(format string, args ...interface{}) {
	On(func() bool { return true }, feature.CONN, format, args...)
}

// Inventory traces to "audit" inventory.
func Inventory(format string, args ...interface{}) {
	On(func() bool { return true }, feature.INVENTORY, format, args...)
}

// Proc traces to "audit" process sampling.
func Proc(format string, args ...interface{}) {
	On(func() bool { return true }, feature.PROC, format, args...)
}
