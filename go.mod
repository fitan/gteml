module github.com/fitan/magic

go 1.16

require (
	cuelang.org/go v0.2.2
	github.com/Depado/ginprom v1.7.2
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/appleboy/gin-jwt/v2 v2.7.0
	github.com/casbin/casbin/v2 v2.39.1
	github.com/casbin/gorm-adapter/v3 v3.4.5
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.10.3 // indirect
	github.com/oam-dev/kubevela-core-api v1.1.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pyroscope-io/pyroscope v0.2.2
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.4
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.7
	go-micro.dev/v4 v4.5.0 // indirect
	go.elastic.co/apm v1.14.0
	go.elastic.co/apm/module/apmhttp v1.14.0
	go.elastic.co/apm/module/apmzap v1.14.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.28.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.24.0
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0
	go.opentelemetry.io/otel/sdk v1.0.0
	go.opentelemetry.io/otel/trace v1.3.0
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.2.3
	gorm.io/gen v0.2.35
	gorm.io/gorm v1.22.5
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/cli-runtime v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/kubectl v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
)
