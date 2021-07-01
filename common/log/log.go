package log

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"time"
)

func InitLog()  {
	initDir()
	initLog()
}

func initDir()  {
	ferr := os.MkdirAll("/var/cloud/", 0666)
	if ferr != nil {
		fmt.Errorf("Failed to create log path. Err: %v", ferr.Error())
	}
}

func initLog()  {
	//初始化hook日志配置
	hook := lumberjack.Logger{
		Filename  : LOGFILE,
		MaxSize   : MAXSIZE,
		MaxBackups: MAXBACKUPS,
		MaxAge    : MAXAGE,
		Compress  : true,
		LocalTime : true,
	}

	var syncer zapcore.WriteSyncer
	syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))

	rawJson, err := ioutil.ReadFile(LOGCFG)
	if err != nil {
		panic(err)
	}

	var cfg zap.Config

	if err := json.Unmarshal(rawJson, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = formatEncodingTime
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	CouldLog, err = cfg.Build2(syncer)
	if err != nil {
		panic(err)
	}
}

func formatEncodingTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}