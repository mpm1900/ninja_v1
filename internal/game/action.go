package game

import (
	"math"
	"math/rand"

	"github.com/google/uuid"
)

const (
	ActionPrioritySwitch  = 10
	ActionPriorityFast    = 1
	ActionPriorityDefault = 0
	ActionPrioritySlow    = -1
)

type ActionConfig struct {
	Name     string      `json:"name"`
	Nature   *NatureSet  `json:"nature"`
	Accuracy *int        `json:"accuracy"`
	Stat     *AttackStat `json:"stat"`
	Power    *int        `json:"power"`
	Recoil   *float64    `json:"recoil"`
}

type ActionMutation Mutation[Game, []Transaction[GameMutation]]

/** [This comment was not written by an LLM]
 * Action Function Members for Action "a"
 *
 * action.Filter(Game, *Context) => can this action be taken with this countext.
 * -- this is often done for a chakra or disabled check
 *
 * action.TargetPredicate(Actor, *Context) => is this actor a valid target for this action
 * -- this is effectively the "targets generator" for users to choose.
 *
 * action.ContextValidate(*Context) => does this context represent a valid targets selection for this action
 * -- this is used to check "is the number of targets correct?" and other checks.
 *
 *
 * action.Delta(Game, *Context) => resolution of the Action
 * -- can include random events
 */
type Action struct {
	ActionMutation
	ID              uuid.UUID                 `json:"ID"`
	Config          ActionConfig              `json:"config"`
	TargetPredicate func(Actor, Context) bool `json:"-"`
	ContextValidate func(Context) bool        `json:"-"`
}

func ResolveAction(game Game, transaction Transaction[Action]) []Transaction[GameMutation] {
	return transaction.Mutation.Delta(game, transaction.Context)
}

func GetAccuracy(game Game, sourceID uuid.UUID, actionAccuracy *int, modifier float64) int {
	if actionAccuracy == nil {
		return 100
	}

	ok, s := game.GetActorByID(sourceID)
	if !ok {
		return 0
	}
	source := s.Resolve(game)
	base_acc := float64(source.Stats[StatAccuracy]) / 100.0
	move_acc := float64(*actionAccuracy) / 100.0
	accuracy := int(math.Floor(base_acc * move_acc * 100.0 * modifier))

	return accuracy
}

func MakeActionRoll() int {
	return rand.Intn(100)
}

type AccuracyResult struct {
	Accuracy int
	Roll     int
	Success  bool
}

func GetAccuracyResult(game Game, actorID uuid.UUID, move *int) AccuracyResult {
	accuracy := GetAccuracy(game, actorID, move, 1.0)
	roll := MakeActionRoll()
	return AccuracyResult{
		Accuracy: accuracy,
		Roll:     roll,
		Success:  accuracy >= roll,
	}
}
