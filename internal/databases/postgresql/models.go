// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgresql

import (
	"github.com/google/uuid"
)

type Application struct {
	ID     uuid.UUID
	Name   string
	Secret uuid.UUID
}

type Group struct {
	ID            uuid.UUID
	Name          string
	ApplicationID uuid.UUID
	Permissions   []byte
}

type User struct {
	ID            uuid.UUID
	Email         string
	Name          string
	Password      string
	Status        string
	ApplicationID uuid.UUID
	GroupID       uuid.UUID
}
