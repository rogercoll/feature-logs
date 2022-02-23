package info

import (
	"fmt"
	"os"

	"github.com/rogercoll/flogs/internal/feature"
	"github.com/rs/zerolog"
)

type tracer struct {
	enabled map[feature.Feature]struct{}
	logger  *zerolog.Logger
}

// Condition lazy-load approach to evaluate condition only
type Condition func() bool

// global tracer instance, aimed to be used the same way log package is.
var global tracer

// EnableOn enables tracing capability for a set of features
// It uses the same logger instance as the agent.
func EnableOn(features []string) {
	if len(features) == 0 {
		return
	}
	zl := zerolog.New(os.Stdout)
	global = tracer{
		enabled: make(map[feature.Feature]struct{}),
		logger:  &zl,
	}

	for _, f := range features {
		global.enabled[feature.Feature(f)] = struct{}{}
	}
}

// On logs data to be shown if trace for a given feature is enabled and condition is met.
func On(condition Condition, f feature.Feature, format string, args ...interface{}) {
	if global.logger == nil {
		return
	}

	if _, ok := global.enabled[f]; ok && condition() {
		global.logger.Info().Str("feature", string(f)).Msg(fmt.Sprintf(format, args...))
	}
}

// IsEnabled returns if feature is enabled.
func IsEnabled(feature feature.Feature) bool {
	_, ok := global.enabled[feature]

	return ok
}
