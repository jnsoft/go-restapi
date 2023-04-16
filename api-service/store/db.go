package database

// A map to store the items with the uuid as the key
// This acts as the storage in lieu of an actual database
var item_store = make(map[string]*Item)
