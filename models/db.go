package models

import (
	"github.com/gofrs/uuid"
)

// TODO: Daniel change this based on schema

type User struct {
	ID               uuid.UUID `gorm:"PRIMARY_KEY"`
	Name             string    `gorm:"unique"`
	FastestSolveTime int
	Records          []MazeRecord
}

type MazeRecord struct {
	ID        uuid.UUID `gorm:"PRIMARY_KEY"`
	UserId    uuid.UUID
	SolveTime int
}
