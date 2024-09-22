.PHONY: run
run:
	air --build.cmd 'go build -o bin/api ./cmd/app' --build.bin './bin/api'	

.PHONY: api 
api:
	go build -o bin/api ./cmd/app

.PHONY: doc
doc:
	go run github.com/swaggo/swag/cmd/swag@latest init -g ./internal/app/api/app.go --pd --parseDepth 1 -o ./api --ot yml
