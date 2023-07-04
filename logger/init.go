package logger

import (
	"go.uber.org/zap"
)

var instance *zap.SugaredLogger

// Initialize logger instance.
func InitLogger() error {
	logger, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	if err != nil {
		return err
	}
	instance = logger.Sugar()
	return nil
}

// release logger memory
func Dispose() error {
	return instance.Sync()
}

// Get logger instance
func GetLogger() *zap.SugaredLogger {
	return instance
}

/// ---
/// Wrapper Function
/// ----
func Error(args ...interface{}) {
	instance.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	instance.Errorf(template, args...)
}

func Info(args ...interface{}) {
	instance.Info(args...)
}

func Infof(template string, args ...interface{}) {
	instance.Infof(template, args...)
}

func Warn(args ...interface{}) {
	instance.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	instance.Warnf(template, args...)
}

func Debug(args ...interface{}) {
	instance.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	instance.Debugf(template, args...)
}
