package db

import (
	"common"
)

// TODO do all the methods

// CRUD db methods
// Create
func CreateNewRecord(name, description string, projectId int) (common.Subscription, error) {
	return common.Subscription{}, nil
}

// Read
func ListRecords(offset, upperLimit int) ([]common.Subscription, error) {
	return []common.Subscription{}, nil
}

// Update
func UpdateRecord(name, description string, id, projectId int, query string) (common.Subscription, error) {
	return common.Subscription{}, nil
}

// Delete
func DeleteRecord(id int) error {
	return nil
}
