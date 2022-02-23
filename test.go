package flogs

import (
	"github.com/rogercoll/flogs/internal/level/info"
	"github.com/rogercoll/flogs/internal/level/trace"
	"github.com/rogercoll/flogs/pkg/connection"
)

func Test() {
	trace.EnableOn([]string{"connect"})
	info.EnableOn([]string{"connect"})
	connection.Foo()
}
