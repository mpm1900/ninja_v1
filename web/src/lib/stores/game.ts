import { Store } from '@tanstack/store'
import type { Game } from '../game/game'
import { NULL_CONTEXT } from '../game/context'

const gameStore = new Store<Game>({
  status: 'init',
  turn: {
    count: 0,
    phase: 'init'
  },
  actors: [],
  actions: [],
  modifiers: [],
  players: [],
  prompt: null,
  log: [],
  active_context: NULL_CONTEXT,
  queued_actions: {},
})

export { gameStore }
