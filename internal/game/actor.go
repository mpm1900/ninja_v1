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
	ResourceHP     Resource = "hp"
	ResourceChakra Resource = "chakra"
)

type AttackStat string

const (
	AttackNinjutsu AttackStat = "ninjutsu"
	AttackGenjutsu AttackStat = "genjutsu"
	AttackTaijutsu AttackStat = "taijutsu"
)

type BaseStat string

const (
	StatHP       BaseStat = BaseStat(ResourceHP)
	StatChakra   BaseStat = BaseStat(ResourceChakra)
	StatNinjutsu BaseStat = BaseStat(AttackNinjutsu)
	StatGenjutsu BaseStat = BaseStat(AttackGenjutsu)
	StatTaijutsu BaseStat = BaseStat(AttackTaijutsu)
	StatSpeed    BaseStat = "speed"
	StatEvasion  BaseStat = "evasion"
	StatAccuracy BaseStat = "accuracy"
)

type ResourceValues struct {
	Max    int
	Offset int
}

type ActorDef struct {
	ActorID uuid.UUID `json:"actor_ID"`
	Name    string    `json:"name"`

	ActionCount      int                    `json:"action_count"`
	Stats            map[BaseStat]int       `json:"stats"`
	NatureDamage     map[Nature]float64     `json:"nature_damage"`
	NatureResistance map[Nature]float64     `json:"nature_resistance"`
	Natures          map[NatureSet][]Nature `json:"natures"`

	InnateModifiers []Modifier  `json:"innate_modifiers"`
	ActionIDs       []uuid.UUID `json:"action_IDs"`
}

type ActorState struct {
	PositionID *uuid.UUID `json:"position_ID"`
	Alive      bool       `json:"alive"`
	Damage     int        `json:"damage"`
}

type Actor struct {
	ActorDef
	ID         uuid.UUID  `json:"ID"`
	PlayerID   uuid.UUID  `json:"player_ID"`
	Level      int        `json:"level"`
	Experience int        `json:"experience"`
	State      ActorState `json:"state"`

	Stages           map[BaseStat]int       `json:"staged_stats"`
	AttackModifiers  map[AttackStat]float64 `json:"attack_modifiers"`
	DefenseModifiers map[AttackStat]float64 `json:"defense_modifiers"`

	Actions []Action `json:"actions"`
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

func GetLevel(experience int) int {
	return int(math.Floor(math.Cbrt(float64(experience))))
}

func GetBaseExperience(level int) int {
	return int(math.Floor(math.Pow(float64(level), 3)))
}

func GetExperienceToNextLevel(level, exp int) int {
	return GetBaseExperience(level+1) - (GetBaseExperience(level) + exp)
}

func MakeActor(def ActorDef, playerID uuid.UUID, experience int) Actor {
	return Actor{
		ActorDef:   def,
		ID:         uuid.New(),
		PlayerID:   playerID,
		Level:      GetLevel(experience),
		Experience: experience,
		State: ActorState{
			Alive:      true,
			Damage:     0,
			PositionID: nil,
		},
		Stages: map[BaseStat]int{
			StatHP:       0,
			StatChakra:   0,
			StatNinjutsu: 0,
			StatGenjutsu: 0,
			StatTaijutsu: 0,
			StatSpeed:    0,
			StatEvasion:  0,
			StatAccuracy: 0,
		},
		AttackModifiers: map[AttackStat]float64{
			AttackGenjutsu: 1,
			AttackNinjutsu: 1,
			AttackTaijutsu: 1,
		},
		DefenseModifiers: map[AttackStat]float64{
			AttackGenjutsu: 1,
			AttackNinjutsu: 1,
			AttackTaijutsu: 1,
		},
	}
}

func resolve(actor Actor, pre Actor) ResolvedActor {
	return ResolvedActor{
		Actor:            actor,
		BaseStats:        maps.Clone(pre.Stats),
		AppliedModifiers: map[uuid.UUID]int{},
	}
}

func MapBaseStat(stat, level int, nature float64) int {
	base := float64((stat * 2) + 15)
	ev := 0.0 // TODO
	ratio := float64((base+(ev/4))*float64(level)) / 100
	return int(math.Floor((ratio + 5) * nature))
}

func MapResourceStat(stat, level int) int {
	return MapBaseStat(stat, level, 1.0) + level + 5
}

func (actor *Actor) MapBase(stat BaseStat) {
	actor.Stats[stat] = MapBaseStat(actor.Stats[stat], actor.Level, 1.0)
}

func (actor *Actor) MapResource(stat BaseStat) {
	actor.Stats[stat] = MapResourceStat(actor.Stats[stat], actor.Level)
}

func MapBaseStats(actor Actor) Actor {
	actor.MapResource(StatHP)
	actor.MapResource(StatChakra)

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
	actor.Stats[stat] = MapStagedStat(actor.Stats[stat], actor.Stages[stat], mod)

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

func GetActorModifiers(game Game) []Transaction[Modifier] {
	var modifiers []Transaction[Modifier]
	activeActors := game.GetActorsFilters(
		ActiveFilter,
	)

	for _, actor := range activeActors {
		context := Context{
			SourcePlayerID:    &actor.PlayerID,
			SourceActorID:     &actor.ID,
			ParentActorID:     &actor.ID,
			TargetActorIDs:    []uuid.UUID{},
			TargetPositionIDs: []uuid.UUID{},
		}
		for _, modifier := range actor.InnateModifiers {
			transaction := MakeModifierTransaction(modifier, context)
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
		func(input Actor, context Context) Actor {
			return MapBaseStats(input)
		},
	),
	MakeModifierMutation(
		nil,
		PriorityMapStagedStats,
		AllFilter,
		func(input Actor, context Context) Actor {
			return MapStagedStats(input)
		},
	),
}

func GetMutationsFromTransactions(transactions []Transaction[Modifier]) []ModifierMutation {
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

	b.Stats = maps.Clone(a.Stats)
	b.Stages = maps.Clone(a.Stages)
	b.NatureDamage = maps.Clone(a.NatureDamage)
	b.NatureResistance = maps.Clone(a.NatureResistance)

	b.Natures = maps.Clone(a.Natures)
	b.InnateModifiers = slices.Clone(a.InnateModifiers)
	b.Actions = slices.Clone(a.Actions)

	return b
}

func resolveActor(actor Actor, mtransactions []Transaction[Modifier], atransactions []Transaction[Modifier]) ResolvedActor {
	applied := make(map[uuid.UUID]int)
	transactions := []Transaction[Modifier]{}
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
				if transaction.ID == *mutation.TransactionID {
					context = transaction.Context
					break
				}
			}
		}

		tx := MakeTransaction(mutation.Mutation, context)
		a, apply := ResolveTransaction(mapped, tx, mapped)
		if apply {
			if mutation.ModifierGroupID != nil {
				if count, ok := applied[*mutation.ModifierGroupID]; ok {
					applied[*mutation.ModifierGroupID] = count + 1
				} else {
					applied[*mutation.ModifierGroupID] = 1
				}
			}
			mapped = a
		}
	}

	resolved := resolve(mapped, actor)
	maps.Copy(resolved.AppliedModifiers, applied)
	return resolved
}

func (a Actor) Resolve(game Game) ResolvedActor {
	resolved := resolveActor(a, game.Modifiers, GetActorModifiers(game))
	pre := resolveActor(a, []Transaction[Modifier]{}, []Transaction[Modifier]{})
	resolved.PreStats = maps.Clone(pre.Stats)

	return resolved
}
