module github.com/fitan/magic

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/Depado/ginprom v1.7.2
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.4
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pyroscope-io/pyroscope v0.2.2
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.24.0
	go.opentelemetry.io/otel v1.0.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0
	go.opentelemetry.io/otel/sdk v1.0.0
	go.opentelemetry.io/otel/trace v1.0.0
	go.uber.org/zap v1.19.1
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.2.0 // indirect
	gorm.io/gen v0.1.20
	gorm.io/gorm v1.22.3 // indirect
)
