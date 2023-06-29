package common

import (
	"github.com/google/uuid"
)

type EntityBase struct {
	ID uuid.UUID `json:"id" bson:"_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}

func NewAbstractEntity(idString string) *EntityBase {
	var id uuid.UUID

	if idString != "" {
		id = uuid.Must(uuid.FromBytes([]byte(idString)))
	} else {
		id = uuid.New()
	}

	var entity = &EntityBase{
		ID: id,
	}

	return entity
}
