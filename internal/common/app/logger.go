package app

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/indigowar/not_amazing_amazon/internal/common/config"
)

func createLogger(cfg *config.Config) *slog.Logger {
	var logWriter io.Writer = os.Stdout
	if cfg.LogFile != "" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		logWriter = io.MultiWriter(os.Stdout, file)
	}

	return slog.New(slog.NewJSONHandler(logWriter, nil))
}
