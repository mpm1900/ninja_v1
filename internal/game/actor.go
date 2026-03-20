package game

import (
	"maps"
	"slices"
	"sort"

	"github.com/google/uuid"
)

type Resource string

const (
	ResourceHP      Resource = "hp"
	ResourceStamina Resource = "stamina"
)

type BaseStat string

const (
	StatHP       BaseStat = "hp"
	StatStamina  BaseStat = "stamina"
	StatNinjutsu BaseStat = "ninjutsu"
	StatGenjutsu BaseStat = "genjutsu"
	StatTaijutsu BaseStat = "taijutsu"
	StatSpeed    BaseStat = "speed"
	StatEvasion  BaseStat = "evasion"
	StatAccuracy BaseStat = "accuracy"
)

type ResourceValues struct {
	Max    int
	Offset int
}

type Actor struct {
	ID       uuid.UUID `json:"ID"`
	ActorID  uuid.UUID `json:"actor_ID"`
	PlayerID uuid.UUID `json:"player_ID"`
	Name     string    `json:"name"`

	Level       int `json:"level"`
	Experience  int `json:"experience"`
	ActionCount int `json:"action_count"`

	Alive  bool `json:"alive"`
	Active bool `json:"active"`

	// Resources        map[Resource]ResourceValues
	Stats            map[BaseStat]int   `json:"stats"`
	Stages           map[BaseStat]int   `json:"staged_stats"`
	NatureDamage     map[Nature]float64 `json:"nature_damage"`
	NatureResistance map[Nature]float64 `json:"nature_resistance"`
	Critical         float64            `json:"critical"`

	Natures         []NatureSet    `json:"naures"`
	InnateModifiers []Modifier     `json:"innate_modifiers"`
	Actions         []Action[Game] `json:"actions"`
}

type ResolvedActor struct {
	Actor
	BaseStats        map[BaseStat]int `json:"base_stats"`
	PreStats         map[BaseStat]int `json:"pre_stats"`
	AppliedModifiers []uuid.UUID      `json:"applied_modifiers"`
}

type ActorMutation = Mutation[Actor, Actor, Context]

const (
	PriorityPreBaseStats    = -11
	PriorityMapBaseStats    = -10
	PriorityPostBaseStats   = -9
	PriorityDefault         = 0
	PriorityPreStagedStats  = 9
	PriorityMapStagedStats  = 10
	PriorityPostStagedStats = 11
	PriorityZero            = 19
	PrioritySet             = 20
)

func resolve(actor Actor, pre Actor) ResolvedActor {
	return ResolvedActor{
		Actor:            actor,
		BaseStats:        pre.Stats,
		AppliedModifiers: []uuid.UUID{},
	}
}

func MapBaseStat(stat, level int) int {
	base := (stat + 15) * 2
	ev := 0 // TODO
	ratio := ((base + ev) * level) / 100
	return ratio + 5
}

func MapResource(stat, level int) int {
	return MapBaseStat(stat, level) + level + 5
}

func MapBaseStats(actor Actor) Actor {
	actor.Stats[StatHP] = MapResource(actor.Stats[StatHP], actor.Level)
	actor.Stats[StatStamina] = MapResource(actor.Stats[StatStamina], actor.Level)

	actor.Stats[StatNinjutsu] = MapBaseStat(actor.Stats[StatNinjutsu], actor.Level)
	actor.Stats[StatGenjutsu] = MapBaseStat(actor.Stats[StatGenjutsu], actor.Level)
	actor.Stats[StatTaijutsu] = MapBaseStat(actor.Stats[StatTaijutsu], actor.Level)
	actor.Stats[StatSpeed] = MapBaseStat(actor.Stats[StatSpeed], actor.Level)

	return actor
}

func MapStagedStat(stat, stage, mod int) int {
	if stage > 0 {
		return stat * ((stage + mod) / mod)
	}
	if stage < 0 {
		return stat * (mod / (-1*stage + mod))
	}
	return stat
}

func MapStagedStats(actor Actor) Actor {
	actor.Stats[StatNinjutsu] = MapStagedStat(actor.Stats[StatNinjutsu], actor.Stages[StatNinjutsu], 2)
	actor.Stats[StatGenjutsu] = MapStagedStat(actor.Stats[StatGenjutsu], actor.Stages[StatGenjutsu], 2)
	actor.Stats[StatTaijutsu] = MapStagedStat(actor.Stats[StatTaijutsu], actor.Stages[StatTaijutsu], 2)
	actor.Stats[StatSpeed] = MapStagedStat(actor.Stats[StatSpeed], actor.Stages[StatSpeed], 2)
	actor.Stats[StatEvasion] = MapStagedStat(actor.Stats[StatEvasion], actor.Stages[StatEvasion], 3)
	actor.Stats[StatAccuracy] = MapStagedStat(actor.Stats[StatAccuracy], actor.Stages[StatAccuracy], 3)
	return actor
}

