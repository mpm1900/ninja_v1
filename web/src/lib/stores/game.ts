import { Store } from '@tanstack/store'
import type { Game } from '../game/game'

const gameStore = new Store<Game>({
  status: 'init',
  actors: [],
  actions: [],
  modifiers: [],
  players: [],
  prompt: null,
  log: [],
})

export { gameStore }
