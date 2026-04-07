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
	// +P.ATK
	FocusAggressive ActorFocus = "aggressive" // -P.DEF
	FocusRelentless ActorFocus = "relentless" // -C.ATK
	FocusReckless   ActorFocus = "reckless"   // -C.DEF
	FocusHeavy      ActorFocus = "heavy"      // -SPE
	// +P.DEF
	FocusPatient   ActorFocus = "patient"   // -P.ATK
	FocusHardened  ActorFocus = "hardened"  // -C.ATK
	FocusTough     ActorFocus = "tough"     // -C.DEF
	FocusSteadfast ActorFocus = "steadfast" // -SPE
	// +C.ATK
	FocusIntelligent ActorFocus = "intelligent" // -P.ATK
	FocusVolatile    ActorFocus = "volatile"    // -P.DEF
	FocusIntense     ActorFocus = "intense"     // -C.DEF
	FocusCalculated  ActorFocus = "calculated"  // -SPE
	// +SPE
	FocusAgile     ActorFocus = "agile"     // -P.ATK
	FocusHasty     ActorFocus = "hasty"     // -P.DEF
	FocusImpulsive ActorFocus = "impulsive" // -C.ATK
	FocusAlert     ActorFocus = "alert"     // -C.DEF
)

type FocusStatDelta struct {
	Up   ActorStat
	Down ActorStat
}

var ActorFocuses = map[ActorFocus]FocusStatDelta{
	FocusNone:        {},
	FocusAggressive:  {Up: StatAttack, Down: StatDefense},
	FocusRelentless:  {Up: StatAttack, Down: StatChakraAttack},
	FocusReckless:    {Up: StatAttack, Down: StatChakraDefense},
	FocusHeavy:       {Up: StatAttack, Down: StatSpeed},
	FocusPatient:     {Up: StatDefense, Down: StatAttack},
	FocusHardened:    {Up: StatDefense, Down: StatChakraAttack},
	FocusTough:       {Up: StatDefense, Down: StatChakraDefense},
	FocusSteadfast:   {Up: StatDefense, Down: StatSpeed},
	FocusIntelligent: {Up: StatChakraAttack, Down: StatAttack},
	FocusVolatile:    {Up: StatChakraAttack, Down: StatDefense},
	FocusIntense:     {Up: StatChakraAttack, Down: StatChakraDefense},
	FocusCalculated:  {Up: StatChakraAttack, Down: StatSpeed},
	FocusAgile:       {Up: StatSpeed, Down: StatAttack},
	FocusHasty:       {Up: StatSpeed, Down: StatDefense},
	FocusImpulsive:   {Up: StatSpeed, Down: StatChakraAttack},
	FocusAlert:       {Up: StatSpeed, Down: StatChakraDefense},
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

/**
 * [ActorDef]
 * Actor Definitions, the non-runtime definition of a shinobi's
 * stats, natures, and information
 */
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
	ActionIDs       []uuid.UUID `json:"action_IDs,omitempty"`
	ActionCount     int         `json:"action_count"`
}

/**
 * [ActorState]
 * - the commonly modified fields
 */
type ActorState struct {
	ActiveTurns   int `json:"active_turns"`
	InactiveTurns int `json:"inactive_turns"`
	// [Alive] whether or not the actor is alive, could
	// - could be computed, but this is here to not have to call .Resolve() on filters
	Alive bool `json:"alive"`
	// [Damage] how much damage this actor has recieved
	Damage int `json:"damage"`
	// [PositionID] current position, nil if not active
	PositionID *uuid.UUID `json:"position_ID"`
	// [Protected]
	// - protected units cannot be damaged by actions
	// - protected units cannot be targeted by enemy actions
	Protected bool `json:"protected"`
	// [Reflect] how much damage is reflected (PureDamage not affected)
	Reflect       float64 `json:"reflect"`
	Seen          bool    `json:"seen"`
	StaminaDamage int     `json:"stamina_damage"`
	// [Stunned] whether or not an actor _can act_
	// - stunned units cannot push actions
	// - stunned units cannot resolve actions (if the status was added during running)
	Stunned bool `json:"stunned"`
	/**
	 * Statuses
	 * Made the choice for the core to not reference status modifiers by name,
	 *  rather, have both reference keys in ActorState below. For example,
	 *  ResolveAction will no check if the source has a modifier "Paralysis,"
	 *  It instead check the flag set by Paralysis here.
	 */
	Statused  bool `json:"statused"`
	Paralyzed bool `json:"paralyzed"`
}

