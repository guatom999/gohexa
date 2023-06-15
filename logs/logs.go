package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// log.Info()
}

func Info(message interface{}, fields ...zapcore.Field) {
	msg, ok := message.(string)
	if ok {
		log.Info(msg, fields...)
	}

}

func Debug(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Error(message interface{}, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
