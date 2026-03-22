package model

type Social struct {
	IconBase `bson:",inline" json:",inline"`
	Handle   string `json:"handle" bson:"handle"`
}