type Summon struct {
	Actor
	Parent *Actor `json:"-"`
	Proxy  bool   `json:"proxy"`
}

type Actor struct {
	ActorDef
	ActorState
	ID         uuid.UUID         `json:"ID"`
	PlayerID   uuid.UUID         `json:"player_ID"`
	Level      int               `json:"level"`
	Experience int               `json:"experience"`
	Focus      ActorFocus        `json:"focus"`
	Stages     map[ActorStat]int `json:"staged_stats"`
	Actions    []Action          `json:"actions"`
	Immunities []uuid.UUID       `json:"immunities"`
	Summon     *Summon           `json:"summon,omitempty"`
}

type ResolvedActor struct {
	Actor
	BaseStats                map[ActorStat]int  `json:"base_stats"`
	PreStats                 map[ActorStat]int  `json:"pre_stats"`
	AppliedModifiers         map[uuid.UUID]int  `json:"applied_modifiers"`
	ResolvedNatureResistance map[Nature]float64 `json:"resolved_nature_resistance"`
	ResolvedNatureDamage     map[Nature]float64 `json:"resolved_nature_damage"`
}

const (
	MutPriorityImmunity        = -30
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
	MutPriorityPostSet         = 21 // toad song
)

func GetLevel(experience int) int {
	return Round(math.Cbrt(float64(experience)))
}

func GetBaseExperience(level int) int {
	return Round(math.Pow(float64(level), 3))
}

func GetExperienceToNextLevel(level, exp int) int {
	return GetBaseExperience(level+1) - (GetBaseExperience(level) + exp)
}

func (a Actor) GetActionByID(g Game, actionID uuid.UUID) (Action, bool) {
	ra := a.Resolve(g)
	return ra.GetActionByID(actionID)
}
func (a Actor) IsActive() bool {
	return a.PositionID != nil
}
func (ra ResolvedActor) GetActionByID(actionID uuid.UUID) (Action, bool) {
	for _, action := range ra.Actions {
		if action.ID == actionID {
			return action, true
		}
	}

	return Action{}, false
}

func makeActions(actionIDs []uuid.UUID, ACTIONS map[uuid.UUID]Action) []Action {
	actions := make([]Action, 0, len(actionIDs))
	for _, id := range actionIDs {
		a, ok := ACTIONS[id]
		if !ok {
			continue
		}
		actions = append(actions, a)
	}

	return actions
}

func cloneActorDef(def ActorDef) ActorDef {
	cloned := def

	cloned.Affiliations = slices.Clone(def.Affiliations)
	cloned.Stats = maps.Clone(def.Stats)
	cloned.NatureDamage = maps.Clone(def.NatureDamage)
	cloned.NatureResistance = maps.Clone(def.NatureResistance)

	cloned.Natures = make(map[NatureSet][]Nature, len(def.Natures))
	for k, v := range def.Natures {
		cloned.Natures[k] = slices.Clone(v)
	}

	cloned.InnateModifiers = slices.Clone(def.InnateModifiers)
	cloned.ActionIDs = slices.Clone(def.ActionIDs)

	return cloned
}

