package instance

import (
	"ninja_v1/internal/game"
)

type Request struct {
	Type    string       `json:"type"`
	Context game.Context `json:"context"`
}
