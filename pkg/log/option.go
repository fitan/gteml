package log

import (
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type option struct {
	traceLevel zapcore.Level
	tp         *otelsdk.TracerProvider
	logLevel   zapcore.Level
	fileName   string
	*zap.Logger
}

type Option func(o *option)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func WithLogger(logger *zap.Logger) Option {
	return func(o *option) {
		o.Logger = logger
	}
}

func WithLogFileName(fileName string) Option {
	return func(o *option) {
		o.fileName = fileName
	}
}

func WithLogLevel(level zapcore.Level) Option {
	return func(o *option) {
		o.logLevel = level
	}
}

func WithTrace(tp *otelsdk.TracerProvider, level zapcore.Level) Option {
	return func(o *option) {
		o.traceLevel = level
		o.tp = tp
	}
}

func NewXlog(fs ...Option) *Xlog {
	xlog := new(Xlog)
	o := new(option)
	for _, f := range fs {
		f(o)
	}
	if o.Logger == nil {
		fileName := o.fileName
		logLever := o.logLevel
		if fileName == "" {
			fileName = "./logs/xlog.log"
		}
		writer := getLogWriter(fileName)
		core := zapcore.NewCore(getEncoder(), writer, logLever)
		log := zap.New(core, zap.AddCaller())
		xlog.Logger = log
	} else {
		xlog.Logger = o.Logger
	}

	xlog.traceLevel = o.traceLevel
	xlog.tp = o.tp

	return xlog
}
