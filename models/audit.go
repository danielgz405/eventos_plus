package models

import "time"

type AuditLog struct {
	User  Profile   `bson:"user"`
	Type  string    `bson:"type"`
	Table string    `bson:"table"`
	Date  time.Time `bson:"date"`
}
