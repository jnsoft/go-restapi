package store

import (
	"fmt"
	"go-restapi/api-service/models"

	"github.com/google/uuid"
)

type ItemStore map[uuid.UUID]*models.Item

// A map to store the items with the uuid as the key
// This acts as the storage in lieu of an actual database
var Item_store = make(map[uuid.UUID]*models.Item)

func (s ItemStore) Init(n int) {

	for i := 0; i < n; i++ {
		item := models.Item{Id: uuid.New(), Name: fmt.Sprintf("item %v", i+1)}
		s[item.Id] = &item
	}
}

func (s ItemStore) Find(id uuid.UUID) (*models.Item, bool) {
	item, ok := s[id]
	if !ok {
		return nil, false
	}
	return item, true
}
