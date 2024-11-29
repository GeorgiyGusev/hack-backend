migration-up:
	migrate -database postgres://minin:minin@kowlad.ru:5432/postgres?sslmode=disable -path migrations up

migration-down:
	migrate -database postgres://minin:minin@localhost:5432/postgres?sslmode=disable -path migrations down