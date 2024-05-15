package repository

type TaskSQLite struct {
	db SQL
}

func NewTaskSQLite(db SQL) TaskSQLite {
	return TaskSQLite{db: db}
}
