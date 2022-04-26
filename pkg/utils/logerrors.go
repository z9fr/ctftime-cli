package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogError(err error) {
	log.Fatal(err)

}
