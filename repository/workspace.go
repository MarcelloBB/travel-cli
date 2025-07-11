package repository

import (
	"travel-cli/db"
	"travel-cli/model"
)

func CreateWorkspace(name string) (*model.Workspace, error) {
	workspace := &model.Workspace{Name: name}
	if err := db.DB.Create(workspace).Error; err != nil {
		return nil, err
	}

	return workspace, nil
}
