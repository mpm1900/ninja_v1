package game

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const (
	ActionPrioritySwitch  = 10
	ActionPriorityProtect = 4
	ActionPriorityP3      = 3
	ActionPriorityP2      = 2
	ActionPriorityP1      = 1
	ActionPriorityDefault = 0
	ActionPrioritySlow    = -1
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
	Description string      `json:"description"`
	LogSuccessF *string     `json:"-"`
	LogFailureF *string     `json:"-"`
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

func ResolveAction(game *Game, transaction Transaction[Action]) []GameTransaction {
	action := transaction.Mutation
	if !action.Filter(*game, transaction.Context) {
		log := fmt.Sprintf("%s failed.", action.Config.Name)
		if action.Config.LogFailureF != nil {
			log = fmt.Sprintf(*action.Config.LogFailureF, action.Config.Name)
		}
		game.PushLog(log)
		return []GameTransaction{}
	}

	if action.Config.Cooldown != nil {
		game.SetActionCooldown(
			*transaction.Context.SourceActorID,
			action.ID,
			*action.Config.Cooldown,
		)
	}

	context := transaction.Context
	if action.MapContext != nil {
		context = action.MapContext(*game, context)
	}

	source, ok := game.GetSource(context)
	if ok {
		log := fmt.Sprintf("%s used %s.", source.Name, action.Config.Name)
		if action.Config.LogSuccessF != nil {
			log = fmt.Sprintf(*action.Config.LogSuccessF, source.Name, action.Config.Name)
		}
		game.PushLog(log)
	}

	queue, ok := game.QueuedActions[source.ID]
	if ok {
		delete(game.QueuedActions, source.ID)
		if queue.Mutation != transaction.Mutation.ID {
			fmt.Println("ERROR: INVALID ACTION EXECTUED")
			return []GameTransaction{}
		}
	}

	return action.Delta(*game, context)
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
