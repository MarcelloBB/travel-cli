package repository

import (
	"travel-cli/db"
	"travel-cli/model"
)

func CreateCollection(name string, idWorkspace uint) (*model.Collection, error) {
	collection := &model.Collection{Title: name, IdWorkspace: idWorkspace}
	if err := db.DB.Create(collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

func ListCollections() ([]model.Collection, error) {
	var collections []model.Collection
	if err := db.DB.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}
