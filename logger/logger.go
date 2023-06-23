package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	onceInit sync.Once
)

func Init(lvl int) error {
	var err error

	onceInit.Do(func() {
		globalLevel := zapcore.Level(lvl)

		highPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return lvl >= int(zapcore.ErrorLevel)
		})

		lowPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return lvl >= int(globalLevel) && lvl < int(zapcore.ErrorLevel)
		})

		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02T15:04:05.999Z"))
		}
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
		)

		log = zap.New(core)
		zap.RedirectStdLog(log)
	})

	return err
}

func Debugf(format string, args ...interface{}) {
	log.Sugar().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Sugar().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Sugar().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Sugar().Errorf(format, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(format string, args ...interface{}) {
	log.Sugar().Panicf(format, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(format string, args ...interface{}) {
	log.Sugar().Fatalf(format, args...)
}
