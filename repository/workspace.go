package repository

import (
	"fmt"
	"travel-cli/db"
	"travel-cli/model"

	"gorm.io/gorm"
)

func CreateWorkspace(name string) (*model.Workspace, error) {
	workspace := &model.Workspace{Name: name}
	if err := db.DB.Create(workspace).Error; err != nil {
		return nil, err
	}

	return workspace, nil
}

func ListWorkspaces() ([]model.Workspace, error) {
	var workspaces []model.Workspace
	if err := db.DB.Order("current DESC").Find(&workspaces).Error; err != nil {
		return nil, err
	}
	return workspaces, nil
}

func SetCurrentWorkspace(workspaceName string) error {
	var workspace model.Workspace

	findError := db.DB.Where("name = ?", workspaceName).First(&workspace).Error
	if findError != nil {
		if findError == gorm.ErrRecordNotFound {
			return fmt.Errorf("workspace '%s' not found. Please create it first or set an existing workspace", workspaceName)
		}
		return findError
	}

	if err := db.DB.Model(&model.Workspace{}).
		Where("id_workspace != ?", workspace.IdWorkspace).
		Update("current", 0).Error; err != nil {
		return fmt.Errorf("failed to unset other workspaces: %w", err)
	}

	if err := db.DB.Model(&workspace).Update("current", 1).Error; err != nil {
		return fmt.Errorf("failed to set workspace '%s' as current: %w", workspaceName, err)
	}

	return nil
}
