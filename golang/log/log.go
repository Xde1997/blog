package bloglog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Field = zap.Field

func String(key, value string) Field {
	return zap.String(key, value)
}

type Logger interface {
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	ErrErr(msg string, err error)
	Fatal(msg string, fields ...Field)
}

type zapLogger struct {
	log *zap.Logger
}

func (s *zapLogger) Info(msg string, fields ...Field) {
	s.log.Info(msg, fields...)
}

func (s *zapLogger) Debug(msg string, fields ...Field) {
	s.log.Debug(msg, fields...)
}

func (s *zapLogger) Error(msg string, fields ...Field) {
	s.log.Error(msg, fields...)
}

func (s *zapLogger) Fatal(msg string, fields ...Field) {
	s.log.Fatal(msg, fields...)
}

func (s *zapLogger) ErrErr(msg string, err error) {
	s.log.Error(msg, String("error", err.Error()))
}

type FileLoggerConfig struct {
	LogPath       string `json:"LogPath,omitempty"`
	ErrorFileName string `json:"ErrorFileName"`
	InfoFileName  string `json:"InfoFileName"`
	MaxSize       int    `json:"MaxSize"`
	MaxBackups    int    `json:"MaxBackups"`
	MaxAge        int    `json:"MaxAge"`
	Console       bool   `json:"Console"`
}

//NewFileLogger New filelogger
func NewFileLogger(cfg *FileLoggerConfig) (Logger, error) {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderJSON := zapcore.NewJSONEncoder(encoder)

	errPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	var allCore []zapcore.Core
	if cfg.ErrorFileName != "" {
		errWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.LogPath + "/" + cfg.ErrorFileName,
			MaxSize:    cfg.MaxSize, // megabytes
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge, // days
		})
		allCore = append(allCore, zapcore.NewCore(encoderJSON, errWriteSyncer, errPriority))
	}
	if cfg.InfoFileName != "" {
		infoWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.LogPath + "/" + cfg.InfoFileName,
			MaxSize:    cfg.MaxSize, // megabytes
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge, // days
		})
		allCore = append(allCore, zapcore.NewCore(encoderJSON, infoWriteSyncer, infoPriority))
	}
	return &zapLogger{
		log: zap.New(zapcore.NewTee(allCore...), zap.AddCaller()),
	}, nil
}

//NewConsoleLogger New ConsoleLogger
func NewConsoleLogger() (Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()

	return &zapLogger{logger}, err
}
