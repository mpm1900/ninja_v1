import type { Actor } from './actor'
import type { ModifierTransaction } from './modifier'

type Context = {
  sourcePlayerID: string | null

  parentActorID: string | null
  sourceActorID: string | null

  targetActorIDs: Array<string>
  targetPositionIDs: Array<string>
}

type Game = {
  actors: Actor[]
  modifiers: ModifierTransaction[]

  // TODO
}

export type { Game, Context }
