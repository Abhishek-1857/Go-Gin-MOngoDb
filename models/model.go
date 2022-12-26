package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type KycInfo struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Country string             `json:"country,omitempty" validate:"required"`
}

type IpfsHash struct {
	IpfsHash string `json:"ipfshash"`
}
