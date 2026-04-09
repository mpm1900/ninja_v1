package mutations

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func AddModifiers(checkProtect bool, modifiers ...game.Modifier) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			for _, modifier := range modifiers {
				mod_tx := game.MakeTransaction(modifier, context)
				mod_tx.Context.ModifierID = &mod_tx.ID

				// logs
				if context.SourceActorID == nil && len(modifier.GameStateMutations) > 0 {
					g.PushLog(game.NewLog(fmt.Sprintf(">>> The battlefield gained %s.", mod_tx.Mutation.Name)))
				}
				for _, actor := range g.GetActionableActors() {
					if game.CheckModifierForActor(mod_tx, g, actor) {
						resolved := actor.Resolve(g)
						if checkProtect && resolved.Protected {
							mod_tx.Context.FilterOutTarget(actor)

							context.SourceActorID = &actor.ID
							g.PushLog(game.NewLogContext(">>> $source$ was protected.", context.WithSource(actor.ID)))
							continue
						} else {
							g.PushLog(game.NewLogContext(fmt.Sprintf(">>> $source$ gained %s.", modifier.Name), context.WithSource(actor.ID)))
						}
					}
				}

				g.On(game.OnModifierAdd, &mod_tx.Context)
				g.Modifiers = append(g.Modifiers, mod_tx)
			}

			return g
		},
	}
}

/**
 * currently not used
 */
func RemoveModifierTxByID(ID uuid.UUID) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			modifiers := []game.Transaction[game.Modifier]{}
			for _, tx := range g.Modifiers {
				if tx.ID != ID {
					modifiers = append(modifiers, tx)
				}
			}

			g.Modifiers = modifiers
			return g
		},
	}
}

var ConsumeItem game.GameMutation = game.GameMutation{
	Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
		if context.SourceActorID == nil {
			return g
		}

		g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
			a.Item = nil
			return a
		})

		return g
	},
}
