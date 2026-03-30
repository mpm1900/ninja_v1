import { Store } from '@tanstack/store'
import type { Game } from '../game/game'

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
})

export { gameStore }
