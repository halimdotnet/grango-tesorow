package logger

import (
	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildEncoder(env string) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if env == constants.EnvProduction {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeName = zapcore.FullNameEncoder

		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoderConfig.ConsoleSeparator = " | "
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeName = zapcore.FullNameEncoder

		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return encoder
}

func mapLogLevel(lvl LogLevel) zapcore.Level {
	switch lvl {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
