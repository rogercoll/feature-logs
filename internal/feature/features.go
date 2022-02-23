package feature

// Feature feature to be traced.
type Feature string

// String stringer method
func (f Feature) String() string {
	return string(f)
}

const (
	INVENTORY Feature = "inventory"
	PROC      Feature = "proc"
	CONN      Feature = "connect"
)
