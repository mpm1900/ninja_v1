package game

import (
	"github.com/google/uuid"
)

type ActionMutation[T any] struct {
	Mutation[T, []Transaction[T, T, Context], Context]
}
type ActionTransaction[T any] struct {
	Transaction[T, []Transaction[T, T, Context], Context]
}
type ActionConfig struct{}
type Action[T any] struct {
	ActionMutation[T] `json:"-"`
	ID                uuid.UUID             `json:"ID"`
	Name              string                `json:"name"`
	Config            ActionConfig          `json:"config"`
	predicate         func(T, Context) bool `json:"-"`
	validate          func(Context) bool    `json:"-"`
}
