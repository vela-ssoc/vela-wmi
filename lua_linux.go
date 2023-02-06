package wmi

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

func WithEnv(env vela.Environment) {
	env.Warn("not support wmi with linux")
}
