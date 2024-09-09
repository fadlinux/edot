package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogInfo : log for info
func LogInfo(message string, input map[string]string) (err error) {
	defer func() {
		if useZapLogger {
			zapClient.Sync()
		}
	}()

	if useZapLogger {
		zapFields := []zapcore.Field{}
		for name, val := range input {
			zapFields = append(zapFields, zap.String(name, val))
		}

		zapClient.Info(message, zapFields...)
	} else {
		logrusClient.Println(message, input)
	}

	return
}

// LogError : log for error
func LogError(message string, errMsg error) (err error) {
	defer func() {
		if useZapLogger {
			zapClient.Sync()
		}
	}()

	if useZapLogger {
		zapClient.Error(message, zap.Error(errMsg))
	} else {
		logrusClient.Errorln(message, errMsg)
	}

	return
}
