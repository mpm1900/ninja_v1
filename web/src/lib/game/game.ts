import type { ActionTransaction } from './action'
import type { Actor } from './actor'
import type { ModifierTransaction } from './modifier'
import type { Player } from './player'

type Game = {
  status: 'idle' | 'running'
  actors: Actor[]
  modifiers: ModifierTransaction[]
  players: Player[]

  actions: ActionTransaction[]
  prompt: ActionTransaction | null
}

export type { Game }
