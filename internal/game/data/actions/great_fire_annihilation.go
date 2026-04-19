package actions

import (
	"fmt"
	"math/rand"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var GreatFireAnnihilation = MakeGreatFireAnnihilation()

func MakeGreatFireAnnihilation() game.Action {
	ID := uuid.MustParse("d97ee3bb-7afa-47de-9f8d-2ee77ba6dfe6")

	config := game.ActionConfig{
		Name:        "Great Fire Annihilation",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(ID, config, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		targets := g.GetTargets(context)
		for _, target := range targets {
			// on 30% chance
			roll := rand.Intn(100)
			if roll > 30 {
				continue
			}
			fmt.Println("BURN! roll=", roll)
			mut_ctx := game.Context{
				SourcePlayerID: &target.PlayerID,
				SourceActorID:  &target.ID,
				ParentActorID:  nil, // do not remove on switch
				TargetActorIDs: []uuid.UUID{target.ID},
			}

			mod := mutations.AddStatus(true, modifiers.Burned)
			mod_tx := game.MakeTransaction(mod, mut_ctx)
			mut := mutations.Burn
			mut_tx := game.MakeTransaction(mut, mut_ctx)
			transactions = append(transactions, mod_tx, mut_tx)
		}
		return transactions
	})
}
