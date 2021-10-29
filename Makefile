run:
	gin.exe -i run main.go

genconf:
	go run tools/gen/main.go genconf -s ./conf.yaml -d ./pkg/types/conf.go

ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema