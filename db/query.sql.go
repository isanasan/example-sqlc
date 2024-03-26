// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  name, email, age
) VALUES (
  ?, ?, ?
)
`

type CreateUserParams struct {
	Name  string
	Email string
	Age   int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Name, arg.Email, arg.Age)
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, age FROM users
WHERE id = ?
`

func (q *Queries) GetUser(ctx context.Context, id uint64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Age,
	)
	return i, err
}

const updateUserAges = `-- name: UpdateUserAges :exec
UPDATE users SET age = ?
WHERE id = ?
`

type UpdateUserAgesParams struct {
	Age int32
	ID  uint64
}

func (q *Queries) UpdateUserAges(ctx context.Context, arg UpdateUserAgesParams) error {
	_, err := q.db.ExecContext(ctx, updateUserAges, arg.Age, arg.ID)
	return err
}
