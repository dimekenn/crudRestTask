migrateup:
    migrate -path db/migration -database "postgresql://dim:123@localhost:5432/crudDb?sslmode=disable" -verbose up