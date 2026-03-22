package main

import (
	"encoding/json"
	"fmt"

	"ninja_v1/internal/game"
	actors "ninja_v1/internal/game/data/actors"

	"github.com/google/uuid"
)

func main1() {
	// Action has unexported fields, so build zero-values first, then include them.
	var action1 game.Action
	var action2 game.Action

	actor := actors.NewItachi(uuid.New(), 50)
	actor.Actions = []game.Action{action1, action2}

	var g game.Game
	modifier := game.Modifier{
		ID:   uuid.New(),
		Name: "Test",
	}
	modifier.Mutations = []game.ModifierMutation{
		{
			ModifierID: &modifier.ID,
			Mutation: game.Mutation[game.Actor, game.Actor, game.Context]{
				Filter: func(actor game.Actor, context *game.Context) bool {
					return actor.ID == *context.SourceActorID
				},
				Delta: func(actor game.Actor, context *game.Context) game.Actor {
					actor.Stats[game.StatNinjutsu] = actor.Stats[game.StatNinjutsu] * 2
					return actor
				},
			},
		},
	}
	resolved := game.ResolveActor(actor, g)

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
