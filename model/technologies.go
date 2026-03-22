package model

type Technology struct {
	IconBase `bson:",inline" json:",inline"`
	Type     string `bson:"type" json:"type"`
}
