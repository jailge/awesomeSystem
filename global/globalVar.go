package global

import (
	"awesomeSystem/config"
	"go.uber.org/zap"
)

var (
	Lg          *zap.Logger
	Settings    config.ServerConfig
	IdentityKey = "id"
)