func GetActors(game Game, predicate func(Actor) bool) []Actor {
	var actors []Actor
	for _, actor := range game.Actors {
		if predicate(actor) {
			actors = append(actors, actor)
		}
	}

	return actors
}

func GetActorModifiers(game Game) []ModifierTransaction {
	var modifiers []ModifierTransaction
	activeActors := GetActors(game, func(a Actor) bool {
		return a.Active
	})

	for _, actor := range activeActors {
		context := Context{
			SourcePlayerID:    actor.PlayerID,
			SourceActorID:     actor.ID,
			ParentActorID:     actor.ID,
			TargetActorIDs:    []uuid.UUID{},
			TargetPositionIDs: []uuid.UUID{},
		}
		for _, modifier := range actor.InnateModifiers {
			transaction := ModifierTransaction{
				ID:       uuid.New(),
				Context:  &context,
				Mutation: modifier,
			}
			modifiers = append(modifiers, transaction)
		}
	}

	return modifiers
}

var SPECIAL_MUTATIONS []ModifierMutation = []ModifierMutation{
	{
		ModifierID: nil,
		ActorMutation: ActorMutation{
			Delta: func(input Actor, context *Context) Actor {
				return MapBaseStats(input)
			},
			Priority: PriorityMapBaseStats,
		},
	},
	{
		ModifierID: nil,
		ActorMutation: ActorMutation{
			Delta: func(input Actor, context *Context) Actor {
				return MapStagedStats(input)
			},
			Priority: PriorityPreStagedStats,
		},
	},
}

func GetMutations(transactions []ModifierTransaction) []ModifierMutation {
	var mutations []ModifierMutation
	for _, transaction := range transactions {
		mutations = append(mutations, transaction.Mutation.Mutations...)
	}

	return mutations
}

func cloneActor(a Actor) Actor {
	b := a

	// b.Resources = maps.Clone(a.Resources)
	b.Stats = maps.Clone(a.Stats)
	b.Stages = maps.Clone(a.Stages)
	b.NatureDamage = maps.Clone(a.NatureDamage)
	b.NatureResistance = maps.Clone(a.NatureResistance)

	b.Natures = slices.Clone(a.Natures)
	b.InnateModifiers = slices.Clone(a.InnateModifiers)
	b.Actions = slices.Clone(a.Actions)

	return b
}

func resolveActor(actor Actor, mtransactions []ModifierTransaction, atransactions []ModifierTransaction) ResolvedActor {
	context := Context{
		SourcePlayerID:    actor.PlayerID,
		SourceActorID:     actor.ID,
		ParentActorID:     actor.ID,
		TargetActorIDs:    []uuid.UUID{},
		TargetPositionIDs: []uuid.UUID{},
	}

	applied := make(map[uuid.UUID]struct{})
	transactions := []ModifierTransaction{}
	transactions = append(transactions, atransactions...)
	transactions = append(transactions, mtransactions...)
	mutations := GetMutations(transactions)
	mutations = append(mutations, SPECIAL_MUTATIONS...)
	sort.Slice(mutations, func(i, j int) bool {
		return mutations[j].Priority > mutations[i].Priority
	})

	mapped := cloneActor(actor)
	for _, mutation := range mutations {
		tx := MakeTransaction(&mutation.ActorMutation, &context)
		m, apply := ResolveTransaction(mapped, &tx, mapped)
		if apply {
			if mutation.ModifierID != nil {
				applied[*mutation.ModifierID] = struct{}{}
			}
			mapped = m
		}
	}

	resolved := resolve(mapped, actor)
	for key := range applied {
		resolved.AppliedModifiers = append(resolved.AppliedModifiers, key)
	}
	return resolved
}

func ResolveActor(actor Actor, mtransactions []ModifierTransaction, atransactions []ModifierTransaction) ResolvedActor {
	resolved := resolveActor(actor, mtransactions, atransactions)
	pre := resolveActor(actor, []ModifierTransaction{}, []ModifierTransaction{})
	resolved.PreStats = maps.Clone(pre.Stats)

	return resolved
}
