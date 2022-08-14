package logging

import (
	"github.com/fatih/structs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ApplicationLog struct {
	RequestId string `json:"requestId" structs:"requestId"`
	Message   string `json:"message" structs:"message"`
}

func Print(level zerolog.Level, requestId string, msg string) {

	applog := ApplicationLog{
		RequestId: requestId,
		Message:   msg,
	}

	m := structs.Map(&applog)

	switch level {
	case zerolog.InfoLevel:
		log.Info().Fields(m).Msg("")
	}
}
