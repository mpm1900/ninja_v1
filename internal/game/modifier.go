package game

import (
	"github.com/google/uuid"
)

const ModifierDurationInf = -1

type ModifierMetadata struct {
	Hazard bool
	Status bool
}

type Modifier struct {
	ID       uuid.UUID  `json:"ID"`
	GroupID  *uuid.UUID `json:"group_ID"`
	Name     string     `json:"name"`
	Duration int        `json:"duration"`

	ActorMutations     []ActorMutation     `json:"-"`
	GameStateMutations []GameStateMutation `json:"-"`
	Triggers           []Trigger           `json:"-"`
}

func ResolveTrigger(game Game, transaction Transaction[Trigger]) []Transaction[GameMutation] {
	if transaction.Mutation.Delta == nil {
		return []Transaction[GameMutation]{}
	}

	return transaction.Mutation.Delta(game, game, transaction.Context)
}

func CheckModifierForActor(tx Transaction[Modifier], game Game, actor Actor) bool {
	game.Actors = []Actor{actor}
	for _, mut := range tx.Mutation.ActorMutations {
		if mut.Filter(game, actor, tx.Context) {
			return true
		}
	}

	return false
}
