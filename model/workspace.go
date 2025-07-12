package model

type Workspace struct {
	IdWorkspace uint   `gorm:"primaryKey" json:"id_workspace"`
	Name        string `json:"name"`
	Current     uint   `gorm:"default:0" json:"current"`
}
