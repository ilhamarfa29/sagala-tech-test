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

API ini digunakan untuk membuat Tugas baru. Silakan isi data tugas sesuai payload dibawah ini. Berikut adalah penjelasan data Payload yang disubmit :
- `task_name` : Nama Tugas yang akan dibuat
- `description` : Deskripsi dari tugas yang akan dibuat
- `task_duration_minutes` : Perkiraan Durasi dari tugas yang akan dibuat
- `due_date` : Deadline waktu dari Tugas

Payload :
```
{
	"task_name": "Nama Tugas",
	"description": "Deskripsi Tugas",
	"task_duration_minutes": 60,
	"due_date": "2024-07-25T10:10:30Z"
}
```

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "task": {
        "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
        "task_name": "Nama Tugas",
        "description": "Deskripsi Tugas",
        "status": "waiting_list",
        "is_deleted": false,
        "task_duration_minutes": 60,
        "due_date": "2024-07-25T10:10:30Z",
        "created_at": "2024-07-21T09:24:41.4058491+07:00"
    }
}
```

### Get List Tasks
Endpoint : [POST] `localhost:5000/tasks`

API ini digunakan untuk mendapatkan list Tugas yang dibutuhkan. Berikut adalah penjelasan data Payload yang disubmit :
- `status` : Filter Status apa yang ingin ditampilkan di List
- `is_deleted` : Filter Tugas yang belum dihapus (`false`) atau yang sudah dihapus (`true`)

Payload :
```
{
    "status" : "in_progress",
    "is_deleted": true
}
```

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "tasks": [
        {
            "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
            "task_name": "Nama Tugas",
            "description": "Deskripsi Tugas",
            "status": "waiting_list",
            "is_deleted": false,
            "task_duration_minutes": 60,
            "due_date": "2024-07-25T17:10:30+07:00",
            "created_at": "2024-07-21T09:24:41+07:00"
        },
	]
}
```

### Get Task Detail by Id
Endpoint : [GET] `localhost:5000/task/{{task_id}}`

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "task": {
        "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
        "task_name": "Nama Tugas",
        "description": "Deskripsi Tugas",
        "status": "waiting_list",
        "is_deleted": false,
        "task_duration_minutes": 60,
        "due_date": "2024-07-25T10:10:30Z",
        "created_at": "2024-07-21T09:24:41.4058491+07:00"
    }
}
```

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

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "task": {
        "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
        "task_name": "Nama Tugas",
        "description": "Deskripsi Tugas",
        "status": "waiting_list",
        "is_deleted": false,
        "task_duration_minutes": 60,
        "due_date": "2024-07-25T10:10:30Z",
        "created_at": "2024-07-21T09:24:41.4058491+07:00"
    }
}
```

### Update Status Task
Endpoint : [PUT] `localhost:5000/task/{{task_id}}/{{status_id}}`

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "task": {
        "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
        "task_name": "Nama Tugas",
        "description": "Deskripsi Tugas",
        "status": "waiting_list",
        "is_deleted": false,
        "task_duration_minutes": 60,
        "due_date": "2024-07-25T10:10:30Z",
        "created_at": "2024-07-21T09:24:41.4058491+07:00"
    }
}
```

### (Soft) Delete Task
Endpoint : [PUT] `localhost:5000/task/remove/{{task_id}}`

**Response**
Response dari API ini adalah:
- `task_id` : Task id digenerate oleh system, jadi akan muncul saat Tugas sudah terbuat
- `task_name` : Nama Tugas seperti yang diinputkan
- `description` : Deskripsi Tugas seperti yang diinputkan
- `status` : Status akan terbuat otomatis di status waiting list
- `is_deleted` : Secara default value dari properi ini pasti bernilai false. Ini menjadi flag kalau tugas sudah dihapus secara soft delete
- `task_duration_minutes` : Perkiraan Durasi Tugas seperti yang diinputkan
- `due_date` : Deadline Tugas sesuai dengan yang diinputkan
- `created_at` : Data waktu saat Tugas dibuat di system

Response :
```
{
    "task": {
        "task_id": "3a8dd53c-511a-4265-8b24-824bc03eb55a",
        "task_name": "Nama Tugas",
        "description": "Deskripsi Tugas",
        "status": "waiting_list",
        "is_deleted": true,
        "task_duration_minutes": 60,
        "due_date": "2024-07-25T10:10:30Z",
        "created_at": "2024-07-21T09:24:41.4058491+07:00"
    }
}
```

### Delete Task
Endpoint : [DELETE] `localhost:5000/task/{{task_id}}`

**Response**
Response dari API ini adalah:
- `message` : Pesan dari hasil proses delete data tugas

Response :
```
{
    "message": "task deleted successfully"
}
```

## GO TEST
for Running go test :
> `go test ./...`