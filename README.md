# sagala-tech-test

Pastikan Postgres sudah terinstall di local komputer.
buat database dengan nama database : sagala_tech_test

go get -u -d github.com/golang-migrate/migrate/cmd/migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Create Migration:
migrate create -ext sql -dir [directory_file_migrations] [migration_name]
example : migrate create -ext sql -dir database/migrations create_table_employees

Running Migration Up:
migrate -database "postgres://[user]:[password]@[host]:[port]/[dbName]?sslmode=disable" -path [directory_file_migrations] up
migrate -database "postgres://postgres:password@localhost:5432/sagala_tech_tesy?sslmode=disable" -path database/migration up

Create Payload :
{
    "employee_name" : "Jonh Wick",
    "job_title" : "Manager",
    "salary" : 53000,
    "department" : "Sales",
    "joined_date" : "2022-10-08T10:30:00Z" 
}