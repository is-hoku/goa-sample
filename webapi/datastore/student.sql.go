// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: student.sql

package datastore

import (
	"context"
	"database/sql"
	"time"
)

const getAllStudents = `-- name: GetAllStudents :many
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `ruby` + "`" + `, ` + "`" + `student_number` + "`" + `, ` + "`" + `date_of_birth` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `expiration_date` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + ` FROM ` + "`" + `students` + "`" + ` ORDER BY ` + "`" + `student_number` + "`" + ` ASC
`

func (q *Queries) GetAllStudents(ctx context.Context) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getAllStudents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Ruby,
			&i.StudentNumber,
			&i.DateOfBirth,
			&i.Address,
			&i.ExpirationDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentByID = `-- name: GetStudentByID :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `ruby` + "`" + `, ` + "`" + `student_number` + "`" + `, ` + "`" + `date_of_birth` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `expiration_date` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + ` FROM ` + "`" + `students` + "`" + ` WHERE ` + "`" + `id` + "`" + `=?
`

func (q *Queries) GetStudentByID(ctx context.Context, id uint64) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByID, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ruby,
		&i.StudentNumber,
		&i.DateOfBirth,
		&i.Address,
		&i.ExpirationDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStudentByNumber = `-- name: GetStudentByNumber :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `ruby` + "`" + `, ` + "`" + `student_number` + "`" + `, ` + "`" + `date_of_birth` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `expiration_date` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + ` FROM ` + "`" + `students` + "`" + ` WHERE ` + "`" + `student_number` + "`" + `=?
`

func (q *Queries) GetStudentByNumber(ctx context.Context, studentNumber uint32) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByNumber, studentNumber)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ruby,
		&i.StudentNumber,
		&i.DateOfBirth,
		&i.Address,
		&i.ExpirationDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setStudent = `-- name: SetStudent :execresult
INSERT INTO ` + "`" + `students` + "`" + ` (` + "`" + `name` + "`" + `, ` + "`" + `ruby` + "`" + `, ` + "`" + `student_number` + "`" + `, ` + "`" + `date_of_birth` + "`" + `, ` + "`" + `address` + "`" + `, ` + "`" + `expiration_date` + "`" + `) VALUES(?, ?, ?, ?, ?, ?)
`

type SetStudentParams struct {
	Name           string
	Ruby           string
	StudentNumber  uint32
	DateOfBirth    time.Time
	Address        string
	ExpirationDate time.Time
}

func (q *Queries) SetStudent(ctx context.Context, arg SetStudentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, setStudent,
		arg.Name,
		arg.Ruby,
		arg.StudentNumber,
		arg.DateOfBirth,
		arg.Address,
		arg.ExpirationDate,
	)
}
