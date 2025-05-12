package logger

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var loggerRegistry = make(map[string]*logrus.Logger)
var loggerMutex sync.Mutex

func GetLogger(module string) *logrus.Logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	if logger, exists := loggerRegistry[module]; exists {
		return logger
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	loggerRegistry[module] = logger
	return logger
}
