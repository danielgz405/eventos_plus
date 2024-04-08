package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	Image       string             `bson:"image" json:"image"`
	ImageRef    string             `bson:"image_ref" json:"image_ref"`
	PaymentData PaymentData        `bson:"payment_data" json:"payment_data"`
	IsCreator   bool               `bson:"is_creator" json:"is_creator"`
	PaypalEmail string             `bson:"paypal_email" json:"paypal_email"`
}

type Profile struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Image       string             `bson:"image" json:"image"`
	ImageRef    string             `bson:"image_ref" json:"image_ref"`
	PaymentData PaymentData        `bson:"payment_data" json:"payment_data"`
	IsCreator   bool               `bson:"is_creator" json:"is_creator"`
	PaypalEmail string             `bson:"paypal_email" json:"paypal_email"`
}

type InsertUser struct {
	Name        string      `bson:"name" json:"name"`
	Email       string      `bson:"email" json:"email"`
	Password    string      `bson:"password" json:"password"`
	Image       string      `bson:"image" json:"image"`
	ImageRef    string      `bson:"image_ref" json:"image_ref"`
	PaymentData PaymentData `bson:"payment_data" json:"payment_data"`
	IsCreator   bool        `bson:"is_creator" json:"is_creator"`
	PaypalEmail string      `bson:"paypal_email" json:"paypal_email"`
}

type UpdateUser struct {
	Id          string      `bson:"_id" json:"_id"`
	Name        string      `bson:"name" json:"name"`
	Email       string      `bson:"email" json:"email"`
	Image       string      `bson:"image" json:"image"`
	ImageRef    string      `bson:"image_ref" json:"image_ref"`
	PaymentData PaymentData `bson:"payment_data" json:"payment_data"`
	IsCreator   bool        `bson:"is_creator" json:"is_creator"`
	PaypalEmail string      `bson:"paypal_email" json:"paypal_email"`
}

type PaymentData struct {
	Number         string `bson:"number" json:"number"`
	Name           string `bson:"name" json:"name"`
	ExpirationDate string `bson:"expiration_date" json:"expiration_date"`
	Cvv            string `bson:"cvv" json:"cvv"`
}
