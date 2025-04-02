package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	ZapLog *zap.SugaredLogger
}

func NewLogger() *Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	return &Logger{ZapLog: sugar}
}

func (L *Logger) Debug(Msg ...interface{}) {
	L.ZapLog.Log(zap.DebugLevel, Msg)
}

func (L *Logger) Error(Msg ...interface{}) {
	L.ZapLog.Error(Msg)
}

func (L *Logger) Info(Msg ...interface{}) {
	L.ZapLog.Info(Msg)
}
