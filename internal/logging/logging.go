package logging

import (
	"github.com/rs/zerolog"
	"os"
)

var loggerCache = make(map[string]zerolog.Logger)

func GetLoggerFor(loggerName string, isDevelopment bool) zerolog.Logger {
	if logger, ok := loggerCache[loggerName]; ok {
		return logger
	} else {
		var consoleWriter zerolog.ConsoleWriter
		var multi zerolog.LevelWriter
		if isDevelopment {
			consoleWriter = zerolog.ConsoleWriter{Out: os.Stdout}
			multi = zerolog.MultiLevelWriter(consoleWriter) //os.Stdout
		} else {
			multi = zerolog.MultiLevelWriter(os.Stdout) //os.Stdout
		}

		logger := zerolog.New(multi).
			With().
			Timestamp().
			Caller().
			Logger()

		loggerCache[loggerName] = logger

		return logger
	}
}
