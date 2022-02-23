package connection

import (
	"github.com/rogercoll/flogs/internal/level/debug"
	"github.com/rogercoll/flogs/internal/level/info"
	"github.com/rogercoll/flogs/internal/level/trace"
)

func Foo() {
	trace.Connect("connection trace")
	info.Connect("connection info")
	debug.Connect("connection debug")
}
