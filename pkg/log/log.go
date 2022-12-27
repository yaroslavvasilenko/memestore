package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var sugar *zap.SugaredLogger

func InitLog() string {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")
	config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info
	pathLog := getVarEnv("LOG_PATH", "./log.txt")
	_, err := initLogFile(pathLog, 10000)
	if err != nil {
		log.Panic(fmt.Errorf("cannot create init log file %s", err.Error()))
	}
	config.OutputPaths = []string{pathLog, "stderr"}
	l, err := config.Build()
	if err != nil {
		log.Panic(fmt.Errorf("cannot create zap logger %s", err.Error()))
	}
	sugar = l.Sugar()

	return pathLog
}

func initLogFile(path string, fileMaxSizeBytes int64) (*os.File, error) {
	fInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		} else {
			return nil, err
		}
	}
	if fInfo.Size() > fileMaxSizeBytes {
		if err = os.Remove(path); err != nil {
			return nil, err
		}
	}
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
}

func getVarEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVal
}

func Log() *zap.SugaredLogger {
	return sugar
}
