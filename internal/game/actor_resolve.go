package game

import (
	"github.com/google/uuid"
)

type actorResolveHandler struct {
	pre              Actor
	actor            Actor
	bypass_modifiers bool
	game             Game
	mutations        []ActorMutation
	transactions     []Transaction[Modifier]
}

func newActorResolveHandler(actor Actor, g Game, bypass_modifiers bool) actorResolveHandler {
	mutations, transactions := GetAllActorMutations(g, bypass_modifiers)
	clone := actor.Clone()
	return actorResolveHandler{
		actor:            clone,
		pre:              clone,
		game:             g,
		bypass_modifiers: bypass_modifiers,
		mutations:        mutations,
		transactions:     transactions,
	}
}

func (ah *actorResolveHandler) applyModifierMutation(mutation ActorMutation) (Actor, bool) {
	if mutation.ModifierGroupID != nil && ah.actor.HasImmunity(*mutation.ModifierGroupID) {
		return ah.actor, false
	}
	context := getContext(ah.actor, ah.transactions, mutation)
	g := ah.game.WithActor(ah.actor)

	tx := MakeTransaction(mutation.Mutation, context)
	next, ok := ResolveTransaction(g, ah.actor, tx, ah.actor)
	if !ok {
		return ah.actor, false
	}

	if mutation.ModifierGroupID == nil {
		return next, true
	}

	if next.AppliedModifiers == nil {
		next.AppliedModifiers = make(map[uuid.UUID]int)
	}

	count, ok := next.AppliedModifiers[*mutation.ModifierGroupID]
	if !ok {
		next.AppliedModifiers[*mutation.ModifierGroupID] = 1
		return next, true
	}

	next.AppliedModifiers[*mutation.ModifierGroupID] = count + 1
	return next, true
}

func (ah *actorResolveHandler) resolveMutations() ResolvedActor {
	for _, mutation := range ah.mutations {
		next, did_apply := ah.applyModifierMutation(mutation)
		if !did_apply {
			continue
		}

		ah.actor = next
	}

	return toResolved(ah.actor, ah.pre)
}

func (ah *actorResolveHandler) resolveNatures(resolved *ResolvedActor) {
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
		resolved.ResolvedNatureDamage[nature] = GetStabModifier(*resolved, &ns)
	}
}

func (ah *actorResolveHandler) resolveActions(resolved *ResolvedActor) {
	for i, _ := range resolved.Actions {
		if resolved.Actions[i].Config.Cooldown == nil {
			resolved.Actions[i].Config.Cooldown = Ptr(resolved.CooldownOffset)
		} else {
			*resolved.Actions[i].Config.Cooldown += resolved.CooldownOffset
		}
	}
}
