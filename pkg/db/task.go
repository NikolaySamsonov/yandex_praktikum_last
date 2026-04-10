package db

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	query := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (?, ?, ?, ?)
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func Tasks(limit int) ([]*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	ORDER BY date
	LIMIT ?
	`

	rows, err := DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		var t Task
		var id int64

		err = rows.Scan(&id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}

		t.ID = strconv.FormatInt(id, 10)
		tasks = append(tasks, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = ?
	`

	var task Task
	var taskID int64

	err := DB.QueryRow(query, id).Scan(
		&taskID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found")
		}
		return nil, err
	}

	task.ID = strconv.FormatInt(taskID, 10)
	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `
	UPDATE scheduler
	SET date = ?, title = ?, comment = ?, repeat = ?
	WHERE id = ?
	`

	res, err := DB.Exec(query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
