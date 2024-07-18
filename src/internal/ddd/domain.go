package ddd

import (
	"time"
)

type BaseDomain struct {
	ID        int       `json:"id" db:"id"`
	PID       string    `json:"pid" db:"pid"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}

func NewBaseDomain(id int, pid string, updatedAt time.Time, createdAt time.Time) BaseDomain {
	return BaseDomain{id, pid, updatedAt, createdAt}
}
