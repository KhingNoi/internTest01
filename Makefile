CREATE_COMMAND := migrate create -ext sql -dir ./migrations/ -seq

create-migrations:
	$(CREATE_COMMAND) $(filter-out $@,$(MAKECMDGOALS))

migrations-up:
	migrate -database "mysql://admin:password@tcp(localhost:3306)/intern_test" -path ./migrations up

migrations-down:
	migrate -database "mysql://admin:password@tcp(localhost:3306)/intern_test" -path ./migrations down

migrations-fix:
	migrate -database "mysql://admin:password@tcp(localhost:3306)/intern_test -path ./migrations force 1

