package game

import (
	"maps"
	"math"
	"slices"
	"sort"

	"github.com/google/uuid"
)

type AttackStat string

const (
	Attack       AttackStat = "attack"
	ChakraAttack AttackStat = "chakra_attack"
)

type DefenseStat string

const (
	Defense       DefenseStat = "defense"
	ChakraDefense DefenseStat = "chakra_defense"
)

type ActorStat string

const (
	StatHP            ActorStat = "hp"
	StatStamina       ActorStat = "stamina"
	StatAttack        ActorStat = ActorStat(Attack)
	StatDefense       ActorStat = ActorStat(Defense)
	StatChakraAttack  ActorStat = ActorStat(ChakraAttack)
	StatChakraDefense ActorStat = ActorStat(ChakraDefense)
	StatSpeed         ActorStat = "speed"
	StatEvasion       ActorStat = "evasion"
	StatAccuracy      ActorStat = "accuracy"
)

type ActorFocus string

const (
	FocusNone ActorFocus = "none"
)

var ActorFocuses map[ActorFocus][]*ActorStat = map[ActorFocus][]*ActorStat{
	FocusNone: {nil, nil},
}

const (
	AffAme      = "ame"
	AffAkatsuki = "akatsuki"
	AffIwa      = "iwa"
	AffKonoha   = "konoha"
	AffKumo     = "kumo"
	AffKuri     = "kuri"
	AffOto      = "oto"
	AffTaki     = "taki"
	AffYuga     = "yuga"

	ClanHatake  = "hatake"
	ClanSenju   = "senju"
	ClanUchiha  = "uchiha"
	ClanUzumaki = "uzumaki"
)

type ResourceValues struct {
	Max    int
	Offset int
}

type ActorDef struct {
	ActorID      uuid.UUID `json:"actor_ID"`
	SpriteURL    string    `json:"sprite_url"`
	Name         string    `json:"name"`
	Clan         string    `json:"clan"`
	Affiliations []string  `json:"affiliations"`

	Stats            map[ActorStat]int      `json:"stats"`
	NatureDamage     map[Nature]float64     `json:"nature_damage"`
	NatureResistance map[Nature]float64     `json:"nature_resistance"`
	Natures          map[NatureSet][]Nature `json:"natures"`

	InnateModifiers []Modifier  `json:"innate_modifiers"`
	ActionIDs       []uuid.UUID `json:"action_IDs"`
	ActionCount     int         `json:"action_count"`
}

type ActorState struct {
	// [Alive] whether or not the actor is alive, could
	// - could be computed, but this is here to not have to call .Resolve() on filters
	Alive bool `json:"alive"`
	// [Damage] how much damage this actor has recieved
	Damage        int `json:"damage"`
	StaminaDamage int `json:"stamina_damage"`
	// [PositionID] current position, nil if not active
	PositionID *uuid.UUID `json:"position_ID"`
	// [Reflect] how much damage is reflected (PureDamage not affected)
	Reflect float64 `json:"reflect"`
	Stunned bool    `json:"stunned"`
}

type Actor struct {
	ActorDef
	ActorState
	ID         uuid.UUID  `json:"ID"`
	PlayerID   uuid.UUID  `json:"player_ID"`
	Level      int        `json:"level"`
	Experience int        `json:"experience"`
	Focus      ActorFocus `json:"focus"`

	Stages map[ActorStat]int `json:"staged_stats"`

	Actions         []Action          `json:"actions"`
	ActionCooldowns map[uuid.UUID]int `json:"action_cooldowns"`
}

type ResolvedActor struct {
	Actor
	BaseStats        map[ActorStat]int `json:"base_stats"`
	PreStats         map[ActorStat]int `json:"pre_stats"`
	AppliedModifiers map[uuid.UUID]int `json:"applied_modifiers"`
}

