package game

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const (
	ActionPrioritySwitch   = 10
	ActionPriorityProtect  = 4
	ActionPriorityVeryFast = 2
	ActionPriorityFast     = 1
	ActionPriorityDefault  = 0
	ActionPrioritySlow     = -1
)

type ActionJutsu string

const (
	Bukijutsu ActionJutsu = "bukijutsu"
	Fuinjutsu ActionJutsu = "fuinjutsu"
	Genjutsu  ActionJutsu = "genjutsu"
	Ninjutsu  ActionJutsu = "ninjutsu"
	Senjutsu  ActionJutsu = "senjutsu"
	Taijutsu  ActionJutsu = "taijutsu"
)

type ActionConfig struct {
	Accuracy    *int        `json:"accuracy"`
	Cooldown    *int        `json:"cooldown"`
	Cost        *int        `json:"cost"`
	LifeSteal   *float64    `json:"life_steal"`
	Name        string      `json:"name"`
	Nature      *NatureSet  `json:"nature"`
	Power       *int        `json:"power"`
	Recoil      *float64    `json:"recoil"`
	Stat        *AttackStat `json:"stat"`
	TargetCount *int        `json:"target_count"`
	Jutsu       ActionJutsu `json:"jutsu"`
}

type ActionTargetType string

const (
	TargetActorID    ActionTargetType = "target-actor-id"
	TargetPositionID ActionTargetType = "target-position-type"
)

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
	ID              uuid.UUID                   `json:"ID"`
	Config          ActionConfig                `json:"config"`
	TargetType      ActionTargetType            `json:"target_type"`
	TargetPredicate func(Actor, Context) bool   `json:"-"`
	ContextValidate func(Context) bool          `json:"-"`
	MapContext      func(Game, Context) Context `json:"-"`
	Cost            GameMutation                `json:"-"`
}

func ResolveAction(game *Game, transaction Transaction[Action]) []Transaction[GameMutation] {
	if !transaction.Mutation.Filter(*game, transaction.Context) {
		fmt.Printf("Action Failed: %s +%v\n", transaction.Mutation.Config.Name, transaction.Context)
		return []Transaction[GameMutation]{}
	}

	if transaction.Mutation.Config.Cooldown != nil {
		game.SetActionCooldown(
			*transaction.Context.SourceActorID,
			transaction.Mutation.ID,
			*transaction.Mutation.Config.Cooldown,
		)
	}

	context := transaction.Context
	if transaction.Mutation.MapContext != nil {
		context = transaction.Mutation.MapContext(*game, context)
	}

	return transaction.Mutation.Delta(*game, context)
}

func GetAccuracy(game Game, source ResolvedActor, target ResolvedActor) float64 {
	ratio := float64(source.Stats[StatAccuracy]) / float64(target.Stats[StatEvasion])
	return ratio
}

func MakeActionRoll() int {
	return rand.Intn(100)
}

type AccuracyResult struct {
	Accuracy int
	Roll     int
	Success  bool
}
