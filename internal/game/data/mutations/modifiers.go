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

				// logs
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

				g.Modifiers = append(g.Modifiers, mod_tx)
			}

			return g
		},
	}
}

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
