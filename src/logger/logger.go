package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

var Logger *zap.SugaredLogger

func InitLogger() {

	logLevel := getZapLevel(viper.GetString("Log.Level"))
	logColor := viper.GetBool("Log.Color")
	logJson := viper.GetBool("Log.Json")

	var logEncoder zapcore.Encoder
	logEncoderConfig := zap.NewProductionEncoderConfig()
	logEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if logColor {
		logEncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	if logJson {
		logEncoder = zapcore.NewJSONEncoder(logEncoderConfig)
	} else {
		logEncoder = zapcore.NewConsoleEncoder(logEncoderConfig)
	}

	core := zapcore.NewCore(
		logEncoder,
		os.Stdout,
		zap.NewAtomicLevelAt(logLevel),
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()

	Logger.Infof("Init logger complete")
}

func SyncLogger() {
	Logger.Infof("Flush logger")
	Logger.Sync()
}
