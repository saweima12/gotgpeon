package logger

import (
	"go.uber.org/zap"
)

var instance *zap.SugaredLogger

// Initialize logger instance.
func InitLogger(mode string) error {
	var logger *zap.Logger
	var err error

	if mode == "dev" {
		logger, err = zap.NewDevelopment(
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		)
	} else {
		logger, err = zap.NewProduction(
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		)
	}

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

var (
	Error  = instance.Error
	Errorf = instance.Errorf
	Info   = instance.Info
	Infof  = instance.Infof
	Warn   = instance.Warn
	Warnf  = instance.Warnf
	Debug  = instance.Debug
	Debugf = instance.Debugf
)
