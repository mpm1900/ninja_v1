package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var powerBoostedID = uuid.New()
var PowerBoosted game.Modifier = game.Modifier{
	ID:          uuid.New(),
	GroupID:     &powerBoostedID,
	Name:        "Power Boosted",
	Description: "Attack Power +50%",
	Show:        true,
	Duration:    0,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&powerBoostedID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.TargetFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.PowerMultiplier += 0.5
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
