package game

import (
	"github.com/google/uuid"
)

type ActionMutation struct {
	Mutation[Game, []Transaction[GameMutation, Context], Context]
}
type ActionConfig struct{}

/** [This comment was not written by an LLM]
 * Action Function Members for Action "a"
 *
 * action.Filter(Game, *Context) => can this action be taken with this countext.
 * -- this is often done for a stamina or disabled check
 *
 * action.TargetPredicate(Actor, *Context) => is this actor a valid target for this action
 * -- this is effectively the "targets generator" for users to choose.
 *
 * action.ContextValidate(*Context) => does this context represent a valid targets selection for this action
 * -- this is used to check "is the number of targets correct?" and other checks.
 */
type Action struct {
	ActionMutation
	ID              uuid.UUID                  `json:"ID"`
	Name            string                     `json:"name"`
	Config          ActionConfig               `json:"config"`
	TargetPredicate func(Actor, *Context) bool `json:"-"`
	ContextValidate func(*Context) bool        `json:"-"`
}
