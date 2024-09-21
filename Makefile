.PHONY: run
run:
	./bin/api

.PHONY: api 
api:
	go build -o bin/api ./cmd/app

.PHONY: admin 
admin:
	go build -o bin/admin ./cmd/admin

.PHONY: doc
doc:
	go run github.com/swaggo/swag/cmd/swag@latest init -g ./internal/app/api/app.go --pd --parseDepth 1 -o ./api --ot yml
