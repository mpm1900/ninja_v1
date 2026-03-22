import type { Actor } from './actor'
import type { ModifierTransaction } from './modifier'

type Context = {
  source_player_ID: string | null

  parent_actor_ID: string | null
  source_actor_ID: string | null

  target_actor_IDs: Array<string>
  target_position_IDs: Array<string>
}

type Game = {
  actors: Actor[]
  modifiers: ModifierTransaction[]

  // TODO
}

export type { Game, Context }
