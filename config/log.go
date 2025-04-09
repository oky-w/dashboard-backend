// Package config contains the application setup configuration settings.
package config

import (
	"os"
	"sync"

	"github.com/okyws/dashboard-backend/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logFile   *os.File
	once      sync.Once
	closeOnce sync.Once
)

// InitiateLog sets up the logger safely
func InitiateLog() {
	once.Do(func() {
		if _, err := os.Stat("log"); os.IsNotExist(err) {
			if err := os.Mkdir("log", os.ModeDir|0750); err != nil {
				log.Fatal().Err(err).Msg(constants.MsgDirrectoryError)
			}
		}

		// process creation log file
		var err error

		logFile, err = os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			log.Fatal().Err(err).Msg(constants.MsgFileOpenError)
		}

		log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
		log.Info().Msg(constants.MsgZerologInit)
	})
}

// CloseLog ensures the log file is closed safely
func CloseLog() {
	closeOnce.Do(func() {
		if logFile != nil {
			if err := logFile.Close(); err != nil {
				log.Error().Err(err).Msg(constants.MsgFileCloseError)
			}
		}
	})
}
