package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catchphrase struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    title        string             `bson:"title,omitempty"`
    description  string             `bson:"description,omitempty"`
}
