package logging

import (
	"github.com/fatih/structs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ApplicationLog struct {
	RequestId string `json:"requestId" structs:"requestId"`
	Service   string `json:"service" structs:"service"`
	Message   string `json:"message" structs:"message"`
}

func Print(level zerolog.Level, requestId string, service string, msg string) {

	applog := ApplicationLog{
		RequestId: requestId,
		Service:   service,
		Message:   msg,
	}

	m := structs.Map(&applog)

	switch level {
	case zerolog.InfoLevel:
		log.Info().Fields(m).Msg("")
	}
}
