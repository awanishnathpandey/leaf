goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" up

goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" down

