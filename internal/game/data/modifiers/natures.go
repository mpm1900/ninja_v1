package modifiers

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var addNatureID = uuid.New()

func AddNature(nature game.NatureSet, duration int) game.Modifier {
	return game.Modifier{
		ID:       addNatureID,
		Name:     fmt.Sprintf("Add Nature: %s", nature),
		Show:     false,
		Duration: duration,
		ActorMutations: []game.ActorMutation{
			game.MakeActorMutation(
				nil,
				game.MutPriorityDefault,
				game.ComposeAF(game.SourceFilter, game.ActiveFilter),
				func(g game.Game, a game.Actor, c game.Context) game.Actor {
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
		Name:     fmt.Sprintf("Remove Nature: %s", nature),
		Show:     false,
		Duration: 0,
		ActorMutations: []game.ActorMutation{
			game.MakeActorMutation(
				nil,
				game.MutPriorityDefault,
				game.ComposeAF(game.SourceFilter, game.ActiveFilter),
				func(g game.Game, a game.Actor, c game.Context) game.Actor {
					delete(a.Natures, nature)
					return a
				},
			),
		},
	}
}
