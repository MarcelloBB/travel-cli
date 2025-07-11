package model

type Collection struct {
	IdCollection uint   `gorm:"primaryKey" json:"id_collection"`
	IdWorkspace  uint   `json:"id_workspace"`
	Title        string `json:"title"`
}
