package logging

import (
	"log"
	"sync"

	config "github/fadlinux/edot/common/util/config"

	"github.com/sirupsen/logrus"

	"go.uber.org/zap"
)

var (
	once         sync.Once
	logrusClient *logrus.Entry
	zapClient    *zap.Logger
	useZapLogger bool
)

// Initialize : init log client
func Initialize() {
	useZapLogger = config.GetBool("feature.enable_zap_logger")

	if useZapLogger {
		once.Do(func() {
			zapClient, _ = zap.NewProduction()
		})
	} else {
		once.Do(func() {
			logrusClient = logrus.NewEntry(logrus.New())
			customFormatter := new(logrus.TextFormatter)
			customFormatter.TimestampFormat = "2006-01-02 15:04:05"
			logrusClient.Logger.SetFormatter(customFormatter)
		})
	}

	log.Println("[logging]init success")
}
