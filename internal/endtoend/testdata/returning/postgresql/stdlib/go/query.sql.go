// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const deleteUserAndReturnID = `-- name: DeleteUserAndReturnID :one
DELETE FROM users
  WHERE name = $1
  RETURNING id
`

func (q *Queries) DeleteUserAndReturnID(ctx context.Context, name sql.NullString) (int32, error) {
	row := q.db.QueryRowContext(ctx, deleteUserAndReturnID, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserAndReturnUser = `-- name: DeleteUserAndReturnUser :one
DELETE FROM users
  WHERE name = $1
  RETURNING name, id
`

func (q *Queries) DeleteUserAndReturnUser(ctx context.Context, name sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUserAndReturnUser, name)
	var i User
	err := row.Scan(&i.Name, &i.ID)
	return i, err
}

const insertUserAndReturnID = `-- name: InsertUserAndReturnID :one
INSERT INTO users (name) VALUES ($1)
  RETURNING id
`

func (q *Queries) InsertUserAndReturnID(ctx context.Context, name sql.NullString) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertUserAndReturnID, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertUserAndReturnUser = `-- name: InsertUserAndReturnUser :one
INSERT INTO users (name) VALUES ($1)
  RETURNING name, id
`

func (q *Queries) InsertUserAndReturnUser(ctx context.Context, name sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUserAndReturnUser, name)
	var i User
	err := row.Scan(&i.Name, &i.ID)
	return i, err
}

const updateUserAndReturnID = `-- name: UpdateUserAndReturnID :one
UPDATE users SET name = $1
  WHERE name = $2
  RETURNING id
`

type UpdateUserAndReturnIDParams struct {
	Name   sql.NullString
	Name_2 sql.NullString
}

func (q *Queries) UpdateUserAndReturnID(ctx context.Context, arg UpdateUserAndReturnIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateUserAndReturnID, arg.Name, arg.Name_2)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateUserAndReturnUser = `-- name: UpdateUserAndReturnUser :one
UPDATE users SET name = $1
  WHERE name = $2
  RETURNING name, id
`

type UpdateUserAndReturnUserParams struct {
	Name   sql.NullString
	Name_2 sql.NullString
}

func (q *Queries) UpdateUserAndReturnUser(ctx context.Context, arg UpdateUserAndReturnUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserAndReturnUser, arg.Name, arg.Name_2)
	var i User
	err := row.Scan(&i.Name, &i.ID)
	return i, err
}