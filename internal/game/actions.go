package game

import (
	"github.com/google/uuid"
)

func NewNoop(groupID *uuid.UUID) ActorMutation {
	return MakeActorMutation(
		groupID,
		0,
		SourceFilter,
		func(a Actor, c Context) Actor {
			return a
		},
	)
}

var SwitchPositions = GameMutation{
	Delta: func(g Game, context Context) Game {
		ok, source := g.GetSource(context)
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
	Delta: func(g Game, context Context) Game {
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
	Delta: func(g Game, context Context) Game {
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
		Name: "Switch",
	},
	TargetType:      TargetActorID,
	TargetPredicate: ComposeAF(TeamFilter, InactiveFilter, AliveFilter),
	ContextValidate: TargetLengthFilter(1),
	ActionMutation: ActionMutation{
		Priority: ActionPrioritySwitch,
		Filter:   AllGameFilter,
		Delta: func(input Game, context Context) []Transaction[GameMutation] {
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
		TargetType:      TargetActorID,
		TargetPredicate: ComposeAF(TeamFilter, InactiveFilter, AliveFilter),
		ContextValidate: TargetLengthFilter(count),
		ActionMutation: ActionMutation{
			Priority: ActionPrioritySwitch,
			Filter:   AllGameFilter,
			Delta: func(input Game, context Context) []Transaction[GameMutation] {
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
