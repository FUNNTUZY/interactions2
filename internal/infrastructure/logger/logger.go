package logger

import (
	"os"

	"interactions/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func InitLogger(cfg config.LoggerConfig) {
	// Устанавливаем время как UNIX timestamp для логов
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Если уровень логирования указан, устанавливаем его
	switch cfg.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Выбор формата вывода логов
	if cfg.LogJsonFormat {
		// Логи в формате JSON
		log.Logger = log.Output(os.Stderr)
	} else {
		// Логи в текстовом формате, удобном для чтения человеком
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Добавление информации об ошибках
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Info().Msg("Logger init sucess")
}
