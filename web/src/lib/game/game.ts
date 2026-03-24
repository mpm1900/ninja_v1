import type { Actor } from './actor'
import type { ModifierTransaction } from './modifier'

type Game = {
  status: 'idle' | 'running'
  actors: Actor[]
  modifiers: ModifierTransaction[]

  // TODO
}

export type { Game }
