package modifiers

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var addNatureID = uuid.New()

func AddNature(nature game.NatureSet) game.Modifier {
	return game.Modifier{
		ID:       addNatureID,
		Name:     fmt.Sprintf("Add Nature: %s", nature),
		Duration: 0,
		Mutations: []game.ActorMutation{
			game.MakeActorMutation(
				nil,
				game.MutPriorityDefault,
				game.ComposeAF(game.SourceFilter, game.ActiveFilter),
				func(a game.Actor, c game.Context) game.Actor {
					a.Natures[nature] = game.NATURES[nature]
					return a
				},
			),
		},
	}
}

var removeNatureID = uuid.New()

func RemoveNature(nature game.NatureSet) game.Modifier {
	return game.Modifier{
		ID:       removeNatureID,
		Name:     fmt.Sprintf("Add Nature: %s", nature),
		Duration: 0,
		Mutations: []game.ActorMutation{
			game.MakeActorMutation(
				nil,
				game.MutPriorityDefault,
				game.ComposeAF(game.SourceFilter, game.ActiveFilter),
				func(a game.Actor, c game.Context) game.Actor {
					delete(a.Natures, nature)
					return a
				},
			),
		},
	}
}
