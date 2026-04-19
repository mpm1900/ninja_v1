package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var shadowCageID = uuid.MustParse("4494f97f-ac70-4fa5-b599-633aa13536f7")

var ShadowCage game.Modifier = game.Modifier{
	ID:          shadowCageID,
	GroupID:     &shadowCageID,
	Name:        "Shadow Cage",
	Description: "Other enemy shinobi: Cannot switch out.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&shadowCageID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.OtherFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.SwitchLocked = true
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
