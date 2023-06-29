package common

import (
	"database/sql"
	"encoding/json"
	"time"
)

type IAggregateRoot interface {
}

type IDTO interface {
	Validate() (bool, error)
	ToBytes() ([]byte, error)
}

type JsonNullTime struct {
	Value sql.NullTime
}

func (v JsonNullTime) MarshalJSON() ([]byte, error) {
	if v.Value.Valid {
		return json.Marshal(v.Value.Time.Format("2006-01-02 15:04:05"))
	} else {
		return json.Marshal(nil)
	}
}

type JsonTime struct {
	Value time.Time
}

func (v JsonTime) MarshalJSON() ([]byte, error) {
	if v.Value.String() != "" {
		return json.Marshal(v.Value.Format("2006-01-02 15:04:05"))
	} else {
		return json.Marshal(nil)
	}
}
