package store

import "go-restapi/api-service/models"

// A map to store the items with the uuid as the key
// This acts as the storage in lieu of an actual database
var Item_store = make(map[string]*models.Item)
