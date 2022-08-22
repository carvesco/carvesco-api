package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	FromAddress string             `json:"fromaddress,omitempty" validate:"required"`
	ToAddress   string             `json:"toaddress,omitempty" validate:"required"`
	Message     string             `json:"message,omitempty" validate:"required"`
}
