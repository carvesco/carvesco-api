package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Serendipia struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Title   string             `json:"title,omitempty" validate:"required"`
	Type    string             `json:"type,omitempty" validate:"required"`
	Details string             `json:"details,omitempty" validate:"required"`
	Review  string             `json:"review,omitempty" validate:"required"`
	Image   string             `json:"image,omitempty" validate:"required"`
	Link    string             `json:"link,omitempty" validate:"required"`
}
