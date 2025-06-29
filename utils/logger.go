package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Logger *zap.Logger

func InitLogger() {
	// Create logs folder if it doesn't exist
	logsDir := "logs"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	// Format current date
	currentDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("log_%s.log", currentDate)
	logFilePath := filepath.Join(logsDir, logFileName)

	// Open log file for appending, create if not exists
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// Logger config
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(file),
		zap.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller())
}
