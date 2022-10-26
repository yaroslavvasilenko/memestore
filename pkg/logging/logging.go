package logging

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"memestore/pkg/config"
)

func InitLog(cfg *config.Config) (*os.File, error) {
	logFile, err := initLogFile(cfg.LogPath, 10000)
	if err != nil {
		return nil, err
	}
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	log.SetLevel(level)

	return logFile, err
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
