package main

import (
	"encoding/json"
	"fmt"

	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func main() {
	// Action has unexported fields, so build zero-values first, then include them.
	var action1 game.Action[game.Game]
	var action2 game.Action[game.Game]

	actor := game.Actor{
		ID:       uuid.New(),
		PlayerID: uuid.New(),
		Name:     "Itachi",

		Level:       50,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       55,
			game.StatStamina:  80,
			game.StatNinjutsu: 120,
			game.StatGenjutsu: 155,
			game.StatTaijutsu: 80,
			game.StatSpeed:    120,
			game.StatEvasion:  0,
			game.StatAccuracy: 1,
		},

		Stages: map[game.BaseStat]int{
			game.StatHP:       0,
			game.StatStamina:  0,
			game.StatNinjutsu: 0,
			game.StatGenjutsu: 0,
			game.StatTaijutsu: 0,
			game.StatSpeed:    0,
			game.StatEvasion:  0,
			game.StatAccuracy: 0,
		},

		Critical: 1.50,
		NatureDamage: map[game.Nature]float64{
			game.NatureFire:      1.30,
			game.NatureWind:      1.00,
			game.NatureLightning: 1.05,
			game.NatureEarth:     0.95,
			game.NatureWater:     0.90,
			game.NatureYin:       1.35,
			game.NatureYang:      0.95,
		},
		NatureResistance: map[game.Nature]float64{
			game.NatureFire:      1.15,
			game.NatureWind:      1.00,
			game.NatureLightning: 1.00,
			game.NatureEarth:     1.00,
			game.NatureWater:     0.95,
			game.NatureYin:       1.25,
			game.NatureYang:      0.90,
		},

		Natures: []game.NatureSet{
			game.NsFire,
			game.NsYin,
		},

		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action[game.Game]{action1, action2},
	}

	var g game.Game
	modifier := game.MakeModifier("Ninjusu doubler")
	modifier.Mutations = []game.ModifierMutation{
		{
			ModifierID: &modifier.ID,
			ActorMutation: game.ActorMutation{
				Filter: func(actor game.Actor, context *game.Context) bool {
					return actor.ID == context.SourceActorID
				},
				Delta: func(actor game.Actor, context *game.Context) game.Actor {
					actor.Stats[game.StatNinjutsu] = actor.Stats[game.StatNinjutsu] * 2
					return actor
				},
			},
		},
	}
	transaction := game.ModifierTransaction{
		ID: uuid.New(),
		Context: &game.Context{
			SourceActorID: actor.ID,
		},
		Mutation: modifier,
	}
	resolved := game.ResolveActor(actor, []game.ModifierTransaction{transaction}, game.GetActorModifiers(g))

	actorJSON, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("=== Input Actor JSON ===")
	fmt.Println(string(actorJSON))

	resolvedJSON, err := json.MarshalIndent(resolved, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("=== Resolved Actor JSON ===")
	fmt.Println(string(resolvedJSON))

	fmt.Println("=== Resolved Actor ===")
	fmt.Printf("ID=%s Name=%s PlayerID=%s\n", resolved.ID, resolved.Name, resolved.PlayerID)
	fmt.Printf("Active=%t Alive=%t Level=%d XP=%d ActionCount=%d\n",
		resolved.Active, resolved.Alive, resolved.Level, resolved.Experience, resolved.ActionCount)
	fmt.Printf("Critical=%.2f\n", resolved.Critical)

	fmt.Printf("Stats: hp=%d stamina=%d ninjutsu=%d genjutsu=%d taijutsu=%d speed=%d evasion=%d accuracy=%d\n",
		resolved.Stats[game.StatHP],
		resolved.Stats[game.StatStamina],
		resolved.Stats[game.StatNinjutsu],
		resolved.Stats[game.StatGenjutsu],
		resolved.Stats[game.StatTaijutsu],
		resolved.Stats[game.StatSpeed],
		resolved.Stats[game.StatEvasion],
		resolved.Stats[game.StatAccuracy],
	)

	fmt.Printf("PreStats: hp=%d stamina=%d ninjutsu=%d genjutsu=%d taijutsu=%d speed=%d evasion=%d accuracy=%d\n",
		resolved.PreStats[game.StatHP],
		resolved.PreStats[game.StatStamina],
		resolved.PreStats[game.StatNinjutsu],
		resolved.PreStats[game.StatGenjutsu],
		resolved.PreStats[game.StatTaijutsu],
		resolved.PreStats[game.StatSpeed],
		resolved.PreStats[game.StatEvasion],
		resolved.PreStats[game.StatAccuracy],
	)

	fmt.Printf("BaseStats: hp=%d stamina=%d ninjutsu=%d genjutsu=%d taijutsu=%d speed=%d evasion=%d accuracy=%d\n",
		resolved.BaseStats[game.StatHP],
		resolved.BaseStats[game.StatStamina],
		resolved.BaseStats[game.StatNinjutsu],
		resolved.BaseStats[game.StatGenjutsu],
		resolved.BaseStats[game.StatTaijutsu],
		resolved.BaseStats[game.StatSpeed],
		resolved.BaseStats[game.StatEvasion],
		resolved.BaseStats[game.StatAccuracy],
	)

	fmt.Printf("Stages: ninjutsu=%d genjutsu=%d taijutsu=%d speed=%d evasion=%d accuracy=%d\n",
		resolved.Stages[game.StatNinjutsu],
		resolved.Stages[game.StatGenjutsu],
		resolved.Stages[game.StatTaijutsu],
		resolved.Stages[game.StatSpeed],
		resolved.Stages[game.StatEvasion],
		resolved.Stages[game.StatAccuracy],
	)

	fmt.Printf("NatureDamage: fire=%.2f wind=%.2f lightning=%.2f earth=%.2f water=%.2f yin=%.2f yang=%.2f\n",
		resolved.NatureDamage[game.NatureFire],
		resolved.NatureDamage[game.NatureWind],
		resolved.NatureDamage[game.NatureLightning],
		resolved.NatureDamage[game.NatureEarth],
		resolved.NatureDamage[game.NatureWater],
		resolved.NatureDamage[game.NatureYin],
		resolved.NatureDamage[game.NatureYang],
	)

	fmt.Printf("NatureResistance: fire=%.2f wind=%.2f lightning=%.2f earth=%.2f water=%.2f yin=%.2f yang=%.2f\n",
		resolved.NatureResistance[game.NatureFire],
		resolved.NatureResistance[game.NatureWind],
		resolved.NatureResistance[game.NatureLightning],
		resolved.NatureResistance[game.NatureEarth],
		resolved.NatureResistance[game.NatureWater],
		resolved.NatureResistance[game.NatureYin],
		resolved.NatureResistance[game.NatureYang],
	)

	fmt.Printf("Natures=%v\n", resolved.Natures)
	fmt.Printf("AppliedModifiers=%v\n", resolved.AppliedModifiers)
}