func MakeActor(def ActorDef, playerID uuid.UUID, experience int, actionIDs []uuid.UUID, ACTIONS map[uuid.UUID]Action) Actor {
	actions := makeActions(actionIDs, ACTIONS)
	clonedDef := cloneActorDef(def)

	return Actor{
		ActorDef:   clonedDef,
		ID:         uuid.New(),
		PlayerID:   playerID,
		Level:      GetLevel(experience),
		Experience: experience,
		Focus:      FocusNone,
		Immunities: []uuid.UUID{},
		ActorState: ActorState{
			ActiveTurns:   0,
			Alive:         true,
			Damage:        0,
			InactiveTurns: 0,
			PositionID:    nil,
			Protected:     false,
			Reflect:       0.0,
			Seen:          false,
			StaminaDamage: 0,
			Stunned:       false,
			Statused:      false,
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
		Actions: actions,
		Summon:  nil,
	}
}

func (a *Actor) SetActions(actionIDs []uuid.UUID, ACTIONS map[uuid.UUID]Action) {
	a.Actions = makeActions(actionIDs, ACTIONS)
}
func (a *Actor) SetPosition(positionID *uuid.UUID) {
	a.PositionID = positionID
	if positionID != nil {
		a.ActiveTurns = 0
		a.Seen = true
	} else {
		a.InactiveTurns = 0
		a.SetSummon(nil)
	}
}
func (a *Actor) PushImmunities(ids ...uuid.UUID) {
	a.Immunities = append(a.Immunities, ids...)
}
func (a Actor) HasImmunity(id uuid.UUID) bool {
	return slices.Contains(a.Immunities, id)
}
func (a *Actor) SetSummon(summon *Summon) {
	if summon != nil {
		summon.Parent = a
	}
	a.Summon = summon
}
func (a *Actor) SetSummonFromActor(actor *Actor, proxy bool) {
	if actor == nil {
		a.Summon = nil
		return
	}

	summon := Summon{
		Actor:  *actor,
		Parent: a,
		Proxy:  proxy,
	}
	a.SetSummon(&summon)
}
func (a *Actor) SetActionCooldown(actionID uuid.UUID, cooldown int) {
	for i := range a.Actions {
		if a.Actions[i].ID == actionID {
			a.Actions[i].Cooldown = &cooldown
		}
	}
}
func (a *Actor) DecrementCooldowns() {
	for i := range a.Actions {
		if a.Actions[i].Cooldown != nil {
			cd := *a.Actions[i].Cooldown - 1
			if cd < 0 {
				a.Actions[i].Cooldown = nil
			} else {
				a.Actions[i].Cooldown = &cd
			}
		}
	}
}
func (a *Actor) RecoverStamina(g Game, ratio float64) {
	resolved := a.Resolve(g)
	amount := Round(float64(resolved.Stats[StatStamina]) * ratio)
	a.StaminaDamage = max(a.StaminaDamage-amount, 0)
}

func (a *Actor) IncrementTurns() {
	if a.IsActive() {
		a.ActiveTurns++
	} else {
		a.InactiveTurns++
	}
}

func (a Actor) GetFocusModifier(stat ActorStat) float64 {
	delta, ok := ActorFocuses[a.Focus]
	if !ok {
		return 1.0
	}

	if delta.Up == stat {
		return 1.1
	}
	if delta.Down == stat {
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
	return Round((ratio + 5) * focus)
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
		stage = min(stage, 6)
		m = float64(stage+mod) / float64(mod)
	} else if stage < 0 {
		stage = max(stage, -6)
		m = float64(mod) / float64(-stage+mod)
	}

	return Round(float64(stat) * m)
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

func newActorContext(actor Actor) Context {
	return Context{
		SourcePlayerID:    &actor.PlayerID,
		SourceActorID:     &actor.ID,
		ParentActorID:     &actor.ID,
		TargetActorIDs:    []uuid.UUID{},
		TargetPositionIDs: []uuid.UUID{},
	}
}

func GetActorModifiers(game Game) []Transaction[Modifier] {
	var modifiers []Transaction[Modifier]
	activeActors := game.GetActorsFilters(
		Context{},
		ActiveFilter,
	)

	for _, actor := range activeActors {
		context := newActorContext(actor)
		for _, modifier := range actor.InnateModifiers {
			transaction := MakeTransaction(modifier, context)
			modifiers = append(modifiers, transaction)
		}
	}

	return modifiers
}

var specialMutations = []ActorMutation{
	MakeActorMutation(
		nil,
		MutPriorityMapBaseStats,
		AllFilter,
		func(g Game, input Actor, context Context) Actor {
			return MapBaseStats(input)
		},
	),
	MakeActorMutation(
		nil,
		MutPriorityMapStagedStats,
		AllFilter,
		func(g Game, input Actor, context Context) Actor {
			return MapStagedStats(input)
		},
	),
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

func getContext(actor Actor, transactions []Transaction[Modifier], mutation ActorMutation) Context {
	context := newActorContext(actor)

	if mutation.TransactionID == nil {
		return context
	}

	for _, transaction := range transactions {
		if transaction.ID == *mutation.TransactionID {
			return transaction.Context
		}
	}

	return context
}

func getMutations(transactions []Transaction[Modifier]) []ActorMutation {
	mutations := make([]ActorMutation, 0)
	for _, transaction := range transactions {
		for _, mut := range transaction.Mutation.ActorMutations {
			mut.TransactionID = &transaction.ID
			mutations = append(mutations, mut)
		}
	}

	return append(mutations, specialMutations...)
}

func applyModifierMutation(gi Game, actor Actor, transactions []Transaction[Modifier], mutation ActorMutation) (Actor, bool) {
	if mutation.ModifierGroupID != nil && actor.HasImmunity(*mutation.ModifierGroupID) {
		return actor, false
	}
	context := getContext(actor, transactions, mutation)
	g := gi.WithActor(actor)

	tx := MakeTransaction(mutation.Mutation, context)
	next, ok := ResolveTransaction(g, actor, tx, actor)
	if !ok {
		return actor, false
	}

	return next, true
}

func resolveActor(actor Actor, g Game, mtransactions []Transaction[Modifier], atransactions []Transaction[Modifier]) ResolvedActor {
	applied := make(map[uuid.UUID]int)
	transactions := []Transaction[Modifier]{}
	transactions = append(transactions, atransactions...)
	transactions = append(transactions, mtransactions...)

	mutations := getMutations(transactions)

	sort.SliceStable(mutations, func(i, j int) bool {
		return mutations[i].Priority < mutations[j].Priority
	})

	mapped := actor.Clone()
	for _, mutation := range mutations {
		next, apply := applyModifierMutation(g, mapped, transactions, mutation)
		if !apply {
			continue
		}

		mapped = next
		if mutation.ModifierGroupID != nil {
			if count, ok := applied[*mutation.ModifierGroupID]; ok {
				applied[*mutation.ModifierGroupID] = count + 1
			} else {
				applied[*mutation.ModifierGroupID] = 1
			}
		}
	}

	resolved := resolve(mapped, actor)
	maps.Copy(resolved.AppliedModifiers, applied)
	resolved.ResolvedNatureResistance = make(map[Nature]float64)
	resolved.ResolvedNatureDamage = make(map[Nature]float64)

	for nature := range resolved.NatureResistance {
		incomingMultiplier := ResolveNatures(
			[]Nature{nature},
			NewNatureSetValues(),
			resolved.NatureResistance,
			resolved.Natures,
		)

		if incomingMultiplier == 0 {
			resolved.ResolvedNatureResistance[nature] = 0
			continue
		}

		resolved.ResolvedNatureResistance[nature] = 1.0 / incomingMultiplier
		ns := NatureSet(nature)
		resolved.ResolvedNatureDamage[nature] = GetStabModifier(resolved, &ns)
	}

	return resolved
}

func (a Actor) Resolve(g Game) ResolvedActor {
	resolved := resolveActor(a, g, g.GetModifiers(), GetActorModifiers(g))
	pre := resolveActor(a, g, []Transaction[Modifier]{}, []Transaction[Modifier]{})
	resolved.PreStats = maps.Clone(pre.Stats)

	return resolved
}

func (r ResolvedActor) HasChakra(amount int) bool {
	return (r.Stats[StatStamina] - r.StaminaDamage) >= amount
}
func (r ResolvedActor) IsProtected(mut GameMutation) (GameTransaction, bool) {
	if r.Protected {
		return MakeTransaction(mut, NewContext()), true
	}

	return GameTransaction{}, false
}
