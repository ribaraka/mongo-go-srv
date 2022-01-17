package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" validate:"required,lte=10" bson:"first_name"`
	LastName  string             `json:"lastName" validate:"required,lte=10" bson:"last_name"`
	Email     string             `json:"email" validate:"required,email" bson:"email"`
	Password  string             `json:"password" validate:"required,gte=8,lte=64" bson:"password"`
}

type Login struct {
	Email    string `json:"emailIn" validate:"required,email"`
	Password string `json:"passwordIn" validate:"required,gte=8,lte=64"`
}
