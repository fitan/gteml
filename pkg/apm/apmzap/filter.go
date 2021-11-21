package apmzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func filter(m map[string]struct{}, fs []zapcore.Field) []zap.Field {
	res := make([]zap.Field, 0, len(m))
	for _, f := range fs {
		if _, ok := m[f.Key]; ok {
			res = append(res, f)
		}
	}
	return res
}
