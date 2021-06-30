package log

import "go.uber.org/zap"

const (
	LOGPATH   = "/var/log/cloud/"
	LOGFILE   = LOGPATH + "cloud.log"
	LOGCFG 	  = LOGPATH + "log.config"
	MAXSIZE	  = 50
	MAXBACKUPS= 31
	MAXAGE	  = 90
)

var CouldLog *zap.Logger

