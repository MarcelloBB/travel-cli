package model

type Workspace struct {
	IdWorkspace uint   `gorm:"primaryKey" json:"id_workspace"`
	Name        string `json:"name"`
}
