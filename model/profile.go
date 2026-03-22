package model

type UserProfile struct {
	ID             string         `json:"-" bson:"_id"`
	Name           string         `json:"name" bson:"name"`
	Email          string         `json:"email" bson:"email"`
	Location       string         `json:"location" bson:"location"`
	PhoneNumber    string         `json:"phone_number" bson:"phone_number"`
	DynamicContent map[string]any `json:"dynamic_content" bson:"dynamic_content"`
	Occupation     Occupation     `json:"occupation" bson:"occupation"`
}
