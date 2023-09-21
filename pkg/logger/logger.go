package logger

import (
	"log"
	"os"

	"gorm.io/gorm/logger"
)

func NewLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			LogLevel:                  0,
		})
}
