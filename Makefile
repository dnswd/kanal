migrateup:
	migrate -path db/migration -database "postgresql://root:toor@127.0.0.1:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:toor@127.0.0.1:5432/postgres?sslmode=disable" -verbose down

migratedirty:
	migrate -path db/migration -database "postgresql://root:toor@127.0.0.1:5432/postgres?sslmode=disable" -verbose force 1

run:
	go run server.go

build-win:
	go build -o .\bin\server server.go; & .\bin\server

build-nix:
	go build -o ./bin server.go; ./bin/server

.PHONY: migrateup migratedown migratedirty run build-win build-nix