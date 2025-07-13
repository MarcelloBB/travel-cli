package repository

import (
	"travel-cli/db"
	"travel-cli/model"
)

func CreateCollection(name string) (*model.Collection, error) {
	currentWorkspace, err := GetCurrentWorkspace()
	if err != nil {
		return nil, err
	}
	collection := &model.Collection{Title: name, IdWorkspace: currentWorkspace.IdWorkspace}
	if err := db.DB.Create(collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

func ListCollections(name ...string) ([]model.Collection, error) {
	var collections []model.Collection
	query := db.DB.Preload("Requests")

	if len(name) > 0 && name[0] != "" {
		query = query.Where("title = ?", name[0])
	}

	if err := query.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func SaveRequestToCollection(requestToSave *model.Request, collectionName string) error {
	currentWorkspace, err := GetCurrentWorkspace()
	if err != nil {
		return err
	}

	var collection model.Collection
	if err := db.DB.Where("title = ? AND id_workspace = ?", collectionName, currentWorkspace.IdWorkspace).First(&collection).Error; err != nil {
		return err
	}

	requestToSave.IdCollection = collection.IdCollection
	if err := db.DB.Create(&requestToSave).Error; err != nil {
		return err
	}

	return nil
}
