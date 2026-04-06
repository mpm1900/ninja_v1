package game

import (
	"github.com/google/uuid"
)

func NotNewlyInactiveFilter(game Game, actor Actor, context Context) bool {
	return actor.InactiveTurns > 0
}

func SwitchFilter(game Game, actor Actor, context Context) bool {
	possible := game.GetActorsFilters(context, ComposeAF(TeamFilter, InactiveFilter, AliveFilter, NotNewlyInactiveFilter))
	if len(possible) == 0 {
		return true
	}

	return NotNewlyInactiveFilter(game, actor, context)
}

func NewNoopSource(groupID *uuid.UUID) ActorMutation {
	return MakeActorMutation(
		groupID,
		0,
		SourceFilter,
		func(g Game, a Actor, c Context) Actor {
			return a
		},
	)
}
func NewNoopParent(groupID *uuid.UUID) ActorMutation {
	return MakeActorMutation(
		groupID,
		0,
		ParentFilter,
		func(g Game, a Actor, c Context) Actor {
			return a
		},
	)
}

var SwitchPositions = GameMutation{
	Delta: func(p Game, g Game, context Context) Game {
		source, ok := g.GetSource(context)
		if !ok || source.PositionID == nil {
			return g
		}

		targets := g.GetTargets(context)
		if len(targets) == 0 {
			return g
		}

		g.SetPosition(targets[0], source.PositionID)
		return g
	},
}

var SetPositions = GameMutation{
	Delta: func(p Game, g Game, context Context) Game {
		targets := g.GetTargets(context)
		positionCount := len(context.TargetPositionIDs)
		if len(targets) == 0 || positionCount == 0 {
			return g
		}

		limit := min(len(targets), positionCount)
		for i := range limit {
			positionID := context.TargetPositionIDs[i]
			g.SetPosition(targets[i], &positionID)
		}

		return g
	},
}

var RemovePositions = GameMutation{
	Delta: func(p Game, g Game, context Context) Game {
		targets := g.GetTargets(context)
		for i := range len(targets) {
			g.SetPosition(targets[i], nil)
		}

		return g
	},
}

var Switch = Action{
	ID: uuid.New(),
	Config: ActionConfig{
		Name:        "Switch",
		Description: "Switches user out and target ally into battle.",
	},
	Locked:          true,
	TargetType:      TargetActorID,
	TargetPredicate: ComposeAF(TeamFilter, InactiveFilter, AliveFilter, SwitchFilter),
	ContextValidate: TargetLengthFilter(1),
	ActionMutation: ActionMutation{
		Priority: ActionPrioritySwitch,
		Filter:   TrueGameFilter,
		Delta: func(p Game, input Game, context Context) []Transaction[GameMutation] {
			transactions := []GameTransaction{
				MakeTransaction(SwitchPositions, context),
			}
			return transactions
		},
	},
}

var SwitchInIds map[int]uuid.UUID = map[int]uuid.UUID{
	1: uuid.New(),
	2: uuid.New(),
	3: uuid.New(),
	4: uuid.New(),
	5: uuid.New(),
}

func SwitchIn(count int) Action {
	return Action{
		ID: SwitchInIds[count],
		Config: ActionConfig{
			Name: "Switch In",
		},
		Locked:          true,
		TargetType:      TargetActorID,
		TargetPredicate: ComposeAF(TeamFilter, InactiveFilter, AliveFilter, SwitchFilter),
		ContextValidate: TargetLengthFilter(count),
		ActionMutation: ActionMutation{
			Priority: ActionPrioritySwitch,
			Filter:   TrueGameFilter,
			Delta: func(parent Game, input Game, context Context) []Transaction[GameMutation] {
				transactions := []GameTransaction{}
				min := min(len(context.TargetActorIDs), len(context.TargetPositionIDs))
				for i := range min {
					c := context
					c.TargetActorIDs = []uuid.UUID{c.TargetActorIDs[i]}
					c.TargetPositionIDs = []uuid.UUID{c.TargetPositionIDs[i]}
					transactions = append(transactions, MakeTransaction(SetPositions, c))
				}

				return transactions
			},
		},
	}
}
