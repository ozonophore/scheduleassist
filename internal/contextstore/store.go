package contextstore

import (
	"context"
	"sync"
)

var contextStorage ContextStorage

type UserContext struct {
	UserID   int64  `json:"user_id"`   // ID пользователя
	State    string `json:"state"`     // Текущее состояние (например, "waiting_for_name")
	LastSeen int64  `json:"last_seen"` // Временная метка последнего взаимодействия
	Context  context.Context
}

type ContextStorage struct {
	memoryStore sync.Map
}

func Get(id int64) *UserContext {
	if val, ok := contextStorage.memoryStore.Load(id); ok {
		return val.(*UserContext)
	} else {
		ctx := UserContext{}
		contextStorage.memoryStore.Store(id, &ctx)
		return &ctx
	}
}
