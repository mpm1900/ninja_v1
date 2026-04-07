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
	Accuracy    *int        `json:"accuracy,omitempty"`
	Cooldown    *int        `json:"cooldown,omitempty"`
	Cost        *int        `json:"cost,omitempty"`
	CritChance  *int        `json:"crit_chance,omitempty"`
	CritMod     float64     `json:"crit_mod,omitempty"`
	LifeSteal   *float64    `json:"life_steal,omitempty"`
	Name        string      `json:"name"`
	Nature      *NatureSet  `json:"nature,omitempty"`
	Power       *int        `json:"power,omitempty"`
	Recoil      *float64    `json:"recoil,omitempty"`
	Stat        *AttackStat `json:"stat,omitempty"`
	TargetCount *int        `json:"target_count,omitempty"`
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

type ActionMutation Mutation[Game, Game, []Transaction[GameMutation]]

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
	ID              uuid.UUID                       `json:"ID"`
	Config          ActionConfig                    `json:"config"`
	Locked          bool                            `json:"locked"`
	Disabled        bool                            `json:"disabled"`
	TargetType      ActionTargetType                `json:"target_type"`
	TargetPredicate func(Game, Actor, Context) bool `json:"-"`
	ContextValidate func(Context) bool              `json:"-"`
	MapContext      func(Game, Context) Context     `json:"-"`
	Cost            GameMutation                    `json:"-"`
	Cooldown        *int                            `json:"cooldown"`
}

func ResolveAction(game *Game, transaction Transaction[Action]) []GameTransaction {
	action := transaction.Mutation

	if action.Disabled || !action.Filter(*game, *game, transaction.Context) {
		context := transaction.Context
		context.ActionID = &action.ID
		logStart := NewLogContext("$source$ used $action$", context)
		logFail := NewLogContext("$action$ failed.", context)
		if action.Config.LogFailureF != nil {
			logFail = NewLog(fmt.Sprintf(*action.Config.LogFailureF, action.Config.Name))
		}

		game.PushLog(logStart)
		game.PushLog(logFail)
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
		log := NewLogContext("$source$ used $action$", context)
		if action.Config.LogSuccessF != nil {
			log = NewLog(fmt.Sprintf(*action.Config.LogSuccessF, source.Name, action.Config.Name))
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

	return action.Delta(*game, *game, context)
}

func GetAccuracy(game Game, source ResolvedActor, target ResolvedActor) float64 {
	ratio := float64(source.Stats[StatAccuracy]) / float64(target.Stats[StatEvasion])
	return ratio
}

func MakeActionRoll() int {
	return rand.Intn(100)
}

type ChanceResult struct {
	Chance  int
	Roll    int
	Success bool
}

func MakeCriticalCheck(action ActionConfig) ChanceResult {
	if action.CritChance == nil {
		return ChanceResult{
			Success: false,
		}
	}

	accuracy := int(*action.CritChance)
	roll := MakeActionRoll()
	return ChanceResult{
		Chance:  accuracy,
		Roll:    roll,
		Success: roll <= accuracy,
	}
}
func GetCriticalModifier(action ActionConfig) float64 {
	result := MakeCriticalCheck(action)
	if result.Success {
		return action.CritMod
	}

	return 1.0
}

func MakeAccuracyCheck(g *Game, action ActionConfig, source ResolvedActor, target ResolvedActor) ChanceResult {
	base_accuracy := GetAccuracy(*g, source, target)
	if action.Accuracy == nil {
		return ChanceResult{
			Success: true,
		}
	}

	accuracy := Round(base_accuracy * float64(*action.Accuracy))
	roll := MakeActionRoll()
	return ChanceResult{
		Chance:  accuracy,
		Roll:    roll,
		Success: roll <= accuracy,
	}
}

func GetActiveActionConfig(g Game, fallback ActionConfig) ActionConfig {
	if g.ActiveTransaction == nil {
		return fallback
	}

	return g.ActiveTransaction.Mutation.Config
}
