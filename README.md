# sagala-tech-test

Hallo Sagala Team, terima kasih telah memberi kesempatan untuk Technical Test.
Berikut untuk tahapan Test yang saya kerjakan.

### Announce
Pengerjaan Task ini saya kerjakan dalam Bahasa Go (Golang) dan menggunakan PostgreSQL. Untuk Tech atau library atau package yang digunakan akan diinfokan saat mengikuti tahap pengecekan

### STEP 1
Pastikan Postgres sudah terinstall di local komputer.
buat database dengan nama database 
> `sagala_tech_test`

### STEP 2
install Package berikut ini yang akan digunakan untuk migrasi database :
> `go get -u -d github.com/golang-migrate/migrate/cmd/migrate`
> `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

**Penggunaan Package**
Create Migration Format: migrate create -ext sql -dir [directory_file_migrations] [migration_name]
example : migrate create -ext sql -dir database/migrations create_table_task

**Running Migration Up:**
Running Migrate Format: migrate -database "postgres://[user]:[password]@[host]:[port]/[dbName]?sslmode=disable" -path [directory_file_migrations] up
example : migrate -database "postgres://postgres:password@localhost:5432/sagala_tech_test?sslmode=disable" -path database/migration up

Untuk tahap ini silakan running ini di terminal :
> `migrate -database "postgres://postgres:password@localhost:5432/sagala_tech_test?sslmode=disable" -path database/migration up`

### STEP 3
You're Ready to GO!
Done and you can Running Go as usual.
> `go run main.go`

## HOW to USE
Setelah melakukan proses sebelumnya pada bagian ini akan fokus pada penjelasan API yang tersedia.
Karena app ini running di local maka ini adalah List API yang sudah dibuat :
- Create Task : [POST] `localhost:5000/task`
- Get List Tasks : [POST] `localhost:5000/tasks`
- Get Task Detail by Id : [GET] `localhost:5000/task/{{task_id}}`
- Update Task : [PUT] `localhost:5000/task/{{task_id}}`
- Update Status Task : [PUT] `localhost:5000/task/{{task_id}}/{{status_id}}`
- (Soft) Delete Task : [PUT] `localhost:5000/task/remove/{{task_id}}`
- Delete Task : [DELETE] `localhost:5000/task/{{task_id}}`

### Create Task
EndPoint : [POST] `localhost:5000/task`
Payload :
```
{
	"task_name": "Test",
	"description": "Test ini meruoakan contoh",
	"task_duration_minutes": 60,
	"due_date": "2024-07-20T10:10:30Z"
}
```

### Get List Tasks
Endpoint : [POST] `localhost:5000/tasks`
Payload :
```
{
    "status" : "in_progress",
    "is_deleted": true
}
```

### Get Task Detail by Id
Endpoint : [GET] `localhost:5000/task/{{task_id}}`

### Update Task
Endpoint : [PUT] `localhost:5000/task/{{task_id}}`
Payload :
```
{
	"task_name": "Test update",
	"description": "Update Test ini meruoakan contoh",
	"status": "waiting_list",
	"is_deleted": false,
	"task_duration_minutes": 60,
	"due_date": "2024-07-20T10:10:30Z"
}
```

### Update Status Task
Endpoint : [PUT] `localhost:5000/task/{{task_id}}/{{status_id}}`

### (Soft) Delete Task
Endpoint : [PUT] `localhost:5000/task/remove/{{task_id}}`

### Delete Task
Endpoint : [DELETE] `localhost:5000/task/{{task_id}}`


## GO TEST
for Running go test :
> `go test ./...`