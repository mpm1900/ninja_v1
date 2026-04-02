import type { ActionTransaction } from './action'
import type { Actor } from './actor'
import type { Context } from './context'
import type { GameLog } from './log'
import type { ModifierTransaction } from './modifier'
import type { Player } from './player'

type Game = {
  status: 'init' | 'idle' | 'running' | 'waiting'
  turn: {
    count: number,
    phase: 'init' | 'start' | 'main' | 'end' | 'cleanup'
  },
  active_context: Context | null
  actors: Actor[]
  modifiers: ModifierTransaction[]
  players: Player[]

  actions: ActionTransaction[]
  prompt: ActionTransaction | null
  log: GameLog[]

  queued_actions: Record<string, {
    ID: string
    context: Context
    mutation: string
  }>
}

export type { Game }
