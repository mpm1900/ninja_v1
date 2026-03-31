package game

import (
	"fmt"

	"github.com/google/uuid"
)

func (g *Game) Validate() bool {
	valid := true
	for _, player := range g.Players {
		missing_pos := make([]uuid.UUID, 0)
		for _, pos := range player.Positions {
			if pos.ActorID == nil {
				missing_pos = append(missing_pos, pos.ID)
				continue
			}

			ok, actor := g.GetActorByID(*pos.ActorID)
			if !ok {
				missing_pos = append(missing_pos, pos.ID)
				continue
			}

			if !actor.Alive {
				missing_pos = append(missing_pos, pos.ID)
				transaction := MakeTransaction(RemovePositions, Context{
					TargetActorIDs: []uuid.UUID{actor.ID},
				})
				g.JumpTransaction(transaction)
			}
		}

		if len(missing_pos) > 0 {
			fmt.Printf("%s needs %v\n", player.ID, missing_pos)
			action := SwitchIn(len(missing_pos))

			context := NewContext()
			context.SourcePlayerID = &player.ID
			context.TargetPositionIDs = missing_pos
			possible_targets := g.GetActors(func(a Actor) bool {
				return action.TargetPredicate(a, context)
			})

			if len(possible_targets) == 0 {
				valid = true
				fmt.Printf("Invalid state, but no possible targets, likely game-over. \n")
				continue
			}

			switch_count := min(len(missing_pos), len(possible_targets))
			action = SwitchIn(switch_count)
			context.ActionID = &action.ID
			transaction := MakeTransaction(action, context)
			transaction.Ready = false
			if !g.HasPlayerPrompt(player.ID) {
				g.AddPrompt(transaction)
			}
			valid = false
		}
	}

	return valid
}

func (g *Game) NextPhase() {
	switch g.Turn.Phase {
	case TurnInit, TurnStart:
		g.Turn.Phase = TurnMain
	case TurnMain:
		g.Turn.Phase = TurnEnd
	case TurnEnd:
		g.Turn.Phase = TurnCleanup
	case TurnCleanup:
		// Keep cleanup stable so callers can run end-of-turn bookkeeping once
		// without immediately wrapping back to main in the same loop tick.
	}
}

func (g *Game) NextTransaction() bool {
	transaction, err := g.Transactions.Dequeue()
	if err != nil {
		return false
	}

	n, ok := ResolveTransaction(*g, transaction, *g)
	if !ok {
		return false
	}

	*g = n
	return true
}

func (g *Game) NextAction() bool {
	transaction, err := g.Actions.Dequeue()
	g.ActiveContext = &transaction.Context

	if err != nil {
		g.ActiveContext = nil
		return false
	}

	g.RunAction(transaction)
	return true
}

func (g *Game) NextPrompt() bool {
	transaction, err := g.Prompts.Dequeue()
	g.ActiveContext = &transaction.Context
	if err != nil {
		return false
	}

	g.RunPrompt(transaction)
	return true
}

func (g *Game) NextTrigger() bool {
	transaction, err := g.Triggers.Dequeue()
	g.ActiveContext = &transaction.Context
	if err != nil {
		return false
	}

	g.RunTrigger(transaction)
	return true
}

func (g *Game) Next() bool {
	if g.NextTransaction() {
		return true
	}

	if g.NextTrigger() {
		return true
	}

	if g.AllPromptsReady() {
		if g.NextPrompt() {
			return true
		}
	}

	if !g.Validate() {
		return false
	}

	if g.NextAction() {
		return true
	}

	g.ActiveContext = nil
	g.NextPhase()

	return false
}
