package domains

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"username" json:"username" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`

	Mobile string `bson:"mobile" json:"mobile" validate:"required"`
	Role string `bson:"role" json:"role" validate:"required"`
	
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserProperty struct {
	Username string `bson:"username" json:"username" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`

	Mobile string `bson:"mobile" json:"mobile" validate:"required"`
	Role string `bson:"role" json:"role" validate:"required"`
	
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserNamePassword struct {
	Username string `bson:"username" json:"username" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type UserMetadata struct {
	Mobile string `bson:"mobile" json:"mobile" validate:"required"`
	Role string `bson:"role" json:"role" validate:"required"`
}