const (
	MutPriorityStateChanges    = -20
	MutPriorityPreBaseStats    = -11
	MutPriorityMapBaseStats    = -10
	MutPriorityPostBaseStats   = -9
	MutPriorityDefault         = 0
	MutPriorityPreStagedStats  = 9
	MutPriorityMapStagedStats  = 10
	MutPriorityPostStagedStats = 11
	MutPriorityZero            = 19
	MutPrioritySet             = 20
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

// TODO make actions selectable
func MakeActor(def ActorDef, playerID uuid.UUID, experience int, ACTIONS map[uuid.UUID]Action) Actor {
	actions := make([]Action, 0)
	for _, id := range def.ActionIDs {
		a, ok := ACTIONS[id]
		if !ok {
			continue
		}
		actions = append(actions, a)
	}
	return Actor{
		ActorDef:   def,
		ID:         uuid.New(),
		PlayerID:   playerID,
		Level:      GetLevel(experience),
		Experience: experience,
		Focus:      FocusNone,
		ActorState: ActorState{
			Alive:      true,
			Damage:     0,
			PositionID: nil,
			Reflect:    0.0,
		},
		Stages: map[ActorStat]int{
			StatHP:            0,
			StatStamina:       0,
			StatAttack:        0,
			StatDefense:       0,
			StatChakraAttack:  0,
			StatChakraDefense: 0,
			StatSpeed:         0,
			StatEvasion:       0,
			StatAccuracy:      0,
		},
		Actions:         actions,
		ActionCooldowns: map[uuid.UUID]int{},
	}
}

func (a *Actor) SetActionCooldown(actionID uuid.UUID, cooldown int) {
	a.ActionCooldowns[actionID] = cooldown
}
func (a *Actor) DecrementCooldowns() {
	for actionID, cooldown := range a.ActionCooldowns {
		if cooldown <= 0 {
			delete(a.ActionCooldowns, actionID)
			continue
		}
		a.ActionCooldowns[actionID] = cooldown - 1
	}
}
func (a *Actor) RecoverStamina(g Game, ratio float64) {
	resolved := a.Resolve(g)
	amount := int(math.Floor(float64(resolved.Stats[StatStamina]) * ratio))
	a.StaminaDamage = max(a.StaminaDamage-amount, 0)
}

func (a Actor) GetFocusModifier(stat ActorStat) float64 {
	stats, ok := ActorFocuses[a.Focus]
	if !ok || len(stats) != 2 {
		return 0.0
	}

	if stats[0] != nil && *stats[0] == stat {
		return 1.1
	}
	if stats[1] != nil && *stats[1] == stat {
		return 0.9
	}

	return 1.0
}

func resolve(actor Actor, pre Actor) ResolvedActor {
	return ResolvedActor{
		Actor:            actor,
		BaseStats:        maps.Clone(pre.Stats),
		AppliedModifiers: map[uuid.UUID]int{},
	}
}

func MapBaseStat(stat, level int, focus float64) int {
	base := float64((stat * 2) + 15)
	ev := 0.0 // TODO
	ratio := float64((base+(ev/4))*float64(level)) / 100
	return int(math.Floor((ratio + 5) * focus))
}

func MapResourceStat(stat, level int, focus float64) int {
	return MapBaseStat(stat, level, focus) + level + 5
}

func (actor *Actor) MapBase(stat ActorStat) {
	actor.Stats[stat] = MapBaseStat(actor.Stats[stat], actor.Level, actor.GetFocusModifier(stat))
}

func (actor *Actor) MapResource(stat ActorStat) {
	actor.Stats[stat] = MapResourceStat(actor.Stats[stat], actor.Level, 1.0)
}

func MapBaseStats(actor Actor) Actor {
	actor.MapResource(StatHP)
	actor.MapResource(StatStamina)

	actor.MapBase(StatAttack)
	actor.MapBase(StatDefense)
	actor.MapBase(StatChakraAttack)
	actor.MapBase(StatChakraDefense)
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

func (actor *Actor) MapStaged(stat ActorStat, mod int) {
	actor.Stats[stat] = MapStagedStat(actor.Stats[stat], actor.Stages[stat], mod)

}

func MapStagedStats(actor Actor) Actor {
	actor.MapStaged(StatAttack, 2)
	actor.MapStaged(StatDefense, 2)
	actor.MapStaged(StatChakraAttack, 2)
	actor.MapStaged(StatChakraDefense, 2)
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

var SPECIAL_MUTATIONS []ActorMutation = []ActorMutation{
	MakeActorMutation(
		nil,
		MutPriorityMapBaseStats,
		AllFilter,
		func(input Actor, context Context) Actor {
			return MapBaseStats(input)
		},
	),
	MakeActorMutation(
		nil,
		MutPriorityMapStagedStats,
		AllFilter,
		func(input Actor, context Context) Actor {
			return MapStagedStats(input)
		},
	),
}

func GetMutationsFromTransactions(transactions []Transaction[Modifier]) []ActorMutation {
	var mutations []ActorMutation
	for _, transaction := range transactions {
		var muts []ActorMutation
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

func (r ResolvedActor) HasChakra(amount int) bool {
	return (r.Stats[StatStamina] - r.StaminaDamage) >= amount
}
