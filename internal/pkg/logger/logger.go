package logger

import (
	"os"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

type Config struct {
	Level        LogLevel `mapstructure:"level"`
	EnableCaller bool     `mapstructure:"enable_caller"`
	EnableTrace  bool     `mapstructure:"enable_trace"`
}

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
	logger        *zap.Logger
}

func New(cfg *Config, env string) Logger {
	if env == "" {
		env = constants.EnvDevelopment
	}

	encoder := buildEncoder(env)
	level := mapLogLevel(cfg.Level)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	var opts []zap.Option

	if cfg.EnableCaller {
		opts = append(opts, zap.AddCaller())
		opts = append(opts, zap.AddCallerSkip(1))
	}

	if cfg.EnableTrace {
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	logx := zap.New(core, opts...)

	return &zapLogger{
		sugaredLogger: logx.Sugar(),
		logger:        logx,
	}
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.sugaredLogger.Debug(args...)
}

func (z *zapLogger) Info(args ...interface{}) {
	z.sugaredLogger.Info(args...)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.sugaredLogger.Warn(args...)
}

func (z *zapLogger) Error(args ...interface{}) {
	z.sugaredLogger.Error(args...)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.sugaredLogger.Fatal(args...)
}

func (z *zapLogger) Debugf(str string, args ...interface{}) {
	z.sugaredLogger.Debugf(str, args...)
}

func (z *zapLogger) Infof(str string, args ...interface{}) {
	z.sugaredLogger.Infof(str, args...)
}

func (z *zapLogger) Warnf(str string, args ...interface{}) {
	z.sugaredLogger.Warnf(str, args...)
}

func (z *zapLogger) Errorf(str string, args ...interface{}) {
	z.sugaredLogger.Errorf(str, args...)
}

func (z *zapLogger) Fatalf(str string, args ...interface{}) {
	z.sugaredLogger.Fatalf(str, args...)
}

func (z *zapLogger) Sync() error {
	return z.sugaredLogger.Sync()
}
