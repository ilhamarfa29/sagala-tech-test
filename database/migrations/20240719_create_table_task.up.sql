CREATE TABLE task (
	task_id varchar(100) PRIMARY KEY,
	task_name varchar(255) NOT NULL,
	description text, 
	status varchar(50),
	is_deleted int,
	task_duration_minutes int,
	due_date timestamp(0) with time zone DEFAULT CURRENT_TIMESTAMP(6),
	created_at timestamp(0) with time zone DEFAULT CURRENT_TIMESTAMP(6),
);