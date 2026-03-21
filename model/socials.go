package model

type Social struct {
	IconBase
	Handle string `json:"handle" bson:"handle"`
}
