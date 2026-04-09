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
	ID: uuid.MustParse("2c6bf2e3-4049-4f1f-b18a-f4fd844ef31a"),
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
	1: uuid.MustParse("f4e25f49-35c1-405d-a34a-c8ba22f595f8"),
	2: uuid.MustParse("11a552fc-a768-4f58-aef1-fecba66d95f8"),
	3: uuid.MustParse("6bc3023c-2b4d-4bc3-8712-dc9f6124f70e"),
	4: uuid.MustParse("6a842cd4-ab84-4a8a-9f8b-f964d5f87f72"),
	5: uuid.MustParse("a8bdb03b-b6bf-4425-8de3-c3ce9e541f32"),
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
