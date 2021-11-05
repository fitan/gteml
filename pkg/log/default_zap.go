package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
)

func DefaultZapCore(fileName string, dir string, openLevel zapcore.Level) zapcore.Core {
	errEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level >= zap.ErrorLevel && zap.ErrorLevel > openLevel
		})
	infoEnable := zap.LevelEnablerFunc(
		func(level zapcore.Level) bool {
			return level < zap.ErrorLevel && level >= zap.DebugLevel && level > openLevel
		})

	infoLogWriter := getLogWriter(path.Join(dir, fileName+"_info.log"))
	errLogWriter := getLogWriter(path.Join(dir, fileName+"_err.log"))

	infoCore := zapcore.NewCore(getEncoder(), infoLogWriter, infoEnable)
	errCore := zapcore.NewCore(getEncoder(), errLogWriter, errEnable)

	return zapcore.NewTee(infoCore, errCore)
}
