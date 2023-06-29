package example

import (
	"time"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/google/uuid"
)

type Example struct {
	common.EntityBase `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ID                uuid.UUID            `json:"id" bson:"_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name              string               `json:"name" gorm:"type:varchar(60);column:name"`
	CreatedAt         common.JsonTime      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         common.JsonTime      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt         *common.JsonNullTime `json:"deleted_at" gorm:"column:deleted_at"`
}

type ExampleProps struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e *Example) TableName() string {
	return "example"
}

func (e *Example) BeforeCreate() error {
	id := uuid.New()

	e.ID = id

	return nil
}

func NewExampleEntity(props ExampleProps) (*Example, error) {
	var example *Example

	abstractEntity := common.NewAbstractEntity(props.ID)

	if common.IsNullOrEmpty(props.Name) {
		return nil, common.IsNullOrEmptyError("name")
	}

	actualDate := time.Now()

	example = &Example{
		Name:      props.Name,
		CreatedAt: common.JsonTime{Value: actualDate},
		UpdatedAt: common.JsonTime{Value: actualDate},
	}

	example.ID = abstractEntity.ID

	return example, nil
}
