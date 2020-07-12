package services

import (
	"wordgame/tdog/core"
)

type (
	Logger struct {
		Base   core.Service
		IpAddr string
	}
)
