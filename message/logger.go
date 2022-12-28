package message

import (
	"context"
	log "github.com/sirupsen/logrus"
	"ojire/utils"
	"time"
)

func LogContext(logCtx string, ctx context.Context) *log.Entry {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc).Format("2006-01-02 15:04:05")

	entry := log.WithFields(log.Fields{
		"topic":   "go-dropping-list",
		"context": logCtx,
		"at":      timeNow,
	})

	if ctx != nil {
		entry = entry.WithFields(
			log.Fields{
				"correlation-id": utils.GetCorrelationIDFromContext(ctx),
			})
	}

	return entry
}

func Log(ctx context.Context, level log.Level, message, logCtx string) {
	entry := LogContext(logCtx, ctx)

	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	}
}

func GetLogger() *log.Logger {
	return log.StandardLogger()
}
