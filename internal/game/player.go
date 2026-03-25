package game

import (
	"fmt"

	"github.com/google/uuid"
)

type Player struct {
	ID                uuid.UUID                `json:"ID"`
	PositionsCapacity int                      `json:"positions_capacity"`
	Positions         map[uuid.UUID]*uuid.UUID `json:"positions"`
}

func NewPlayer(ID uuid.UUID) Player {
	return Player{
		ID:                ID,
		PositionsCapacity: 2,
		Positions:         make(map[uuid.UUID]*uuid.UUID),
	}
}

func (p *Player) EnsureOpenPositionID() *uuid.UUID {
	for pid, aid := range p.Positions {
		if aid == nil {
			id := pid
			return &id
		}
	}

	if len(p.Positions) >= p.PositionsCapacity {
		return nil
	}

	id := uuid.New()
	p.Positions[id] = nil
	return &id
}

func (p *Player) SetPosition(pid uuid.UUID, aid *uuid.UUID) error {
	_, ok := p.Positions[pid]

	// if this position exists, just set it,
	// this effectively means "swap" or "clear"
	if ok {
		p.Positions[pid] = aid
		return nil
	}

	for openPID, openAID := range p.Positions {
		if openAID == nil {
			p.Positions[openPID] = aid
			return nil
		}
	}

	if len(p.Positions) < p.PositionsCapacity {
		p.Positions[pid] = aid
		return nil
	}

	return fmt.Errorf("[SetPosition] pid didn't exist and there's no room.")
}

func (p *Player) AddPosition(aid *uuid.UUID) error {
	return p.SetPosition(uuid.New(), aid)
}
