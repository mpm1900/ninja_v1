package game

import (
	"maps"
	"math"
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

	Natures         []NatureSet `json:"natures"`
	InnateModifiers []Modifier  `json:"innate_modifiers"`
	Actions         []Action    `json:"actions"`
}

type ResolvedActor struct {
	Actor
	BaseStats        map[BaseStat]int  `json:"base_stats"`
	PreStats         map[BaseStat]int  `json:"pre_stats"`
	AppliedModifiers map[uuid.UUID]int `json:"applied_modifiers"`
}

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
		AppliedModifiers: map[uuid.UUID]int{},
	}
}

func MapBaseStat(stat, level int) int {
	base := (stat + 15) * 2
	ev := 0 // TODO
	ratio := float64((base+ev)*level) / 100
	return int(math.Floor(ratio + 5))
}

func MapResourceStat(stat, level int) int {
	return MapBaseStat(stat, level) + level + 5
}

func (actor *Actor) MapBase(stat BaseStat) {
	actor.Stats[stat] = MapBaseStat(actor.Stats[stat], actor.Level)
}

func (actor *Actor) MapResource(stat BaseStat) {
	actor.Stats[stat] = MapResourceStat(actor.Stats[stat], actor.Level)
}

func MapBaseStats(actor Actor) Actor {
	actor.MapResource(StatHP)
	actor.MapResource(StatStamina)

	actor.MapBase(StatNinjutsu)
	actor.MapBase(StatGenjutsu)
	actor.MapBase(StatTaijutsu)
	actor.MapBase(StatSpeed)

	return actor
}

func MapStagedStat(stat, stage, mod int) int {
	m := 1.0
	if stage > 0 {
		m = float64(stage+mod) / float64(mod)
	} else if stage < 0 {
		m = float64(mod) / float64(-stage+mod)
	}

	return int(math.Floor(float64(stat) * m))
}

func (actor *Actor) MapStaged(stat BaseStat, mod int) {
	actor.Stats[stat] = MapStagedStat(actor.Stats[stat], actor.Stages[stat], 2)
}

func MapStagedStats(actor Actor) Actor {
	actor.MapStaged(StatNinjutsu, 2)
	actor.MapStaged(StatGenjutsu, 2)
	actor.MapStaged(StatTaijutsu, 2)
	actor.MapStaged(StatSpeed, 2)
	actor.MapStaged(StatEvasion, 3)
	actor.MapStaged(StatAccuracy, 3)
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

func GetActorModifiers(game Game) []Transaction[Modifier, Context] {
	var modifiers []Transaction[Modifier, Context]
	activeActors := GetActors(game, func(a Actor) bool {
		return a.Active
	})

	for _, actor := range activeActors {
		context := Context{
			SourcePlayerID:    &actor.PlayerID,
			SourceActorID:     &actor.ID,
			ParentActorID:     &actor.ID,
			TargetActorIDs:    []uuid.UUID{},
			TargetPositionIDs: []uuid.UUID{},
		}
		for _, modifier := range actor.InnateModifiers {
			transaction := MakeModifierTransaction(&modifier, &context)
			modifiers = append(modifiers, transaction)
		}
	}

	return modifiers
}

var SPECIAL_MUTATIONS []ModifierMutation = []ModifierMutation{
	MakeModifierMutation(
		nil,
		PriorityMapBaseStats,
		AllFilter,
		func(input Actor, context *Context) Actor {
			return MapBaseStats(input)
		},
	),
	MakeModifierMutation(
		nil,
		PriorityMapStagedStats,
		AllFilter,
		func(input Actor, context *Context) Actor {
			return MapStagedStats(input)
		},
	),
}

func GetMutationsFromTransactions(transactions []Transaction[Modifier, Context]) []ModifierMutation {
	var mutations []ModifierMutation
	for _, transaction := range transactions {
		var muts []ModifierMutation
		for _, mut := range transaction.Mutation.Mutations {
			mut.TransactionID = &transaction.ID
			muts = append(muts, mut)
		}
		mutations = append(mutations, muts...)
	}

	return mutations
}

func (a Actor) Clone() Actor {
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

func resolveActor(actor Actor, mtransactions []Transaction[Modifier, Context], atransactions []Transaction[Modifier, Context]) ResolvedActor {
	applied := make(map[uuid.UUID]int)
	transactions := []Transaction[Modifier, Context]{}
	transactions = append(transactions, atransactions...)
	transactions = append(transactions, mtransactions...)
	mutations := GetMutationsFromTransactions(transactions)
	mutations = append(mutations, SPECIAL_MUTATIONS...)
	sort.Slice(mutations, func(i, j int) bool {
		return mutations[i].Priority < mutations[j].Priority
	})

	mapped := actor.Clone()
	for _, mutation := range mutations {
		context := Context{
			SourcePlayerID:    &actor.PlayerID,
			SourceActorID:     &actor.ID,
			ParentActorID:     &actor.ID,
			TargetActorIDs:    []uuid.UUID{},
			TargetPositionIDs: []uuid.UUID{},
		}

		if mutation.TransactionID != nil {
			for _, transaction := range transactions {
				if transaction.ID == *mutation.TransactionID && transaction.Context != nil {
					context = *transaction.Context
					break
				}
			}
		}

		tx := MakeTransaction(&mutation.Mutation, &context)
		a, apply := ResolveTransaction(mapped, &tx, mapped)
		if apply {
			if mutation.ModifierID != nil {
				if count, ok := applied[*mutation.ModifierID]; ok {
					applied[*mutation.ModifierID] = count + 1
				} else {
					applied[*mutation.ModifierID] = 0
				}
			}
			mapped = a
		}
	}

	resolved := resolve(mapped, actor)
	maps.Copy(resolved.AppliedModifiers, applied)
	return resolved
}

func ResolveActor(actor Actor, game Game) ResolvedActor {
	resolved := resolveActor(actor, game.Modifiers, GetActorModifiers(game))
	pre := resolveActor(actor, []Transaction[Modifier, Context]{}, []Transaction[Modifier, Context]{})
	resolved.PreStats = maps.Clone(pre.Stats)

	return resolved
}
