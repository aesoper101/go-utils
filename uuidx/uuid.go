package uuidx

import (
	uuid "github.com/satori/go.uuid"
)

// MustUUID returns a new UUIDv4
func MustUUID() string {
	u := uuid.NewV4()
	return u.String()
}
