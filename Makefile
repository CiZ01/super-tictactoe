APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/internal/migrations
DATABASE_URL = postgres://ciz:0053@localhost:5050/postgres?sslmode=disable

security:
	gosec -quiet ./...

lint:
	golangci-lint run ./...

build:
	go build -o ./server/main ./server/main/

run:
	go build -o ./server/main.go ./server/main.go
	./server/main

tests:
	go test -v ./cmd/main/


# Migrations
migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

# Profiles
generate-profiles:
	go test ./cmd/main 		-cpuprofile cpu.prof 		-memprofile mem.prof 		-bench .

mem-profile: generate-profiles
	# Type: alloc_space
	go tool pprof mem.prof

cpu-profile: generate-profiles
	# Type: cpu
	go tool pprof cpu.prof

# Coverage

coverage:
	go test ./cmd/main -coverprofile=coverage.out -coverpkg=./...

coverage-html: coverage
	go tool cover -html=coverage.out

coverage-func: coverage
	go tool cover -func=coverage.out

# Tracing

traces:
	go test ./cmd/main -trace trace.out
	go tool trace pkg.test trace.out

