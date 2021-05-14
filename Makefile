build-api:
		go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/api ./cmd/app/main.go
run-api:
		./.bin/api
build-cron:
		go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/cron ./cmd/cron/main.go
run-cron:
		./.bin/cron
test-integration:
		go test -v ./tests/...