package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type option struct {
	traceLevel zapcore.Level
	//tp         *otelsdk.TracerProvider
	openTrace bool
	logLevel  zapcore.Level
	dir       string
	fileName  string
	filter    map[string]struct{}
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

func WithLogFileName(fileName string, dir string) Option {
	return func(o *option) {
		o.fileName = fileName
		o.dir = dir
	}
}

func WithLogLevel(level zapcore.Level) Option {
	return func(o *option) {
		o.logLevel = level
	}
}

func WithTrace(level zapcore.Level, filter map[string]struct{}) Option {
	return func(o *option) {
		o.traceLevel = level
		//o.tp = tp
		o.openTrace = true

		o.filter = filter

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
		dir := o.dir
		if fileName == "" {
			fileName = "xlog"
		}
		if dir == "" {
			dir = "./logs"
		}

		core := DefaultZapCore(fileName, dir, logLever)
		//writer := getLogWriter(fileName)
		//core := zapcore.NewCore(getEncoder(), writer, logLever)
		log := zap.New(core, zap.AddCaller())
		xlog.Logger = log
	} else {
		xlog.Logger = o.Logger
	}

	if o.openTrace {
		xlog.traceLevel = o.traceLevel
		xlog.openTrace = o.openTrace
		xlog.filter = o.filter
	}

	return xlog
}
