package config

import (
	log "github/fadlinux/edot/common/util/log"

	"github.com/sirupsen/logrus"
)

// UnmarshalKey takes a single key and unmarshals it into a Struct.
func UnmarshalKey(key string, rawVal interface{}, configName ...string) (err error) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	err = vInstances[name].UnmarshalKey(key, rawVal)
	return
}
