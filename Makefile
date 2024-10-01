.PHONY: run
run: doc
	@air --build.cmd 'go build -o bin/api ./cmd/app' --build.bin './bin/api'	

.PHONY: api 
api:
	@go build -o bin/api ./cmd/app

.PHONY: doc
doc:
	@go run github.com/swaggo/swag/cmd/swag@v1.16 fmt internal && \
	go run github.com/swaggo/swag/cmd/swag@v1.16 init -g internal/app/api/app.go -p pascalcase --pd --parseDepth 1 -o ./api --ot go,json
