// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package postgresql

import (
	"context"

	"github.com/google/uuid"
)

const getUser = `-- name: GetUser :one
SELECT id, email, name, password, status, application_id, group_id FROM users
WHERE
    email = $2 AND application_id = $1
`

type GetUserParams struct {
	ApplicationID uuid.UUID
	Email         string
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (User, error) {
	row := q.db.QueryRow(ctx, getUser, arg.ApplicationID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.Password,
		&i.Status,
		&i.ApplicationID,
		&i.GroupID,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users
    ( "email", "name", "password", "application_id", "group_id" ) VALUES
    ( $1, $2, $3, $4, $5 )
`

type InsertUserParams struct {
	Email         string
	Name          string
	Password      string
	ApplicationID uuid.UUID
	GroupID       uuid.UUID
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.Email,
		arg.Name,
		arg.Password,
		arg.ApplicationID,
		arg.GroupID,
	)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT name, email, status FROM users
WHERE
    application_id = $1
ORDER BY name ASC
`

type ListUsersRow struct {
	Name   string
	Email  string
	Status string
}

func (q *Queries) ListUsers(ctx context.Context, applicationID uuid.UUID) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers, applicationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(&i.Name, &i.Email, &i.Status); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserStatus = `-- name: UpdateUserStatus :exec
UPDATE users
SET status = $3
WHERE email = $2 AND application_id = $1
`

type UpdateUserStatusParams struct {
	ApplicationID uuid.UUID
	Email         string
	Status        string
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error {
	_, err := q.db.Exec(ctx, updateUserStatus, arg.ApplicationID, arg.Email, arg.Status)
	return err
}
