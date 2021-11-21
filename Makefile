export ELASTIC_APM_SERVICE_NAME=gteml
export ELASTIC_APM_SERVER_URL=http://localhost:8200

build:
	go build -ldflags "-X main.GitCommitId=`git rev-parse HEAD` -X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"  -mod vendor -o output/main main.go

run:
	go run main.go

gen-conf:
	go run tools/main.go genconf -s ./conf.yaml -d ./pkg/types/conf_gen.go

ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema

migrate:
	go run tools/gen/main.go  migrate

watch:
	gin.exe -i run main.go
