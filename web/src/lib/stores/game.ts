import { Store } from '@tanstack/store'
import type { Game } from '../game/game'

const gameStore = new Store<Game>({
  status: 'idle',
  actors: [],
  modifiers: [],
  players: [],
})

export { gameStore }
