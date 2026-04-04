package mutations

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func AddModifiers(modifiers ...game.Modifier) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			for _, modifier := range modifiers {
				mod_tx := game.MakeTransaction(modifier, context)
				g.Modifiers = append(g.Modifiers, mod_tx)

				// logs
				for _, actor := range g.GetActionableActors() {
					if game.CheckModifier(mod_tx, actor) {
						context.SourceActorID = &actor.ID
						g.PushLog(game.NewLogContext(fmt.Sprintf(">>> $source$ gained %s", modifier.Name), context))
					}
				}
			}

			return g
		},
	}
}

func RemoveModifierTxByID(ID uuid.UUID) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
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
